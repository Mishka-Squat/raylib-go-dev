package rl

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -std=gnu99 -Wno-missing-braces -Wno-unused-result -Wno-implicit-function-declaration

#include "external/stb_rect_pack.h"
*/
import "C"
import (
	"fmt"
	"iter"
	"runtime"
	"slices"
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/goex/image/colorex"
)

type PackContext struct {
	context C.stbrp_context
	nodes   []C.stbrp_node
}

type PackRect C.stbrp_rect

func (r PackRect) Id() int {
	return int(r.id)
}

func (r PackRect) Rect() rect2.Float32 {
	return MakeRectangle(int(r.x), int(r.y), int(r.w), int(r.h))
}

func (r PackRect) X() int {
	return int(r.x)
}

func (r PackRect) Y() int {
	return int(r.y)
}

func (r PackRect) W() int {
	return int(r.w)
}

func (r PackRect) H() int {
	return int(r.h)
}

func PackRects(width, height int, rects ...PackRect) ([]PackRect, error) {
	ctx := NewPackContext(width, height, len(rects))
	return ctx.PackRects(rects...)
}

func NewPackContext(width, height, size int) *PackContext {
	ctx := &PackContext{
		nodes: make([]C.stbrp_node, size),
	}
	C.stbrp_init_target(&ctx.context, C.int(width), C.int(height), unsafe.SliceData(ctx.nodes), C.int(len(ctx.nodes)))
	return ctx
}

func NewRect(id int, w, h int) PackRect {
	return PackRect{
		id: C.int(id),
		w:  C.stbrp_coord(w),
		h:  C.stbrp_coord(h),
	}
}

func (ctx *PackContext) PackRects(rects ...PackRect) ([]PackRect, error) {
	p := runtime.Pinner{}
	defer p.Unpin()

	crects := unsafe.SliceData(rects)
	p.Pin(ctx.context.active_head)
	p.Pin(ctx.context.free_head)
	res := C.stbrp_pack_rects(&ctx.context, (*C.stbrp_rect)(crects), C.int(len(rects)))

	if res == 0 {
		return nil, fmt.Errorf("failed to pack rects")
	}
	return rects, nil
}

type ImageAtlas struct {
	Image Image
	Atlas []rect2.Float32
}

func LoadImageAtlas(width, height int, fileNames ...string) (ImageAtlas, error) {
	images := make([]Image, len(fileNames))
	defer func() {
		for _, i := range images {
			i.Unload()
		}
	}()
	for i := range images {
		images[i] = LoadImage(fileNames[i])
	}

	rect := make([]PackRect, len(images))
	for i, img := range images {
		img_size := img.GetSize().AddXY(2, 2)
		rect[i].w = C.int(img_size.X)
		rect[i].h = C.int(img_size.Y)
	}
	icons, err := PackRects(width, height, rect...)
	if err != nil {
		return ImageAtlas{}, err
	}

	atlasImage := GenImageColor(width, height, Blank)
	atlasAtlas := make([]rect2.Float32, len(icons))
	for i, icon := range icons {
		rect := icon.Rect().ShrinkXYWH(1, 1, 1, 1)
		atlasAtlas[i] = rect
		images[i].DrawDef(&atlasImage, rect)
	}

	return ImageAtlas{
		Image: atlasImage,
		Atlas: atlasAtlas,
	}, nil
}

func LoadImageAtlasEx(width, height int, imgFn func(path string) Image, fileNames ...string) (ImageAtlas, error) {
	images := make([]Image, len(fileNames))
	defer func() {
		for _, i := range images {
			i.Unload()
		}
	}()
	for i := range images {
		images[i] = imgFn(fileNames[i])
	}

	rect := make([]PackRect, len(images))
	for i, img := range images {
		img_size := img.GetSize().AddXY(2, 2)
		rect[i].w = C.int(img_size.X)
		rect[i].h = C.int(img_size.Y)
	}
	icons, err := PackRects(width, height, rect...)
	if err != nil {
		return ImageAtlas{}, err
	}

	atlasImage := GenImageColor(width, height, Blank)
	atlasAtlas := make([]rect2.Float32, len(icons))
	for i, icon := range icons {
		rect := icon.Rect().ShrinkXYWH(1, 1, 1, 1)
		atlasAtlas[i] = rect
		images[i].DrawDef(&atlasImage, rect)
	}

	return ImageAtlas{
		Image: atlasImage,
		Atlas: atlasAtlas,
	}, nil
}

func (i *ImageAtlas) Unload() {
	li := i.Image
	i.Image = Image{}
	li.Unload()
}

type TextureAtlasItem struct {
	Texture *Texture2D
	Rect    rect2.Float32
}

func (t TextureAtlasItem) DrawExDef(position vector2.Float32) {
	DrawTexturePro(*t.Texture, t.Rect, rect2.NewFloat32(position, t.Rect.Size), vector2.Zero[float32](), 0, White)
}

func (t TextureAtlasItem) DrawProDef(destRec rect2.Float32) {
	DrawTexturePro(*t.Texture, t.Rect, destRec, vector2.Zero[float32](), 0, White)
}

func (t TextureAtlasItem) DrawProTintedDef(destRec rect2.Float32, tint colorex.RGBA) {
	DrawTexturePro(*t.Texture, t.Rect, destRec, vector2.Zero[float32](), 0, tint)
}

type TextureAtlas struct {
	Texture Texture2D
	Atlas   []rect2.Float32
}

func LoadTextureAtlas(width, height int, fileNames ...string) (TextureAtlas, error) {
	ia, err := LoadImageAtlas(width, height, fileNames...)
	if err != nil {
		return TextureAtlas{}, err
	}
	defer ia.Unload()

	iaImage := ia.Image
	return TextureAtlas{
		Texture: LoadTextureFromImage(iaImage),
		Atlas:   ia.Atlas,
	}, nil
}

func LoadTextureAtlasSeq(width, height int, fileNames iter.Seq[string]) (TextureAtlas, error) {
	return LoadTextureAtlas(width, height, slices.Collect(fileNames)...)
}

func LoadTextureAtlasEx(width, height int, imgFn func(path string) Image, fileNames ...string) (TextureAtlas, error) {
	ia, err := LoadImageAtlasEx(width, height, imgFn, fileNames...)
	if err != nil {
		return TextureAtlas{}, err
	}
	defer ia.Unload()

	iaImage := ia.Image
	return TextureAtlas{
		Texture: LoadTextureFromImage(iaImage),
		Atlas:   ia.Atlas,
	}, nil
}

func LoadTextureAtlasExSeq(width, height int, imgFn func(path string) Image, fileNames iter.Seq[string]) (TextureAtlas, error) {
	return LoadTextureAtlasEx(width, height, imgFn, slices.Collect(fileNames)...)
}

func (t *TextureAtlas) Unload() {
	lt := t.Texture
	t.Texture = Texture2D{}
	lt.Unload()
}

func (t *TextureAtlas) GetItem(id int) TextureAtlasItem {
	return TextureAtlasItem{
		Texture: &t.Texture,
		Rect:    t.Atlas[id],
	}
}

func (t *TextureAtlas) GetItemSet(ids ...int) []TextureAtlasItem {
	itemSet := make([]TextureAtlasItem, len(ids))
	for i, id := range ids {
		itemSet[i] = t.GetItem(id)
	}
	return itemSet
}

func (t *TextureAtlas) GetItemSetIter(ids iter.Seq[int]) iter.Seq[TextureAtlasItem] {
	return func(yield func(TextureAtlasItem) bool) {
		for id := range ids {
			if !yield(t.GetItem(id)) {
				return
			}
		}
	}
}

func (t *TextureAtlas) ToItemSet() []TextureAtlasItem {
	itemSet := make([]TextureAtlasItem, len(t.Atlas))
	for i, item := range t.Atlas {
		itemSet[i] = TextureAtlasItem{
			Texture: &t.Texture,
			Rect:    item,
		}
	}
	return itemSet
}
