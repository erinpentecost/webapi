// Code generated by webidl-bind. DO NOT EDIT.

package pointerevents

import "syscall/js"

import (
	"github.com/gowebapi/webapi"
	"github.com/gowebapi/webapi/device/inputcapabilities"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/html/htmlevent"
)

// using following types:
// domcore.EventTarget
// htmlevent.MouseEvent
// inputcapabilities.InputDeviceCapabilities
// webapi.Window

// source idl files:
// pointerevents.idl

// transform files:
// pointerevents.go.md

// ReleasableApiResource is used to release underlaying
// allocated resources.
type ReleasableApiResource interface {
	Release()
}

type releasableApiResourceList []ReleasableApiResource

func (a releasableApiResourceList) Release() {
	for _, v := range a {
		v.Release()
	}
}

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// dictionary: PointerEventInit
type PointerEventInit struct {
	Bubbles            bool
	Cancelable         bool
	Composed           bool
	View               *webapi.Window
	Detail             int
	SourceCapabilities *inputcapabilities.InputDeviceCapabilities
	CtrlKey            bool
	ShiftKey           bool
	AltKey             bool
	MetaKey            bool
	ModifierAltGraph   bool
	ModifierCapsLock   bool
	ModifierFn         bool
	ModifierFnLock     bool
	ModifierHyper      bool
	ModifierNumLock    bool
	ModifierScrollLock bool
	ModifierSuper      bool
	ModifierSymbol     bool
	ModifierSymbolLock bool
	Button             int
	Buttons            int
	RelatedTarget      *domcore.EventTarget
	ScreenX            float64
	ScreenY            float64
	ClientX            float64
	ClientY            float64
	MovementX          int
	MovementY          int
	PointerId          int
	Width              float64
	Height             float64
	Pressure           float32
	TangentialPressure float32
	TiltX              int
	TiltY              int
	Twist              int
	PointerType        string
	IsPrimary          bool
	CoalescedEvents    []*PointerEvent
	PredictedEvents    []*PointerEvent
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PointerEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.View.JSValue()
	out.Set("view", value3)
	value4 := _this.Detail
	out.Set("detail", value4)
	value5 := _this.SourceCapabilities.JSValue()
	out.Set("sourceCapabilities", value5)
	value6 := _this.CtrlKey
	out.Set("ctrlKey", value6)
	value7 := _this.ShiftKey
	out.Set("shiftKey", value7)
	value8 := _this.AltKey
	out.Set("altKey", value8)
	value9 := _this.MetaKey
	out.Set("metaKey", value9)
	value10 := _this.ModifierAltGraph
	out.Set("modifierAltGraph", value10)
	value11 := _this.ModifierCapsLock
	out.Set("modifierCapsLock", value11)
	value12 := _this.ModifierFn
	out.Set("modifierFn", value12)
	value13 := _this.ModifierFnLock
	out.Set("modifierFnLock", value13)
	value14 := _this.ModifierHyper
	out.Set("modifierHyper", value14)
	value15 := _this.ModifierNumLock
	out.Set("modifierNumLock", value15)
	value16 := _this.ModifierScrollLock
	out.Set("modifierScrollLock", value16)
	value17 := _this.ModifierSuper
	out.Set("modifierSuper", value17)
	value18 := _this.ModifierSymbol
	out.Set("modifierSymbol", value18)
	value19 := _this.ModifierSymbolLock
	out.Set("modifierSymbolLock", value19)
	value20 := _this.Button
	out.Set("button", value20)
	value21 := _this.Buttons
	out.Set("buttons", value21)
	value22 := _this.RelatedTarget.JSValue()
	out.Set("relatedTarget", value22)
	value23 := _this.ScreenX
	out.Set("screenX", value23)
	value24 := _this.ScreenY
	out.Set("screenY", value24)
	value25 := _this.ClientX
	out.Set("clientX", value25)
	value26 := _this.ClientY
	out.Set("clientY", value26)
	value27 := _this.MovementX
	out.Set("movementX", value27)
	value28 := _this.MovementY
	out.Set("movementY", value28)
	value29 := _this.PointerId
	out.Set("pointerId", value29)
	value30 := _this.Width
	out.Set("width", value30)
	value31 := _this.Height
	out.Set("height", value31)
	value32 := _this.Pressure
	out.Set("pressure", value32)
	value33 := _this.TangentialPressure
	out.Set("tangentialPressure", value33)
	value34 := _this.TiltX
	out.Set("tiltX", value34)
	value35 := _this.TiltY
	out.Set("tiltY", value35)
	value36 := _this.Twist
	out.Set("twist", value36)
	value37 := _this.PointerType
	out.Set("pointerType", value37)
	value38 := _this.IsPrimary
	out.Set("isPrimary", value38)
	value39 := js.Global().Get("Array").New(len(_this.CoalescedEvents))
	for __idx39, __seq_in39 := range _this.CoalescedEvents {
		__seq_out39 := __seq_in39.JSValue()
		value39.SetIndex(__idx39, __seq_out39)
	}
	out.Set("coalescedEvents", value39)
	value40 := js.Global().Get("Array").New(len(_this.PredictedEvents))
	for __idx40, __seq_in40 := range _this.PredictedEvents {
		__seq_out40 := __seq_in40.JSValue()
		value40.SetIndex(__idx40, __seq_out40)
	}
	out.Set("predictedEvents", value40)
	return out
}

// PointerEventInitFromJS is allocating a new
// PointerEventInit object and copy all values from
// input javascript object
func PointerEventInitFromJS(value js.Wrapper) *PointerEventInit {
	input := value.JSValue()
	var out PointerEventInit
	var (
		value0  bool                                       // javascript: boolean {bubbles Bubbles bubbles}
		value1  bool                                       // javascript: boolean {cancelable Cancelable cancelable}
		value2  bool                                       // javascript: boolean {composed Composed composed}
		value3  *webapi.Window                             // javascript: Window {view View view}
		value4  int                                        // javascript: long {detail Detail detail}
		value5  *inputcapabilities.InputDeviceCapabilities // javascript: InputDeviceCapabilities {sourceCapabilities SourceCapabilities sourceCapabilities}
		value6  bool                                       // javascript: boolean {ctrlKey CtrlKey ctrlKey}
		value7  bool                                       // javascript: boolean {shiftKey ShiftKey shiftKey}
		value8  bool                                       // javascript: boolean {altKey AltKey altKey}
		value9  bool                                       // javascript: boolean {metaKey MetaKey metaKey}
		value10 bool                                       // javascript: boolean {modifierAltGraph ModifierAltGraph modifierAltGraph}
		value11 bool                                       // javascript: boolean {modifierCapsLock ModifierCapsLock modifierCapsLock}
		value12 bool                                       // javascript: boolean {modifierFn ModifierFn modifierFn}
		value13 bool                                       // javascript: boolean {modifierFnLock ModifierFnLock modifierFnLock}
		value14 bool                                       // javascript: boolean {modifierHyper ModifierHyper modifierHyper}
		value15 bool                                       // javascript: boolean {modifierNumLock ModifierNumLock modifierNumLock}
		value16 bool                                       // javascript: boolean {modifierScrollLock ModifierScrollLock modifierScrollLock}
		value17 bool                                       // javascript: boolean {modifierSuper ModifierSuper modifierSuper}
		value18 bool                                       // javascript: boolean {modifierSymbol ModifierSymbol modifierSymbol}
		value19 bool                                       // javascript: boolean {modifierSymbolLock ModifierSymbolLock modifierSymbolLock}
		value20 int                                        // javascript: short {button Button button}
		value21 int                                        // javascript: unsigned short {buttons Buttons buttons}
		value22 *domcore.EventTarget                       // javascript: EventTarget {relatedTarget RelatedTarget relatedTarget}
		value23 float64                                    // javascript: double {screenX ScreenX screenX}
		value24 float64                                    // javascript: double {screenY ScreenY screenY}
		value25 float64                                    // javascript: double {clientX ClientX clientX}
		value26 float64                                    // javascript: double {clientY ClientY clientY}
		value27 int                                        // javascript: long {movementX MovementX movementX}
		value28 int                                        // javascript: long {movementY MovementY movementY}
		value29 int                                        // javascript: long {pointerId PointerId pointerId}
		value30 float64                                    // javascript: double {width Width width}
		value31 float64                                    // javascript: double {height Height height}
		value32 float32                                    // javascript: float {pressure Pressure pressure}
		value33 float32                                    // javascript: float {tangentialPressure TangentialPressure tangentialPressure}
		value34 int                                        // javascript: long {tiltX TiltX tiltX}
		value35 int                                        // javascript: long {tiltY TiltY tiltY}
		value36 int                                        // javascript: long {twist Twist twist}
		value37 string                                     // javascript: DOMString {pointerType PointerType pointerType}
		value38 bool                                       // javascript: boolean {isPrimary IsPrimary isPrimary}
		value39 []*PointerEvent                            // javascript: sequence<PointerEvent> {coalescedEvents CoalescedEvents coalescedEvents}
		value40 []*PointerEvent                            // javascript: sequence<PointerEvent> {predictedEvents PredictedEvents predictedEvents}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	if input.Get("view").Type() != js.TypeNull && input.Get("view").Type() != js.TypeUndefined {
		value3 = webapi.WindowFromJS(input.Get("view"))
	}
	out.View = value3
	value4 = (input.Get("detail")).Int()
	out.Detail = value4
	if input.Get("sourceCapabilities").Type() != js.TypeNull && input.Get("sourceCapabilities").Type() != js.TypeUndefined {
		value5 = inputcapabilities.InputDeviceCapabilitiesFromJS(input.Get("sourceCapabilities"))
	}
	out.SourceCapabilities = value5
	value6 = (input.Get("ctrlKey")).Bool()
	out.CtrlKey = value6
	value7 = (input.Get("shiftKey")).Bool()
	out.ShiftKey = value7
	value8 = (input.Get("altKey")).Bool()
	out.AltKey = value8
	value9 = (input.Get("metaKey")).Bool()
	out.MetaKey = value9
	value10 = (input.Get("modifierAltGraph")).Bool()
	out.ModifierAltGraph = value10
	value11 = (input.Get("modifierCapsLock")).Bool()
	out.ModifierCapsLock = value11
	value12 = (input.Get("modifierFn")).Bool()
	out.ModifierFn = value12
	value13 = (input.Get("modifierFnLock")).Bool()
	out.ModifierFnLock = value13
	value14 = (input.Get("modifierHyper")).Bool()
	out.ModifierHyper = value14
	value15 = (input.Get("modifierNumLock")).Bool()
	out.ModifierNumLock = value15
	value16 = (input.Get("modifierScrollLock")).Bool()
	out.ModifierScrollLock = value16
	value17 = (input.Get("modifierSuper")).Bool()
	out.ModifierSuper = value17
	value18 = (input.Get("modifierSymbol")).Bool()
	out.ModifierSymbol = value18
	value19 = (input.Get("modifierSymbolLock")).Bool()
	out.ModifierSymbolLock = value19
	value20 = (input.Get("button")).Int()
	out.Button = value20
	value21 = (input.Get("buttons")).Int()
	out.Buttons = value21
	if input.Get("relatedTarget").Type() != js.TypeNull && input.Get("relatedTarget").Type() != js.TypeUndefined {
		value22 = domcore.EventTargetFromJS(input.Get("relatedTarget"))
	}
	out.RelatedTarget = value22
	value23 = (input.Get("screenX")).Float()
	out.ScreenX = value23
	value24 = (input.Get("screenY")).Float()
	out.ScreenY = value24
	value25 = (input.Get("clientX")).Float()
	out.ClientX = value25
	value26 = (input.Get("clientY")).Float()
	out.ClientY = value26
	value27 = (input.Get("movementX")).Int()
	out.MovementX = value27
	value28 = (input.Get("movementY")).Int()
	out.MovementY = value28
	value29 = (input.Get("pointerId")).Int()
	out.PointerId = value29
	value30 = (input.Get("width")).Float()
	out.Width = value30
	value31 = (input.Get("height")).Float()
	out.Height = value31
	value32 = (float32)((input.Get("pressure")).Float())
	out.Pressure = value32
	value33 = (float32)((input.Get("tangentialPressure")).Float())
	out.TangentialPressure = value33
	value34 = (input.Get("tiltX")).Int()
	out.TiltX = value34
	value35 = (input.Get("tiltY")).Int()
	out.TiltY = value35
	value36 = (input.Get("twist")).Int()
	out.Twist = value36
	value37 = (input.Get("pointerType")).String()
	out.PointerType = value37
	value38 = (input.Get("isPrimary")).Bool()
	out.IsPrimary = value38
	__length39 := input.Get("coalescedEvents").Length()
	__array39 := make([]*PointerEvent, __length39, __length39)
	for __idx39 := 0; __idx39 < __length39; __idx39++ {
		var __seq_out39 *PointerEvent
		__seq_in39 := input.Get("coalescedEvents").Index(__idx39)
		__seq_out39 = PointerEventFromJS(__seq_in39)
		__array39[__idx39] = __seq_out39
	}
	value39 = __array39
	out.CoalescedEvents = value39
	__length40 := input.Get("predictedEvents").Length()
	__array40 := make([]*PointerEvent, __length40, __length40)
	for __idx40 := 0; __idx40 < __length40; __idx40++ {
		var __seq_out40 *PointerEvent
		__seq_in40 := input.Get("predictedEvents").Index(__idx40)
		__seq_out40 = PointerEventFromJS(__seq_in40)
		__array40[__idx40] = __seq_out40
	}
	value40 = __array40
	out.PredictedEvents = value40
	return &out
}

// interface: PointerEvent
type PointerEvent struct {
	htmlevent.MouseEvent
}

// PointerEventFromJS is casting a js.Wrapper into PointerEvent.
func PointerEventFromJS(value js.Wrapper) *PointerEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PointerEvent{}
	ret.Value_JS = input
	return ret
}

func NewPointerEvent(_type string, eventInitDict *PointerEventInit) (_result *PointerEvent) {
	_klass := js.Global().Get("PointerEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *PointerEvent // javascript: PointerEvent _what_return_name
	)
	_converted = PointerEventFromJS(_returned)
	_result = _converted
	return
}

// PointerId returning attribute 'pointerId' with
// type int (idl: long).
func (_this *PointerEvent) PointerId() int {
	var ret int
	value := _this.Value_JS.Get("pointerId")
	ret = (value).Int()
	return ret
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *PointerEvent) Width() float64 {
	var ret float64
	value := _this.Value_JS.Get("width")
	ret = (value).Float()
	return ret
}

// Height returning attribute 'height' with
// type float64 (idl: double).
func (_this *PointerEvent) Height() float64 {
	var ret float64
	value := _this.Value_JS.Get("height")
	ret = (value).Float()
	return ret
}

// Pressure returning attribute 'pressure' with
// type float32 (idl: float).
func (_this *PointerEvent) Pressure() float32 {
	var ret float32
	value := _this.Value_JS.Get("pressure")
	ret = (float32)((value).Float())
	return ret
}

// TangentialPressure returning attribute 'tangentialPressure' with
// type float32 (idl: float).
func (_this *PointerEvent) TangentialPressure() float32 {
	var ret float32
	value := _this.Value_JS.Get("tangentialPressure")
	ret = (float32)((value).Float())
	return ret
}

// TiltX returning attribute 'tiltX' with
// type int (idl: long).
func (_this *PointerEvent) TiltX() int {
	var ret int
	value := _this.Value_JS.Get("tiltX")
	ret = (value).Int()
	return ret
}

// TiltY returning attribute 'tiltY' with
// type int (idl: long).
func (_this *PointerEvent) TiltY() int {
	var ret int
	value := _this.Value_JS.Get("tiltY")
	ret = (value).Int()
	return ret
}

// Twist returning attribute 'twist' with
// type int (idl: long).
func (_this *PointerEvent) Twist() int {
	var ret int
	value := _this.Value_JS.Get("twist")
	ret = (value).Int()
	return ret
}

// PointerType returning attribute 'pointerType' with
// type string (idl: DOMString).
func (_this *PointerEvent) PointerType() string {
	var ret string
	value := _this.Value_JS.Get("pointerType")
	ret = (value).String()
	return ret
}

// IsPrimary returning attribute 'isPrimary' with
// type bool (idl: boolean).
func (_this *PointerEvent) IsPrimary() bool {
	var ret bool
	value := _this.Value_JS.Get("isPrimary")
	ret = (value).Bool()
	return ret
}

func (_this *PointerEvent) GetCoalescedEvents() (_result []*PointerEvent) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getCoalescedEvents", _args[0:_end]...)
	var (
		_converted []*PointerEvent // javascript: sequence<PointerEvent> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*PointerEvent, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *PointerEvent
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = PointerEventFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *PointerEvent) GetPredictedEvents() (_result []*PointerEvent) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getPredictedEvents", _args[0:_end]...)
	var (
		_converted []*PointerEvent // javascript: sequence<PointerEvent> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*PointerEvent, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *PointerEvent
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = PointerEventFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}
