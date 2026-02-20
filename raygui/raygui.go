package raygui

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/raygui/src -std=gnu99 -Wno-unused-result

#define RAYGUI_IMPLEMENTATION
#include "raygui.h"
#include <stdlib.h>
*/
import "C"

import (
	"slices"
	"strings"
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
	"golang.org/x/exp/constraints"
)

type BoundsFn func(bounds rect2.Float32)

const (
	SCROLLBAR_LEFT_SIDE = iota
	SCROLLBAR_RIGHT_SIDE
)

// StyleProp - Style property
type StyleProp struct {
	controlId     uint16
	propertyId    uint16
	propertyValue uint32
}

// State .
type State int8

const (
	STATE_NORMAL State = iota
	STATE_FOCUSED
	STATE_PRESSED
	STATE_DISABLED

	STATE_MAX
)

func (s State) ColorProperty(base ControlProperty) ControlProperty {
	return base + ControlProperty(s*PropertyControlColorStride)
}

// TextAlignment .
type TextAlignment = int8

// Gui control text alignment
const (
	TEXT_ALIGN_LEFT TextAlignment = iota
	TEXT_ALIGN_CENTER
	TEXT_ALIGN_RIGHT
)

// TextAlignmentVertical .
type TextAlignmentVertical = int8

// Gui control text alignment vertical
const (
	TEXT_ALIGN_TOP TextAlignmentVertical = iota
	TEXT_ALIGN_MIDDLE
	TEXT_ALIGN_BOTTOM
)

// TextWrapMode .
type TextWrapMode = int8

// Gui control text wrap mode
// NOTE: Useful for multiline text
const (
	TEXT_WRAP_NONE TextWrapMode = iota
	TEXT_WRAP_CHAR
	TEXT_WRAP_WORD
)

// Control .
type Control = int8

// DEFAULT - Gui controls
const (
	Control_DEFAULT Control = iota
	Control_LABEL
	Control_BUTTON
	Control_TOGGLE
	Control_SLIDER
	Control_PROGRESSBAR
	Control_CHECKBOX
	Control_COMBOBOX
	Control_DROPDOWNBOX
	Control_TEXTBOX
	Control_VALUEBOX
	Control_SPINNER
	Control_LISTVIEW
	Control_COLORPICKER
	Control_SCROLLBAR
	Control_STATUSBAR
)

// ControlProperty .
type ControlProperty = int8

// Gui base properties for every control
// NOTE: RAYGUI_MAX_PROPS_BASE properties (by default 16 properties)
const (
	ControlProperty_BORDER_COLOR_NORMAL ControlProperty = iota
	ControlProperty_BASE_COLOR_NORMAL
	ControlProperty_TEXT_COLOR_NORMAL
	ControlProperty_BORDER_COLOR_FOCUSED
	ControlProperty_BASE_COLOR_FOCUSED
	ControlProperty_TEXT_COLOR_FOCUSED
	ControlProperty_BORDER_COLOR_PRESSED
	ControlProperty_BASE_COLOR_PRESSED
	ControlProperty_TEXT_COLOR_PRESSED
	ControlProperty_BORDER_COLOR_DISABLED
	ControlProperty_BASE_COLOR_DISABLED
	ControlProperty_TEXT_COLOR_DISABLED
	ControlProperty_BORDER_WIDTH
	ControlProperty_TEXT_PADDING
	ControlProperty_TEXT_ALIGNMENT
)

const PropertyControlColorStride = 3 // BORDER, BASE, TEXT

// DEFAULT extended properties
// NOTE: Those properties are common to all controls or global
const (
	DefaultProperty_TEXT_SIZE ControlProperty = iota + 16
	DefaultProperty_TEXT_SPACING
	DefaultProperty_LINE_COLOR
	DefaultProperty_BACKGROUND_COLOR
	DefaultProperty_TEXT_LINE_SPACING
	DefaultProperty_TEXT_ALIGNMENT_VERTICAL
	DefaultProperty_TEXT_WRAP_MODE
)

const (
	ToggleProperty_GROUP_PADDING ControlProperty = 16
)

const (
	// Slider size of internal bar
	SliderProperty_SLIDER_WIDTH ControlProperty = 16
	// Slider/SliderBar internal bar padding
	SliderProperty_SLIDER_PADDING = 17
)

const (
	// ProgressBar internal padding
	ProgressBarProperty_PROGRESS_PADDING ControlProperty = 16
)

const (
	ScrollBarProperty_ARROWS_SIZE ControlProperty = iota + 16
	ScrollBarProperty_ARROWS_VISIBLE
	ScrollBarProperty_SCROLL_SLIDER_PADDING
	ScrollBarProperty_SCROLL_SLIDER_SIZE
	ScrollBarProperty_SCROLL_PADDING
	ScrollBarProperty_SCROLL_SPEED
)

const (
	CheckBoxProperty_CHECK_PADDING ControlProperty = 16
)

const (
	// ComboBox right button width
	ComboBoxProperty_COMBO_BUTTON_WIDTH ControlProperty = 16
	// ComboBox button separation
	ComboBoxProperty_COMBO_BUTTON_SPACING = 17
)

const (
	// DropdownBox arrow separation from border and items
	DropdownBoxProperty_ARROW_PADDING ControlProperty = 16
	// DropdownBox items separation
	DropdownBoxProperty_DROPDOWN_ITEMS_SPACING = 17
)

const (
	// TextBox/TextBoxMulti/ValueBox/Spinner inner text padding
	TextBoxProperty_TEXT_INNER_PADDING ControlProperty = 16
	// TextBoxMulti lines separation
	TextBoxProperty_TEXT_LINES_SPACING = 17
)

const (
	// Spinner left/right buttons width
	SpinnerProperty_SPIN_BUTTON_WIDTH ControlProperty = 16
	// Spinner buttons separation
	SpinnerProperty_SPIN_BUTTON_SPACING = 17
)

const (
	// ListView items height
	ListViewProperty_LIST_ITEMS_HEIGHT ControlProperty = 16
	// ListView items separation
	ListViewProperty_LIST_ITEMS_SPACING = 17
	// ListView scrollbar size (usually width)
	ListViewProperty_SCROLLBAR_WIDTH = 18
	// ListView scrollbar side (0-left, 1-right)
	ListViewProperty_SCROLLBAR_SIDE = 19
)

const (
	ColorPickerProperty_COLOR_SELECTOR_SIZE ControlProperty = 16
	// colorex.RGBA right hue bar width
	ColorPickerProperty_HUEBAR_WIDTH = 17
	// colorex.RGBA right hue bar separation from panel
	ColorPickerProperty_HUEBAR_PADDING = 18
	// colorex.RGBA right hue bar selector height
	ColorPickerProperty_HUEBAR_SELECTOR_HEIGHT = 19
	// colorex.RGBA right hue bar selector overflow
	ColorPickerProperty_HUEBAR_SELECTOR_OVERFLOW = 20
)

//func GetTextSize(text string) vector2.Float32 {
//	ctext := textAlloc(text)
//	ret := C.GetTextSize(ctext)
//	return *govec2ptr(&ret)
//}

// GuiEnable - Enable gui controls (global state)
func Enable() {
	C.GuiEnable()
}

// GuiDisable - Disable gui controls (global state)
func Disable() {
	C.GuiDisable()
}

// GuiLock - Lock gui controls (global state)
func Lock() {
	C.GuiLock()
}

// GuiUnlock - Unlock gui controls (global state)
func Unlock() {
	C.GuiUnlock()
}

// GuiIsLocked - Check if gui is locked (global state)
func IsLocked() bool {
	return bool(C.GuiIsLocked())
}

// GuiFade - Set gui controls alpha (global state), alpha goes from 0.0f to 1.0f
func Fade(color colorex.RGBA, alpha float32) {
	ccolor := (*C.Color)(unsafe.Pointer(&color))
	calpha := C.float(alpha)
	C.GuiFade(*ccolor, calpha)
}

// GuiSetState - Set gui state (global state)
func SetState(state State) {
	cstate := C.int(state)
	C.GuiSetState(cstate)
}

// GuiGetState - Get gui state (global state)
func GetState() State {
	return State(C.GuiGetState())
}

// GuiSetStyle .
func SetStyle[T constraints.Integer](control Control, property ControlProperty, value T) {
	ccontrol := C.int(control)
	cproperty := C.int(property)
	cvalue := C.int(value)
	C.GuiSetStyle(ccontrol, cproperty, cvalue)
}

func SetStyleColor(control Control, property ControlProperty, color colorex.RGBA) {
	ccontrol := C.int(control)
	cproperty := C.int(property)
	cvalue := C.int(rl.ColorToInt(color))
	C.GuiSetStyle(ccontrol, cproperty, cvalue)
}

// GuiGetStyle - Get one style property
func GetStyle[T constraints.Integer](control Control, property ControlProperty) T {
	ccontrol := C.int(control)
	cproperty := C.int(property)
	return T(C.GuiGetStyle(ccontrol, cproperty))
}

func GetStyleColor(control Control, property ControlProperty) colorex.RGBA {
	ccontrol := C.int(control)
	cproperty := C.int(property)
	return rl.GetColor(uint(C.GuiGetStyle(ccontrol, cproperty)))
}

// GuiWindowBox - Window Box control, shows a window that can be closed
func WindowBox(bounds rect2.Float32, title string) bool {
	cbounds := crect2ptr(&bounds)
	ctitle := textAlloc(title)
	return C.GuiWindowBox(*cbounds, ctitle) != 0
}

// GuiGroupBox - Group Box control with text name
func GroupBox(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiGroupBox(*cbounds, ctext)
}

func GroupBoxEx(bounds rect2.Float32, text string, fn BoundsFn) {
	GroupBox(bounds, text)
	fn(bounds.ShrinkXYWH(8, 8, 8, 8))
}

// GuiLine - Line separator control, could contain text
func Line(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiLine(*cbounds, ctext)
}

// Panel - Panel control, useful to group controls
func Panel(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiPanel(*cbounds, ctext)
}

func PanelEx(bounds rect2.Float32, text string, fn BoundsFn) {
	Panel(bounds, text)
	fn(bounds.ShrinkXYWH(0, 24, 0, 0))
}

// ScrollPanel control - Scroll Panel control
func ScrollPanel(bounds rect2.Float32, text string, content rect2.Float32, scroll *vector2.Float32, view *rect2.Float32) int32 {
	ctext := textAlloc(text)
	cbounds := crect2ptr(&bounds)
	ccontent := crect2ptr(&content)
	cscroll := cvec2ptr(scroll)
	cview := crect2ptr(view)

	res := C.GuiScrollPanel(*cbounds, ctext, *ccontent, cscroll, cview)

	return int32(res)
}

// ScrollBar control (used by GuiScrollPanel())
func ScrollBar(bounds rect2.Float32, value, minValue, maxValue int32) int32 {
	cbounds := crect2ptr(&bounds)
	cvalue := C.int(value)
	cminValue := C.int(minValue)
	cmaxValue := C.int(maxValue)

	return int32(C.GuiScrollBar(*cbounds, cvalue, cminValue, cmaxValue))
}

// Label control, shows text
func Label(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiLabel(*cbounds, ctext)
}

// Button control, returns true when clicked
func Button(bounds rect2.Float32, text string) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	return C.GuiButton(*cbounds, ctext) != 0
}

// LabelButton control, show true when clicked
func LabelButton(bounds rect2.Float32, text string) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	return C.GuiLabelButton(*cbounds, ctext) != 0
}

// Toggle control, returns true when active
func Toggle(bounds rect2.Float32, text string, active bool) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cactive := C.bool(active)
	C.GuiToggle(*cbounds, ctext, &cactive)
	return bool(cactive)
}

// ToggleGroup control, returns active toggle index
func ToggleGroup(bounds rect2.Float32, text string, active int32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cactive := C.int(active)
	C.GuiToggleGroup(*cbounds, ctext, &cactive)
	return int32(cactive)
}

// ToggleSlider control, returns true when clicked
func ToggleSlider(bounds rect2.Float32, text string, active int32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cactive := C.int(active)
	C.GuiToggleSlider(*cbounds, ctext, &cactive)
	return int32(cactive)
}

// CheckBox control, returns true when active
func CheckBox(bounds rect2.Float32, text string, checked bool) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cchecked := C.bool(checked)
	C.GuiCheckBox(*cbounds, ctext, &cchecked)
	return bool(cchecked)
}

// ComboBox control, returns selected item index
func ComboBox(bounds rect2.Float32, text string, active int32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cactive := C.int(active)
	C.GuiComboBox(*cbounds, ctext, &cactive)
	return int32(cactive)
}

type ComboBoxState[T any] struct {
	items           []T
	citemNames      *C.char
	activeItemIndex int32
}

func (s *ComboBoxState[T]) Free() {
	if s.citemNames != nil {
		C.free(unsafe.Pointer(s.citemNames))
	}
}

func (s *ComboBoxState[T]) SetItems(items []T, nameFn func(i T) string) {
	s.items = items
	itemNames := ""
	for i, v := range s.items {
		if i != 0 {
			itemNames += ";"
		}
		itemNames += nameFn(v)
	}

	s.Free()
	s.citemNames = C.CString(itemNames)
}

func (s *ComboBoxState[T]) IsActiveItem() bool {
	return s.activeItemIndex != -1
}

func (s *ComboBoxState[T]) ActiveItem() T {
	return s.items[s.activeItemIndex]
}

func (s *ComboBoxState[T]) SetActiveItemIndex(i int32) {
	s.activeItemIndex = i
}

func (s *ComboBoxState[T]) ComboBox(bounds rect2.Float32) {
	cbounds := crect2ptr(&bounds)
	cactive := (*C.int)(&s.activeItemIndex)

	C.GuiComboBox(*cbounds, s.citemNames, cactive)
}

// Spinner control, returns selected value
func Spinner(bounds rect2.Float32, text string, value *int32, minValue, maxValue int, editMode bool) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)

	if value == nil {
		value = new(int32)
	}
	cvalue := C.int(*value)
	defer func() {
		*value = int32(cvalue)
	}()

	cminValue := C.int(minValue)
	cmaxValue := C.int(maxValue)
	ceditMode := C.bool(editMode)

	C.GuiSpinner(*cbounds, ctext, &cvalue, cminValue, cmaxValue, ceditMode)
	return int32(cvalue)
}

// Slider control
func Slider(bounds rect2.Float32, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds := crect2ptr(&bounds)
	ctextLeft := textAlloc(textLeft)
	ctextRight := textAlloc(textRight)
	cvalue := C.float(value)
	cminValue := C.float(minValue)
	cmaxValue := C.float(maxValue)
	C.GuiSlider(*cbounds, ctextLeft, ctextRight, &cvalue, cminValue, cmaxValue)
	return float32(cvalue)
}

// SliderBar control, returns selected value
func SliderBar(bounds rect2.Float32, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds := crect2ptr(&bounds)
	ctextLeft := textAlloc(textLeft)
	ctextRight := textAlloc(textRight)
	cvalue := C.float(value)
	cminValue := C.float(minValue)
	cmaxValue := C.float(maxValue)
	C.GuiSliderBar(*cbounds, ctextLeft, ctextRight, &cvalue, cminValue, cmaxValue)
	return float32(cvalue)
}

// ProgressBar control, shows current progress value
func ProgressBar(bounds rect2.Float32, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	cbounds := crect2ptr(&bounds)
	ctextLeft := textAlloc(textLeft)
	ctextRight := textAlloc(textRight)
	cvalue := C.float(value)
	cminValue := C.float(minValue)
	cmaxValue := C.float(maxValue)
	C.GuiProgressBar(*cbounds, ctextLeft, ctextRight, &cvalue, cminValue, cmaxValue)
	return float32(cvalue)
}

// StatusBar control, shows info text
func StatusBar(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiStatusBar(*cbounds, ctext)
}

// DummyRec control for placeholders
func DummyRec(bounds rect2.Float32, text string) {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	C.GuiDummyRec(*cbounds, ctext)
}

// Grid control, returns mouse cell position
func Grid(bounds rect2.Float32, text string, spacing float32, subdivs int32, mouseCell *vector2.Float32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cspacing := C.float(spacing)
	csubdivs := C.int(subdivs)
	cmouseCell := cvec2ptr(mouseCell)
	res := C.GuiGrid(*cbounds, ctext, cspacing, csubdivs, cmouseCell)
	return int32(res)
}

// ListView control, returns selected list item index
func ListView(bounds rect2.Float32, text string, scrollIndex *int32, active int32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)

	if scrollIndex == nil {
		scrollIndex = new(int32)
	}
	cscrollIndex := C.int(*scrollIndex)
	defer func() {
		*scrollIndex = int32(cscrollIndex)
	}()

	cactive := C.int(active)

	C.GuiListView(*cbounds, ctext, &cscrollIndex, &cactive)
	return int32(cactive)
}

// MessageBox control, displays a message
func MessageBox(bounds rect2.Float32, title string, message string, buttons string) int32 {
	cbounds := crect2ptr(&bounds)
	ctitle := textAlloc(title)
	cmessage := textAlloc(message)
	cbuttons := textAlloc(buttons)
	return int32(C.GuiMessageBox(*cbounds, ctitle, cmessage, cbuttons))
}

// ColorPicker control (multiple color controls)
func ColorPicker(bounds rect2.Float32, text string, color colorex.RGBA) colorex.RGBA {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	var ccolor C.struct_Color
	ccolor.r = C.uchar(color.R)
	ccolor.g = C.uchar(color.G)
	ccolor.b = C.uchar(color.B)
	ccolor.a = C.uchar(color.A)
	C.GuiColorPicker(*cbounds, ctext, &ccolor)
	var goRes colorex.RGBA
	goRes.A = byte(ccolor.a)
	goRes.R = byte(ccolor.r)
	goRes.G = byte(ccolor.g)
	goRes.B = byte(ccolor.b)
	return goRes
}

// ColorPanel control
func ColorPanel(bounds rect2.Float32, text string, color colorex.RGBA) colorex.RGBA {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	var ccolor C.struct_Color
	ccolor.b = C.uchar(color.B)
	ccolor.a = C.uchar(color.A)
	ccolor.r = C.uchar(color.R)
	ccolor.g = C.uchar(color.G)
	C.GuiColorPanel(*cbounds, ctext, &ccolor)
	var goRes colorex.RGBA
	goRes.A = byte(ccolor.a)
	goRes.R = byte(ccolor.r)
	goRes.G = byte(ccolor.g)
	goRes.B = byte(ccolor.b)
	return goRes
}

// ColorBarAlpha control
func ColorBarAlpha(bounds rect2.Float32, text string, alpha float32) float32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	calpha := C.float(alpha)
	C.GuiColorBarAlpha(*cbounds, ctext, &calpha)
	return float32(calpha)
}

// ColorBarHue control
func ColorBarHue(bounds rect2.Float32, text string, value float32) float32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	cvalue := C.float(value)
	C.GuiColorBarHue(*cbounds, ctext, &cvalue)
	return float32(cvalue)
}

// ColorPickerHSV - Color Picker control that avoids conversion to RGB on each call (multiple color controls)
func ColorPickerHSV(bounds rect2.Float32, text string, colorHSV *vector3.Float32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	ccolorHSV := cvec3ptr(colorHSV)

	return int32(C.GuiColorPickerHSV(*cbounds, ctext, ccolorHSV))
}

// ColorPanelHSV - Color Panel control that returns HSV color value, used by GuiColorPickerHSV()
func ColorPanelHSV(bounds rect2.Float32, text string, colorHSV *vector3.Float32) int32 {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)
	ccolorHSV := cvec3ptr(colorHSV)

	return int32(C.GuiColorPanelHSV(*cbounds, ctext, ccolorHSV))
}

// DropdownBox control
func DropdownBox(bounds rect2.Float32, text string, active *int32, editMode bool) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)

	if active == nil {
		active = new(int32)
	}
	cactive := C.int(*active)
	defer func() {
		*active = int32(cactive)
	}()

	ceditMode := C.bool(editMode)

	return C.GuiDropdownBox(*cbounds, ctext, &cactive, ceditMode) != 0
}

type DropdownBoxState[T any] struct {
	items           []T
	citemNames      *C.char
	activeItemIndex int32
	edit            bool
}

func (s *DropdownBoxState[T]) Free() {
	if s.citemNames != nil {
		C.free(unsafe.Pointer(s.citemNames))
	}
}

func (s *DropdownBoxState[T]) SetItems(items []T, nameFn func(i T) string) {
	s.items = items
	itemNames := ""
	for i, v := range s.items {
		if i != 0 {
			itemNames += ";"
		}
		itemNames += nameFn(v)
	}

	s.Free()
	s.citemNames = C.CString(itemNames)
}

func (s *DropdownBoxState[T]) IsActiveItem() bool {
	return s.activeItemIndex != -1
}

func (s *DropdownBoxState[T]) ActiveItem() T {
	return s.items[s.activeItemIndex]
}

func (s *DropdownBoxState[T]) SetActiveItem(fn func(T) bool) {
	s.activeItemIndex = (int32)(slices.IndexFunc(s.items, fn))
}

func (s *DropdownBoxState[T]) SetActiveItemIndex(i int32) {
	s.activeItemIndex = i
}

func (s *DropdownBoxState[T]) DropdownBox(bounds rect2.Float32) bool {
	cbounds := crect2ptr(&bounds)
	cactive := (*C.int)(&s.activeItemIndex)
	cedit := (C.bool)(s.edit)

	if C.GuiDropdownBox(*cbounds, s.citemNames, cactive, cedit) != 0 {
		s.edit = !s.edit

		return !s.edit
	}

	return false
}

// ValueBox control, updates input text with numbers
func ValueBox(bounds rect2.Float32, text string, value *int32, minValue, maxValue int, editMode bool) bool {
	cbounds := crect2ptr(&bounds)
	ctext := textAlloc(text)

	if value == nil {
		value = new(int32)
	}
	cvalue := C.int(*value)
	defer func() {
		*value = int32(cvalue)
	}()

	cminValue := C.int(minValue)
	cmaxValue := C.int(maxValue)
	ceditMode := C.bool(editMode)

	return C.GuiValueBox(*cbounds, ctext, &cvalue, cminValue, cmaxValue, ceditMode) != 0
}

// TextBox control, updates input text
func TextBox(bounds rect2.Float32, text *string, textSize int, editMode bool) bool {
	cbounds := crect2ptr(&bounds)

	bs := []byte(*text)
	if len(bs) == 0 {
		bs = []byte{byte(0)}
	}
	if 0 < len(bs) && bs[len(bs)-1] != byte(0) { // minimalize allocation
		bs = append(bs, byte(0)) // for next input symbols
	}
	ctext := (*C.char)(unsafe.Pointer(&bs[0]))
	defer func() {
		*text = strings.TrimSpace(strings.Trim(string(bs), "\x00"))
		// no need : C.free(unsafe.Pointer(ctext))
	}()

	ctextSize := C.int(textSize)
	ceditMode := C.bool(editMode)

	return C.GuiTextBox(*cbounds, ctext, ctextSize, ceditMode) != 0
}

// LoadStyle file over global style variable (.rgs)
func LoadStyle(fileName string) {
	cfileName := textAlloc(fileName)
	C.GuiLoadStyle(cfileName)
}

// LoadStyleDefault over global style
func LoadStyleDefault() {
	C.GuiLoadStyleDefault()
}

// LoadStyleFromMemory - Load style from memory (binary only)
func LoadStyleFromMemory(data []byte) {
	C.GuiLoadStyleFromMemory((*C.uchar)(unsafe.Pointer(&data[0])), C.int(len(data)))
}

// IconText gets text with icon id prepended (if supported)
func IconText(iconId int32, text string) string {
	ciconId := C.int(iconId)
	ctext := textAlloc(text)
	return C.GoString(C.GuiIconText(ciconId, ctext))
}

// Icons enumeration
const (
	ICON_NONE int32 = iota
	ICON_FOLDER_FILE_OPEN
	ICON_FILE_SAVE_CLASSIC
	ICON_FOLDER_OPEN
	ICON_FOLDER_SAVE
	ICON_FILE_OPEN
	ICON_FILE_SAVE
	ICON_FILE_EXPORT
	ICON_FILE_ADD
	ICON_FILE_DELETE
	ICON_FILETYPE_TEXT
	ICON_FILETYPE_AUDIO
	ICON_FILETYPE_IMAGE
	ICON_FILETYPE_PLAY
	ICON_FILETYPE_VIDEO
	ICON_FILETYPE_INFO
	ICON_FILE_COPY
	ICON_FILE_CUT
	ICON_FILE_PASTE
	ICON_CURSOR_HAND
	ICON_CURSOR_POINTER
	ICON_CURSOR_CLASSIC
	ICON_PENCIL
	ICON_PENCIL_BIG
	ICON_BRUSH_CLASSIC
	ICON_BRUSH_PAINTER
	ICON_WATER_DROP
	ICON_COLOR_PICKER
	ICON_RUBBER
	ICON_COLOR_BUCKET
	ICON_TEXT_T
	ICON_TEXT_A
	ICON_SCALE
	ICON_RESIZE
	ICON_FILTER_POINT
	ICON_FILTER_BILINEAR
	ICON_CROP
	ICON_CROP_ALPHA
	ICON_SQUARE_TOGGLE
	ICON_SYMMETRY
	ICON_SYMMETRY_HORIZONTAL
	ICON_SYMMETRY_VERTICAL
	ICON_LENS
	ICON_LENS_BIG
	ICON_EYE_ON
	ICON_EYE_OFF
	ICON_FILTER_TOP
	ICON_FILTER
	ICON_TARGET_POINT
	ICON_TARGET_SMALL
	ICON_TARGET_BIG
	ICON_TARGET_MOVE
	ICON_CURSOR_MOVE
	ICON_CURSOR_SCALE
	ICON_CURSOR_SCALE_RIGHT
	ICON_CURSOR_SCALE_LEFT
	ICON_UNDO
	ICON_REDO
	ICON_REREDO
	ICON_MUTATE
	ICON_ROTATE
	ICON_REPEAT
	ICON_SHUFFLE
	ICON_EMPTYBOX
	ICON_TARGET
	ICON_TARGET_SMALL_FILL
	ICON_TARGET_BIG_FILL
	ICON_TARGET_MOVE_FILL
	ICON_CURSOR_MOVE_FILL
	ICON_CURSOR_SCALE_FILL
	ICON_CURSOR_SCALE_RIGHT_FILL
	ICON_CURSOR_SCALE_LEFT_FILL
	ICON_UNDO_FILL
	ICON_REDO_FILL
	ICON_REREDO_FILL
	ICON_MUTATE_FILL
	ICON_ROTATE_FILL
	ICON_REPEAT_FILL
	ICON_SHUFFLE_FILL
	ICON_EMPTYBOX_SMALL
	ICON_BOX
	ICON_BOX_TOP
	ICON_BOX_TOP_RIGHT
	ICON_BOX_RIGHT
	ICON_BOX_BOTTOM_RIGHT
	ICON_BOX_BOTTOM
	ICON_BOX_BOTTOM_LEFT
	ICON_BOX_LEFT
	ICON_BOX_TOP_LEFT
	ICON_BOX_CENTER
	ICON_BOX_CIRCLE_MASK
	ICON_POT
	ICON_ALPHA_MULTIPLY
	ICON_ALPHA_CLEAR
	ICON_DITHERING
	ICON_MIPMAPS
	ICON_BOX_GRID
	ICON_GRID
	ICON_BOX_CORNERS_SMALL
	ICON_BOX_CORNERS_BIG
	ICON_FOUR_BOXES
	ICON_GRID_FILL
	ICON_BOX_MULTISIZE
	ICON_ZOOM_SMALL
	ICON_ZOOM_MEDIUM
	ICON_ZOOM_BIG
	ICON_ZOOM_ALL
	ICON_ZOOM_CENTER
	ICON_BOX_DOTS_SMALL
	ICON_BOX_DOTS_BIG
	ICON_BOX_CONCENTRIC
	ICON_BOX_GRID_BIG
	ICON_OK_TICK
	ICON_CROSS
	ICON_ARROW_LEFT
	ICON_ARROW_RIGHT
	ICON_ARROW_DOWN
	ICON_ARROW_UP
	ICON_ARROW_LEFT_FILL
	ICON_ARROW_RIGHT_FILL
	ICON_ARROW_DOWN_FILL
	ICON_ARROW_UP_FILL
	ICON_AUDIO
	ICON_FX
	ICON_WAVE
	ICON_WAVE_SINUS
	ICON_WAVE_SQUARE
	ICON_WAVE_TRIANGULAR
	ICON_CROSS_SMALL
	ICON_PLAYER_PREVIOUS
	ICON_PLAYER_PLAY_BACK
	ICON_PLAYER_PLAY
	ICON_PLAYER_PAUSE
	ICON_PLAYER_STOP
	ICON_PLAYER_NEXT
	ICON_PLAYER_RECORD
	ICON_MAGNET
	ICON_LOCK_CLOSE
	ICON_LOCK_OPEN
	ICON_CLOCK
	ICON_TOOLS
	ICON_GEAR
	ICON_GEAR_BIG
	ICON_BIN
	ICON_HAND_POINTER
	ICON_LASER
	ICON_COIN
	ICON_EXPLOSION
	ICON_1UP
	ICON_PLAYER
	ICON_PLAYER_JUMP
	ICON_KEY
	ICON_DEMON
	ICON_TEXT_POPUP
	ICON_GEAR_EX
	ICON_CRACK
	ICON_CRACK_POINTS
	ICON_STAR
	ICON_DOOR
	ICON_EXIT
	ICON_MODE_2D
	ICON_MODE_3D
	ICON_CUBE
	ICON_CUBE_FACE_TOP
	ICON_CUBE_FACE_LEFT
	ICON_CUBE_FACE_FRONT
	ICON_CUBE_FACE_BOTTOM
	ICON_CUBE_FACE_RIGHT
	ICON_CUBE_FACE_BACK
	ICON_CAMERA
	ICON_SPECIAL
	ICON_LINK_NET
	ICON_LINK_BOXES
	ICON_LINK_MULTI
	ICON_LINK
	ICON_LINK_BROKE
	ICON_TEXT_NOTES
	ICON_NOTEBOOK
	ICON_SUITCASE
	ICON_SUITCASE_ZIP
	ICON_MAILBOX
	ICON_MONITOR
	ICON_PRINTER
	ICON_PHOTO_CAMERA
	ICON_PHOTO_CAMERA_FLASH
	ICON_HOUSE
	ICON_HEART
	ICON_CORNER
	ICON_VERTICAL_BARS
	ICON_VERTICAL_BARS_FILL
	ICON_LIFE_BARS
	ICON_INFO
	ICON_CROSSLINE
	ICON_HELP
	ICON_FILETYPE_ALPHA
	ICON_FILETYPE_HOME
	ICON_LAYERS_VISIBLE
	ICON_LAYERS
	ICON_WINDOW
	ICON_HIDPI
	ICON_FILETYPE_BINARY
	ICON_HEX
	ICON_SHIELD
	ICON_FILE_NEW
	ICON_FOLDER_ADD
	ICON_ALARM
	ICON_CPU
	ICON_ROM
	ICON_STEP_OVER
	ICON_STEP_INTO
	ICON_STEP_OUT
	ICON_RESTART
	ICON_BREAKPOINT_ON
	ICON_BREAKPOINT_OFF
	ICON_BURGER_MENU
	ICON_CASE_SENSITIVE
	ICON_REG_EXP
	ICON_FOLDER
	ICON_FILE
	ICON_SAND_TIMER
	ICON_220
	ICON_221
	ICON_222
	ICON_223
	ICON_224
	ICON_225
	ICON_226
	ICON_227
	ICON_228
	ICON_229
	ICON_230
	ICON_231
	ICON_232
	ICON_233
	ICON_234
	ICON_235
	ICON_236
	ICON_237
	ICON_238
	ICON_239
	ICON_240
	ICON_241
	ICON_242
	ICON_243
	ICON_244
	ICON_245
	ICON_246
	ICON_247
	ICON_248
	ICON_249
	ICON_250
	ICON_251
	ICON_252
	ICON_253
	ICON_254
	ICON_255
)

// TextInputBox control, ask for text
func TextInputBox(bounds rect2.Float32, title, message, buttons string, text *string, textMaxSize int32, secretViewActive *bool) int32 {
	cbounds := crect2ptr(&bounds)

	ctitle := textAlloc(title)
	cmessage := textAlloc(message)
	cbuttons := textAlloc(buttons)

	bs := []byte(*text)
	if len(bs) == 0 {
		bs = []byte{byte(0)}
	}
	if 0 < len(bs) && bs[len(bs)-1] != byte(0) { // minimalize allocation
		bs = append(bs, byte(0)) // for next input symbols
	}
	ctext := (*C.char)(unsafe.Pointer(&bs[0]))
	defer func() {
		*text = strings.TrimSpace(strings.Trim(string(bs), "\x00"))
		// no need : C.free(unsafe.Pointer(ctext))
	}()

	ctextMaxSize := C.int(textMaxSize)

	csecretViewActive := C.bool(*secretViewActive)
	defer func() {
		*secretViewActive = bool(csecretViewActive)
	}()

	return int32(C.GuiTextInputBox(*cbounds, ctitle, cmessage, cbuttons, ctext, ctextMaxSize, &csecretViewActive))
}

// ListViewEx control with extended parameters
func ListViewEx(bounds rect2.Float32, text []string, focus, scrollIndex *int32, active int32) int32 {
	cbounds := crect2ptr(&bounds)

	ctext := NewCStringArrayFromSlice(text)
	defer ctext.Free()

	count := C.int(len(text))

	if focus == nil {
		focus = new(int32)
	}
	cfocus := C.int(*focus)
	defer func() {
		*focus = int32(cfocus)
	}()

	if scrollIndex == nil {
		scrollIndex = new(int32)
	}
	cscrollIndex := C.int(*scrollIndex)
	defer func() {
		*scrollIndex = int32(cscrollIndex)
	}()

	cactive := C.int(active)

	C.GuiListViewEx(*cbounds, (**C.char)(ctext.Pointer), count, &cscrollIndex, &cactive, &cfocus)
	return int32(cactive)
}

type ListViewState[T any] struct {
	items           []T
	citemNames      CStringArray
	focusIndex      int32
	scrollIndex     int32
	activeItemIndex int32
}

func (s *ListViewState[T]) Free() {
	if !s.citemNames.IsNil() {
		s.citemNames.Free()
	}
}

func (s *ListViewState[T]) SetItems(items []T, nameFn func(i T) string) {
	s.items = items
	itemNames := make([]string, len(s.items))
	for i, v := range s.items {
		itemNames[i] = nameFn(v)
	}

	s.Free()
	s.citemNames = NewCStringArrayFromSlice(itemNames)
}

func (s *ListViewState[T]) IsActiveItem() bool {
	return s.activeItemIndex != -1
}

func (s *ListViewState[T]) ActiveItem() T {
	return s.items[s.activeItemIndex]
}

func (s *ListViewState[T]) ListView(bounds rect2.Float32) {
	cbounds := crect2ptr(&bounds)
	count := C.int(len(s.items))
	cfocus := (*C.int)(&s.focusIndex)
	cscrollIndex := (*C.int)(&s.scrollIndex)
	cactive := (*C.int)(&s.activeItemIndex)

	C.GuiListViewEx(*cbounds, (**C.char)(s.citemNames.Pointer), count, cscrollIndex, cactive, cfocus)
}

// TabBar control
func TabBar(bounds rect2.Float32, text []string, active *int32) int32 {
	cbounds := crect2ptr(&bounds)

	ctext := NewCStringArrayFromSlice(text)
	defer ctext.Free()

	count := C.int(len(text))

	if active == nil {
		active = new(int32)
	}
	cactive := C.int(*active)
	defer func() {
		*active = int32(cactive)
	}()
	return int32(C.GuiTabBar(*cbounds, (**C.char)(ctext.Pointer), count, &cactive))
}

// SetFont - set custom font (global state)
func SetFont(font rl.Font) {
	cfont := (*C.Font)(unsafe.Pointer(&font))
	C.GuiSetFont(*cfont)
}

// GetFont - get custom font (global state)
func GetFont() rl.Font {
	ret := C.GuiGetFont()
	ptr := unsafe.Pointer(&ret)
	return *(*rl.Font)(ptr)
}

// SetAlpha - set alpha (global state)
func SetAlpha(alpha float32) {
	C.GuiSetAlpha((C.float)(alpha))
}
