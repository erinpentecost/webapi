// Code generated by webidl-bind. DO NOT EDIT.

package notification

import "syscall/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.EventHandler
// domcore.EventTarget
// javascript.FrozenArray
// javascript.PromiseFinally

// source idl files:
// notifications.idl
// promises.idl

// transform files:
// notifications.go.md
// promises.go.md

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

// enum: NotificationDirection
type Direction int

const (
	Auto Direction = iota
	LeftToRight
	RightToLeft
)

var notificationDirectionToWasmTable = []string{
	"auto", "ltr", "rtl",
}

var notificationDirectionFromWasmTable = map[string]Direction{
	"auto": Auto, "ltr": LeftToRight, "rtl": RightToLeft,
}

// JSValue is converting this enum into a java object
func (this *Direction) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this Direction) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(notificationDirectionToWasmTable) {
		return notificationDirectionToWasmTable[idx]
	}
	panic("unknown input value")
}

// DirectionFromJS is converting a javascript value into
// a Direction enum value.
func DirectionFromJS(value js.Value) Direction {
	key := value.String()
	conv, ok := notificationDirectionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: NotificationPermission
type PermissionMode int

const (
	Default PermissionMode = iota
	Denied
	Granted
)

var notificationPermissionToWasmTable = []string{
	"default", "denied", "granted",
}

var notificationPermissionFromWasmTable = map[string]PermissionMode{
	"default": Default, "denied": Denied, "granted": Granted,
}

// JSValue is converting this enum into a java object
func (this *PermissionMode) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this PermissionMode) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(notificationPermissionToWasmTable) {
		return notificationPermissionToWasmTable[idx]
	}
	panic("unknown input value")
}

// PermissionModeFromJS is converting a javascript value into
// a PermissionMode enum value.
func PermissionModeFromJS(value js.Value) PermissionMode {
	key := value.String()
	conv, ok := notificationPermissionFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: NotificationPermissionCallback
type PermissionCallbackFunc func(permission PermissionMode)

// PermissionCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PermissionCallback js.Func

func PermissionCallbackToJS(callback PermissionCallbackFunc) *PermissionCallback {
	if callback == nil {
		return nil
	}
	ret := PermissionCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 PermissionMode // javascript: NotificationPermission permission
		)
		_p0 = PermissionModeFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PermissionCallbackFromJS(_value js.Value) PermissionCallbackFunc {
	return func(permission PermissionMode) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := permission.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromisePermissionModeOnFulfilledFunc func(value PermissionMode)

// PromisePermissionModeOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionModeOnFulfilled js.Func

func PromisePermissionModeOnFulfilledToJS(callback PromisePermissionModeOnFulfilledFunc) *PromisePermissionModeOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionModeOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 PermissionMode // javascript: NotificationPermission value
		)
		_p0 = PermissionModeFromJS(args[0])
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePermissionModeOnFulfilledFromJS(_value js.Value) PromisePermissionModeOnFulfilledFunc {
	return func(value PermissionMode) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromisePermissionModeOnRejectedFunc func(reason js.Value)

// PromisePermissionModeOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromisePermissionModeOnRejected js.Func

func PromisePermissionModeOnRejectedToJS(callback PromisePermissionModeOnRejectedFunc) *PromisePermissionModeOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromisePermissionModeOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)
		// returning no return value
		return nil
	}))
	return &ret
}

func PromisePermissionModeOnRejectedFromJS(_value js.Value) PromisePermissionModeOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: NotificationAction
type NotificationAction struct {
	Action string
	Title  string
	Icon   string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *NotificationAction) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Action
	out.Set("action", value0)
	value1 := _this.Title
	out.Set("title", value1)
	value2 := _this.Icon
	out.Set("icon", value2)
	return out
}

// NotificationActionFromJS is allocating a new
// NotificationAction object and copy all values from
// input javascript object
func NotificationActionFromJS(value js.Wrapper) *NotificationAction {
	input := value.JSValue()
	var out NotificationAction
	var (
		value0 string // javascript: DOMString {action Action action}
		value1 string // javascript: DOMString {title Title title}
		value2 string // javascript: USVString {icon Icon icon}
	)
	value0 = (input.Get("action")).String()
	out.Action = value0
	value1 = (input.Get("title")).String()
	out.Title = value1
	value2 = (input.Get("icon")).String()
	out.Icon = value2
	return &out
}

// dictionary: NotificationOptions
type Options struct {
	Dir                Direction
	Lang               string
	Body               string
	Tag                string
	Image              string
	Icon               string
	Badge              string
	Timestamp          int
	Renotify           bool
	Silent             bool
	RequireInteraction bool
	Data               js.Value
	Actions            []*NotificationAction
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *Options) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Dir.JSValue()
	out.Set("dir", value0)
	value1 := _this.Lang
	out.Set("lang", value1)
	value2 := _this.Body
	out.Set("body", value2)
	value3 := _this.Tag
	out.Set("tag", value3)
	value4 := _this.Image
	out.Set("image", value4)
	value5 := _this.Icon
	out.Set("icon", value5)
	value6 := _this.Badge
	out.Set("badge", value6)
	value7 := _this.Timestamp
	out.Set("timestamp", value7)
	value8 := _this.Renotify
	out.Set("renotify", value8)
	value9 := _this.Silent
	out.Set("silent", value9)
	value10 := _this.RequireInteraction
	out.Set("requireInteraction", value10)
	value11 := _this.Data
	out.Set("data", value11)
	value12 := js.Global().Get("Array").New(len(_this.Actions))
	for __idx12, __seq_in12 := range _this.Actions {
		__seq_out12 := __seq_in12.JSValue()
		value12.SetIndex(__idx12, __seq_out12)
	}
	out.Set("actions", value12)
	return out
}

// OptionsFromJS is allocating a new
// Options object and copy all values from
// input javascript object
func OptionsFromJS(value js.Wrapper) *Options {
	input := value.JSValue()
	var out Options
	var (
		value0  Direction             // javascript: NotificationDirection {dir Dir dir}
		value1  string                // javascript: DOMString {lang Lang lang}
		value2  string                // javascript: DOMString {body Body body}
		value3  string                // javascript: DOMString {tag Tag tag}
		value4  string                // javascript: USVString {image Image image}
		value5  string                // javascript: USVString {icon Icon icon}
		value6  string                // javascript: USVString {badge Badge badge}
		value7  int                   // javascript: unsigned long long {timestamp Timestamp timestamp}
		value8  bool                  // javascript: boolean {renotify Renotify renotify}
		value9  bool                  // javascript: boolean {silent Silent silent}
		value10 bool                  // javascript: boolean {requireInteraction RequireInteraction requireInteraction}
		value11 js.Value              // javascript: any {data Data data}
		value12 []*NotificationAction // javascript: sequence<NotificationAction> {actions Actions actions}
	)
	value0 = DirectionFromJS(input.Get("dir"))
	out.Dir = value0
	value1 = (input.Get("lang")).String()
	out.Lang = value1
	value2 = (input.Get("body")).String()
	out.Body = value2
	value3 = (input.Get("tag")).String()
	out.Tag = value3
	value4 = (input.Get("image")).String()
	out.Image = value4
	value5 = (input.Get("icon")).String()
	out.Icon = value5
	value6 = (input.Get("badge")).String()
	out.Badge = value6
	value7 = (input.Get("timestamp")).Int()
	out.Timestamp = value7
	value8 = (input.Get("renotify")).Bool()
	out.Renotify = value8
	value9 = (input.Get("silent")).Bool()
	out.Silent = value9
	value10 = (input.Get("requireInteraction")).Bool()
	out.RequireInteraction = value10
	value11 = input.Get("data")
	out.Data = value11
	__length12 := input.Get("actions").Length()
	__array12 := make([]*NotificationAction, __length12, __length12)
	for __idx12 := 0; __idx12 < __length12; __idx12++ {
		var __seq_out12 *NotificationAction
		__seq_in12 := input.Get("actions").Index(__idx12)
		__seq_out12 = NotificationActionFromJS(__seq_in12)
		__array12[__idx12] = __seq_out12
	}
	value12 = __array12
	out.Actions = value12
	return &out
}

// interface: Notification
type Notification struct {
	domcore.EventTarget
}

// NotificationFromJS is casting a js.Wrapper into Notification.
func NotificationFromJS(value js.Wrapper) *Notification {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Notification{}
	ret.Value_JS = input
	return ret
}

// Permission returning attribute 'permission' with
// type PermissionMode (idl: NotificationPermission).
func Permission() PermissionMode {
	var ret PermissionMode
	_klass := js.Global().Get("Notification")
	value := _klass.Get("permission")
	ret = PermissionModeFromJS(value)
	return ret
}

// MaxActions returning attribute 'maxActions' with
// type uint (idl: unsigned long).
func MaxActions() uint {
	var ret uint
	_klass := js.Global().Get("Notification")
	value := _klass.Get("maxActions")
	ret = (uint)((value).Int())
	return ret
}

func RequestPermission(deprecatedCallback *PermissionCallback) (_result *PromisePermissionMode) {
	_klass := js.Global().Get("Notification")
	_method := _klass.Get("requestPermission")
	var (
		_args [1]interface{}
		_end  int
	)
	if deprecatedCallback != nil {

		var __callback0 js.Value
		if deprecatedCallback != nil {
			__callback0 = (*deprecatedCallback).Value
		} else {
			__callback0 = js.Null()
		}
		_p0 := __callback0
		_args[0] = _p0
		_end++
	}
	_returned := _method.Invoke(_args[0:_end]...)
	var (
		_converted *PromisePermissionMode // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionModeFromJS(_returned)
	_result = _converted
	return
}

func New(title string, options *Options) (_result *Notification) {
	_klass := js.Global().Get("Notification")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := title
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Notification // javascript: Notification _what_return_name
	)
	_converted = NotificationFromJS(_returned)
	_result = _converted
	return
}

// OnClick returning attribute 'onclick' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnClick() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclick")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnClick setting attribute 'onclick' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnClick(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onclick", input)
}

// OnShow returning attribute 'onshow' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnShow() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onshow")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnShow setting attribute 'onshow' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnShow(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onshow", input)
}

// OnError returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnError() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnError setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnError(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onerror", input)
}

// OnClose returning attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) OnClose() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclose")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnClose setting attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *Notification) SetOnClose(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onclose", input)
}

// Title returning attribute 'title' with
// type string (idl: DOMString).
func (_this *Notification) Title() string {
	var ret string
	value := _this.Value_JS.Get("title")
	ret = (value).String()
	return ret
}

// Dir returning attribute 'dir' with
// type Direction (idl: NotificationDirection).
func (_this *Notification) Dir() Direction {
	var ret Direction
	value := _this.Value_JS.Get("dir")
	ret = DirectionFromJS(value)
	return ret
}

// Lang returning attribute 'lang' with
// type string (idl: DOMString).
func (_this *Notification) Lang() string {
	var ret string
	value := _this.Value_JS.Get("lang")
	ret = (value).String()
	return ret
}

// Body returning attribute 'body' with
// type string (idl: DOMString).
func (_this *Notification) Body() string {
	var ret string
	value := _this.Value_JS.Get("body")
	ret = (value).String()
	return ret
}

// Tag returning attribute 'tag' with
// type string (idl: DOMString).
func (_this *Notification) Tag() string {
	var ret string
	value := _this.Value_JS.Get("tag")
	ret = (value).String()
	return ret
}

// Image returning attribute 'image' with
// type string (idl: USVString).
func (_this *Notification) Image() string {
	var ret string
	value := _this.Value_JS.Get("image")
	ret = (value).String()
	return ret
}

// Icon returning attribute 'icon' with
// type string (idl: USVString).
func (_this *Notification) Icon() string {
	var ret string
	value := _this.Value_JS.Get("icon")
	ret = (value).String()
	return ret
}

// Badge returning attribute 'badge' with
// type string (idl: USVString).
func (_this *Notification) Badge() string {
	var ret string
	value := _this.Value_JS.Get("badge")
	ret = (value).String()
	return ret
}

// Vibrate returning attribute 'vibrate' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *Notification) Vibrate() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("vibrate")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

// Timestamp returning attribute 'timestamp' with
// type int (idl: unsigned long long).
func (_this *Notification) Timestamp() int {
	var ret int
	value := _this.Value_JS.Get("timestamp")
	ret = (value).Int()
	return ret
}

// Renotify returning attribute 'renotify' with
// type bool (idl: boolean).
func (_this *Notification) Renotify() bool {
	var ret bool
	value := _this.Value_JS.Get("renotify")
	ret = (value).Bool()
	return ret
}

// Silent returning attribute 'silent' with
// type bool (idl: boolean).
func (_this *Notification) Silent() bool {
	var ret bool
	value := _this.Value_JS.Get("silent")
	ret = (value).Bool()
	return ret
}

// RequireInteraction returning attribute 'requireInteraction' with
// type bool (idl: boolean).
func (_this *Notification) RequireInteraction() bool {
	var ret bool
	value := _this.Value_JS.Get("requireInteraction")
	ret = (value).Bool()
	return ret
}

// Data returning attribute 'data' with
// type Any (idl: any).
func (_this *Notification) Data() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("data")
	ret = value
	return ret
}

// Actions returning attribute 'actions' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *Notification) Actions() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("actions")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *Notification) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

// interface: Promise
type PromisePermissionMode struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromisePermissionMode) JSValue() js.Value {
	return _this.Value_JS
}

// PromisePermissionModeFromJS is casting a js.Wrapper into PromisePermissionMode.
func PromisePermissionModeFromJS(value js.Wrapper) *PromisePermissionMode {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromisePermissionMode{}
	ret.Value_JS = input
	return ret
}

func (_this *PromisePermissionMode) Then(onFulfilled *PromisePermissionModeOnFulfilled, onRejected *PromisePermissionModeOnRejected) (_result *PromisePermissionMode) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromisePermissionMode // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionModeFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionMode) Catch(onRejected *PromisePermissionModeOnRejected) (_result *PromisePermissionMode) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromisePermissionMode // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionModeFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromisePermissionMode) Finally(onFinally *javascript.PromiseFinally) (_result *PromisePermissionMode) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromisePermissionMode // javascript: Promise _what_return_name
	)
	_converted = PromisePermissionModeFromJS(_returned)
	_result = _converted
	return
}
