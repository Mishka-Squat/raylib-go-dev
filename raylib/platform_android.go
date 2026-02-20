//go:build android
// +build android

package rl

/*
#include "raylib.h"
#include "platforms/raylib_android.h"
#include <stdlib.h>
#include <android/asset_manager.h>

static AAssetManager* GetAssetManager() {
	return GetAndroidApp()->activity->assetManager;
}

static const char* GetInternalDataPath() {
	return GetAndroidApp()->activity->internalDataPath;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"strings"
	"time"
	"unsafe"

	"golang.org/x/exp/constraints"
)

var callbackHolder func()

// InitWindow - Initialize Window and OpenGL Graphics
func InitWindow[WT, HT constraints.Integer](width WT, height HT, title string) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	ctitle := textAlloc(title)

	C.InitWindow(cwidth, cheight, ctitle)

	SetLoadFileDataCallback(func(fileName string) []byte {
		asset, err := OpenAsset(fileName)
		if err != nil {
			return nil
		}
		data, err := ReadAll(asset)
		if err != nil {
			return nil
		}

		return data
	})
	SetLoadFileTextCallback(func(fileName string) string {
		asset, err := OpenAsset(fileName)
		if err != nil {
			return ""
		}
		data, err := ReadAll(asset)
		if err != nil {
			return ""
		}

		return string(data)
	})
}

var androidMainFn func()

func RayLibANativeActivity_onCreate(activity unsafe.Pointer, savedState unsafe.Pointer, savedStateSize uint, mainFn func()) {
	androidMainFn = mainFn
	C.RayLibANativeActivity_onCreate((*C.ANativeActivity)(activity), savedState, (C.size_t)(savedStateSize))
}

//export android_run
func android_run() {
	androidMainFn()
}

// ShowCursor - Shows cursor
func ShowCursor() {
	return
}

// HideCursor - Hides cursor
func HideCursor() {
	return
}

// IsCursorHidden - Returns true if cursor is not visible
func IsCursorHidden() bool {
	return false
}

// IsCursorOnScreen - Check if cursor is on the current screen.
func IsCursorOnScreen() bool {
	return false
}

// EnableCursor - Enables cursor
func EnableCursor() {
	return
}

// DisableCursor - Disables cursor
func DisableCursor() {
	return
}

// IsFileDropped - Check if a file have been dropped into window
func IsFileDropped() bool {
	return false
}

// LoadDroppedFiles - Load dropped filepaths
func LoadDroppedFiles() (files []string) {
	return
}

// UnloadDroppedFiles - Unload dropped filepaths
func UnloadDroppedFiles() {
	return
}

// Open implements fs.FS interface - opens the named file for reading
func (a Asset) Open(name string) (fs.File, error) {
	return openAssetFile(a.root, name)
}

// ReadFile implements fs.ReadFileFS interface - reads the entire file
func (a Asset) ReadFile(name string) ([]byte, error) {
	return readAssetFile(a.root, name)
}

// ReadDir implements fs.ReadDirFS interface - reads the directory
func (a Asset) ReadDir(name string) ([]fs.DirEntry, error) {
	return readAssetDir(a.root, name)
}

// androidAsset implements fs.File interface for Android assets
type androidAsset struct {
	ptr  *C.AAsset
	name string
	size int64
}

func (a androidAsset) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	n = int(C.AAsset_read(a.ptr, unsafe.Pointer(&p[0]), C.size_t(len(p))))
	if n == 0 && len(p) > 0 {
		return 0, io.EOF
	}

	return n, nil
}

func (a androidAsset) Seek(offset int64, whence int) (int64, error) {
	off := C.AAsset_seek(a.ptr, C.off_t(offset), C.int(whence))
	if off == -1 {
		return 0, fmt.Errorf("bad result for offset=%d, whence=%d", offset, whence)
		return 0, fmt.Errorf("seek failed for offset=%d, whence=%d", offset, whence)
	}

	return int64(off), nil
}

func (a androidAsset) Close() error {
	C.AAsset_close(a.ptr)

	return nil
}

func (a androidAsset) Stat() (fs.FileInfo, error) {
	return &androidAssetInfo{
		name: path.Base(a.name),
		size: a.size,
		mode: fs.FileMode(0444), // read-only
	}, nil
}

// androidAssetInfo implements fs.FileInfo for Android assets
type androidAssetInfo struct {
	name string
	size int64
	mode fs.FileMode
	dir  bool
}

func (i androidAssetInfo) Name() string       { return i.name }
func (i androidAssetInfo) Size() int64        { return i.size }
func (i androidAssetInfo) Mode() fs.FileMode  { return i.mode }
func (i androidAssetInfo) ModTime() time.Time { return time.Time{} }
func (i androidAssetInfo) IsDir() bool        { return i.dir }
func (i androidAssetInfo) Sys() interface{}   { return nil }

// androidDirEntry implements fs.DirEntry for Android assets
type androidDirEntry struct {
	name string
	dir  bool
}

func (e androidDirEntry) Name() string               { return e.name }
func (e androidDirEntry) IsDir() bool                { return e.dir }
func (e androidDirEntry) Type() fs.FileMode          { return e.Mode().Type() }
func (e androidDirEntry) Info() (fs.FileInfo, error) { return e.fileInfo(), nil }

func (e androidDirEntry) Mode() fs.FileMode {
	if e.dir {
		return fs.FileMode(0555) | fs.ModeDir
	}

	return fs.FileMode(0444)
}

func (e androidDirEntry) fileInfo() fs.FileInfo {
	return &androidAssetInfo{
		name: e.name,
		size: 0,
		mode: e.Mode(),
		dir:  e.dir,
	}
}

func openAssetFile(root, name string) (fs.File, error) {
	fullPath := name
	if root != "" {
		fullPath = path.Join(root, name)
	}

	cname := C.CString(fullPath)
	defer C.free(unsafe.Pointer(cname))

	asset := C.AAssetManager_open(C.GetAssetManager(), cname, C.AASSET_MODE_UNKNOWN)
	if asset == nil {
		if C.IsAssetDir(cname) != 0 {
			return nil, fmt.Errorf("cannot open directory as file: %s", fullPath)
		}

		return nil, fmt.Errorf("asset file not found: %s", fullPath)
	}

	size := int64(C.GetAssetLength(asset))

	return &androidAsset{
		ptr:  asset,
		name: fullPath,
		size: size,
	}, nil
}

func readAssetFile(root, name string) ([]byte, error) {
	file, err := openAssetFile(root, name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, info.Size())
	n, err := io.ReadFull(file, data)
	if err != nil && err != io.EOF && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, err
	}

	return data[:n], nil
}

func readAssetDir(root, name string) ([]fs.DirEntry, error) {
	fullPath := name
	if root != "" {
		fullPath = path.Join(root, name)
	}
	if fullPath == "." {
		fullPath = ""
	}

	cpath := C.CString(fullPath)
	defer C.free(unsafe.Pointer(cpath))

	dir := C.AAssetManager_openDir(C.GetAssetManager(), cpath)
	if dir == nil {
		return nil, fmt.Errorf("cannot open directory: %s", fullPath)
	}
	defer C.AAssetDir_close(dir)

	var entries []fs.DirEntry
	seenNames := make(map[string]bool)

	for {
		cfilename := C.AAssetDir_getNextFileName(dir)
		if cfilename == nil {
			break
		}

		filename := C.GoString(cfilename)

		// Extract immediate child name (handle nested paths)
		parts := strings.SplitN(filename, "/", 2)
		childName := parts[0]

		if !seenNames[childName] {
			seenNames[childName] = true

			// Check if it's a directory by checking if there are more parts
			isDir := len(parts) > 1
			if !isDir {
				// Double-check if it's actually a directory
				checkPath := path.Join(fullPath, childName)
				cCheckPath := C.CString(checkPath)
				isDir = C.IsAssetDir(cCheckPath) != 0
				C.free(unsafe.Pointer(cCheckPath))
			}

			entries = append(entries, &androidDirEntry{
				name: childName,
				dir:  isDir,
			})
		}
	}

	return entries, nil
}

func getInternalStoragePath() string {
	return C.GoString(C.GetInternalDataPath())
}
