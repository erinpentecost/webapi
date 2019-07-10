// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package channel

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/file"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// file.Blob
// javascript.ArrayBuffer
// javascript.FrozenArray
// javascript.Object

// source idl files:
// html.idl

// transform files:
// html.go.md

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

// enum: BinaryType
type BinaryType int

const (
	Blob BinaryType = iota
	Arraybuffer
)

var binaryTypeToWasmTable = []string{
	"blob", "arraybuffer",
}

var binaryTypeFromWasmTable = map[string]BinaryType{
	"blob": Blob, "arraybuffer": Arraybuffer,
}

// JSValue is converting this enum into a java object
func (this *BinaryType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this BinaryType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(binaryTypeToWasmTable) {
		return binaryTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// BinaryTypeFromJS is converting a javascript value into
// a BinaryType enum value.
func BinaryTypeFromJS(value js.Value) BinaryType {
	key := value.String()
	conv, ok := binaryTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: CloseEventInit
type CloseEventInit struct {
	Bubbles    bool
	Cancelable bool
	Composed   bool
	WasClean   bool
	Code       int
	Reason     string
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *CloseEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.WasClean
	out.Set("wasClean", value3)
	value4 := _this.Code
	out.Set("code", value4)
	value5 := _this.Reason
	out.Set("reason", value5)
	return out
}

// CloseEventInitFromJS is allocating a new
// CloseEventInit object and copy all values from
// input javascript object
func CloseEventInitFromJS(value js.Wrapper) *CloseEventInit {
	input := value.JSValue()
	var out CloseEventInit
	var (
		value0 bool   // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool   // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool   // javascript: boolean {composed Composed composed}
		value3 bool   // javascript: boolean {wasClean WasClean wasClean}
		value4 int    // javascript: unsigned short {code Code code}
		value5 string // javascript: USVString {reason Reason reason}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("wasClean")).Bool()
	out.WasClean = value3
	value4 = (input.Get("code")).Int()
	out.Code = value4
	value5 = (input.Get("reason")).String()
	out.Reason = value5
	return &out
}

// dictionary: MessageEventInit
type MessageEventInit struct {
	Bubbles     bool
	Cancelable  bool
	Composed    bool
	Data        js.Value
	Origin      string
	LastEventId string
	Source      *Union
	Ports       []*MessagePort
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *MessageEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.Data
	out.Set("data", value3)
	value4 := _this.Origin
	out.Set("origin", value4)
	value5 := _this.LastEventId
	out.Set("lastEventId", value5)
	value6 := _this.Source.JSValue()
	out.Set("source", value6)
	value7 := js.Global().Get("Array").New(len(_this.Ports))
	for __idx7, __seq_in7 := range _this.Ports {
		__seq_out7 := __seq_in7.JSValue()
		value7.SetIndex(__idx7, __seq_out7)
	}
	out.Set("ports", value7)
	return out
}

// MessageEventInitFromJS is allocating a new
// MessageEventInit object and copy all values from
// input javascript object
func MessageEventInitFromJS(value js.Wrapper) *MessageEventInit {
	input := value.JSValue()
	var out MessageEventInit
	var (
		value0 bool           // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool           // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool           // javascript: boolean {composed Composed composed}
		value3 js.Value       // javascript: any {data Data data}
		value4 string         // javascript: USVString {origin Origin origin}
		value5 string         // javascript: DOMString {lastEventId LastEventId lastEventId}
		value6 *Union         // javascript: Union {source Source source}
		value7 []*MessagePort // javascript: sequence<MessagePort> {ports Ports ports}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = input.Get("data")
	out.Data = value3
	value4 = (input.Get("origin")).String()
	out.Origin = value4
	value5 = (input.Get("lastEventId")).String()
	out.LastEventId = value5
	if input.Get("source").Type() != js.TypeNull && input.Get("source").Type() != js.TypeUndefined {
		value6 = UnionFromJS(input.Get("source"))
	}
	out.Source = value6
	__length7 := input.Get("ports").Length()
	__array7 := make([]*MessagePort, __length7, __length7)
	for __idx7 := 0; __idx7 < __length7; __idx7++ {
		var __seq_out7 *MessagePort
		__seq_in7 := input.Get("ports").Index(__idx7)
		__seq_out7 = MessagePortFromJS(__seq_in7)
		__array7[__idx7] = __seq_out7
	}
	value7 = __array7
	out.Ports = value7
	return &out
}

// dictionary: PostMessageOptions
type PostMessageOptions struct {
	Transfer []*javascript.Object
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *PostMessageOptions) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := js.Global().Get("Array").New(len(_this.Transfer))
	for __idx0, __seq_in0 := range _this.Transfer {
		__seq_out0 := __seq_in0.JSValue()
		value0.SetIndex(__idx0, __seq_out0)
	}
	out.Set("transfer", value0)
	return out
}

// PostMessageOptionsFromJS is allocating a new
// PostMessageOptions object and copy all values from
// input javascript object
func PostMessageOptionsFromJS(value js.Wrapper) *PostMessageOptions {
	input := value.JSValue()
	var out PostMessageOptions
	var (
		value0 []*javascript.Object // javascript: sequence<object> {transfer Transfer transfer}
	)
	__length0 := input.Get("transfer").Length()
	__array0 := make([]*javascript.Object, __length0, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *javascript.Object
		__seq_in0 := input.Get("transfer").Index(__idx0)
		__seq_out0 = javascript.ObjectFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	value0 = __array0
	out.Transfer = value0
	return &out
}

// interface: BroadcastChannel
type BroadcastChannel struct {
	domcore.EventTarget
}

// BroadcastChannelFromJS is casting a js.Wrapper into BroadcastChannel.
func BroadcastChannelFromJS(value js.Wrapper) *BroadcastChannel {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &BroadcastChannel{}
	ret.Value_JS = input
	return ret
}

func NewBroadcastChannel(name string) (_result *BroadcastChannel) {
	_klass := js.Global().Get("BroadcastChannel")
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *BroadcastChannel // javascript: BroadcastChannel _what_return_name
	)
	_converted = BroadcastChannelFromJS(_returned)
	_result = _converted
	return
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *BroadcastChannel) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BroadcastChannel) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BroadcastChannel) SetOnmessage(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessage", input)
}

// Onmessageerror returning attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BroadcastChannel) Onmessageerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessageerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessageerror setting attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *BroadcastChannel) SetOnmessageerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessageerror", input)
}

func (_this *BroadcastChannel) PostMessage(message interface{}) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *BroadcastChannel) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

// interface: CloseEvent
type CloseEvent struct {
	domcore.Event
}

// CloseEventFromJS is casting a js.Wrapper into CloseEvent.
func CloseEventFromJS(value js.Wrapper) *CloseEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &CloseEvent{}
	ret.Value_JS = input
	return ret
}

func NewCloseEvent(_type string, eventInitDict *CloseEventInit) (_result *CloseEvent) {
	_klass := js.Global().Get("CloseEvent")
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
		_converted *CloseEvent // javascript: CloseEvent _what_return_name
	)
	_converted = CloseEventFromJS(_returned)
	_result = _converted
	return
}

// WasClean returning attribute 'wasClean' with
// type bool (idl: boolean).
func (_this *CloseEvent) WasClean() bool {
	var ret bool
	value := _this.Value_JS.Get("wasClean")
	ret = (value).Bool()
	return ret
}

// Code returning attribute 'code' with
// type int (idl: unsigned short).
func (_this *CloseEvent) Code() int {
	var ret int
	value := _this.Value_JS.Get("code")
	ret = (value).Int()
	return ret
}

// Reason returning attribute 'reason' with
// type string (idl: USVString).
func (_this *CloseEvent) Reason() string {
	var ret string
	value := _this.Value_JS.Get("reason")
	ret = (value).String()
	return ret
}

// interface: MessageChannel
type MessageChannel struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *MessageChannel) JSValue() js.Value {
	return _this.Value_JS
}

// MessageChannelFromJS is casting a js.Wrapper into MessageChannel.
func MessageChannelFromJS(value js.Wrapper) *MessageChannel {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MessageChannel{}
	ret.Value_JS = input
	return ret
}

func NewMessageChannel() (_result *MessageChannel) {
	_klass := js.Global().Get("MessageChannel")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *MessageChannel // javascript: MessageChannel _what_return_name
	)
	_converted = MessageChannelFromJS(_returned)
	_result = _converted
	return
}

// Port1 returning attribute 'port1' with
// type MessagePort (idl: MessagePort).
func (_this *MessageChannel) Port1() *MessagePort {
	var ret *MessagePort
	value := _this.Value_JS.Get("port1")
	ret = MessagePortFromJS(value)
	return ret
}

// Port2 returning attribute 'port2' with
// type MessagePort (idl: MessagePort).
func (_this *MessageChannel) Port2() *MessagePort {
	var ret *MessagePort
	value := _this.Value_JS.Get("port2")
	ret = MessagePortFromJS(value)
	return ret
}

// interface: MessageEvent
type MessageEvent struct {
	domcore.Event
}

// MessageEventFromJS is casting a js.Wrapper into MessageEvent.
func MessageEventFromJS(value js.Wrapper) *MessageEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MessageEvent{}
	ret.Value_JS = input
	return ret
}

func NewMessageEvent(_type string, eventInitDict *MessageEventInit) (_result *MessageEvent) {
	_klass := js.Global().Get("MessageEvent")
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
		_converted *MessageEvent // javascript: MessageEvent _what_return_name
	)
	_converted = MessageEventFromJS(_returned)
	_result = _converted
	return
}

// Data returning attribute 'data' with
// type Any (idl: any).
func (_this *MessageEvent) Data() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("data")
	ret = value
	return ret
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *MessageEvent) Origin() string {
	var ret string
	value := _this.Value_JS.Get("origin")
	ret = (value).String()
	return ret
}

// LastEventId returning attribute 'lastEventId' with
// type string (idl: DOMString).
func (_this *MessageEvent) LastEventId() string {
	var ret string
	value := _this.Value_JS.Get("lastEventId")
	ret = (value).String()
	return ret
}

// Source returning attribute 'source' with
// type Union (idl: Union).
func (_this *MessageEvent) Source() *Union {
	var ret *Union
	value := _this.Value_JS.Get("source")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = UnionFromJS(value)
	}
	return ret
}

// Ports returning attribute 'ports' with
// type javascript.FrozenArray (idl: FrozenArray).
func (_this *MessageEvent) Ports() *javascript.FrozenArray {
	var ret *javascript.FrozenArray
	value := _this.Value_JS.Get("ports")
	ret = javascript.FrozenArrayFromJS(value)
	return ret
}

func (_this *MessageEvent) InitMessageEvent(_type string, bubbles *bool, cancelable *bool, data interface{}, origin *string, lastEventId *string, source *Union, ports []*MessagePort) {
	var (
		_args [8]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if bubbles != nil {
		_p1 := bubbles
		_args[1] = _p1
		_end++
	}
	if cancelable != nil {
		_p2 := cancelable
		_args[2] = _p2
		_end++
	}
	if data != nil {
		_p3 := data
		_args[3] = _p3
		_end++
	}
	if origin != nil {
		_p4 := origin
		_args[4] = _p4
		_end++
	}
	if lastEventId != nil {
		_p5 := lastEventId
		_args[5] = _p5
		_end++
	}
	if source != nil {
		_p6 := source.JSValue()
		_args[6] = _p6
		_end++
	}
	if ports != nil {
		_p7 := js.Global().Get("Array").New(len(ports))
		for __idx7, __seq_in7 := range ports {
			__seq_out7 := __seq_in7.JSValue()
			_p7.SetIndex(__idx7, __seq_out7)
		}
		_args[7] = _p7
		_end++
	}
	_this.Value_JS.Call("initMessageEvent", _args[0:_end]...)
	return
}

// interface: MessagePort
type MessagePort struct {
	domcore.EventTarget
}

// MessagePortFromJS is casting a js.Wrapper into MessagePort.
func MessagePortFromJS(value js.Wrapper) *MessagePort {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &MessagePort{}
	ret.Value_JS = input
	return ret
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MessagePort) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MessagePort) SetOnmessage(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessage", input)
}

// Onmessageerror returning attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MessagePort) Onmessageerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessageerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessageerror setting attribute 'onmessageerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *MessagePort) SetOnmessageerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessageerror", input)
}

func (_this *MessagePort) PostMessage(message interface{}, transfer []*javascript.Object) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_p1 := js.Global().Get("Array").New(len(transfer))
	for __idx1, __seq_in1 := range transfer {
		__seq_out1 := __seq_in1.JSValue()
		_p1.SetIndex(__idx1, __seq_out1)
	}
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *MessagePort) PostMessage2(message interface{}, options *PostMessageOptions) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("postMessage", _args[0:_end]...)
	return
}

func (_this *MessagePort) Start() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("start", _args[0:_end]...)
	return
}

func (_this *MessagePort) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

// interface: WebSocket
type WebSocket struct {
	domcore.EventTarget
}

// WebSocketFromJS is casting a js.Wrapper into WebSocket.
func WebSocketFromJS(value js.Wrapper) *WebSocket {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &WebSocket{}
	ret.Value_JS = input
	return ret
}

const (
	CONNECTING int = 0
	OPEN       int = 1
	CLOSING    int = 2
	CLOSED     int = 3
)

func NewWebSocket(url string, protocols *Union) (_result *WebSocket) {
	_klass := js.Global().Get("WebSocket")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := url
	_args[0] = _p0
	_end++
	if protocols != nil {
		_p1 := protocols.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *WebSocket // javascript: WebSocket _what_return_name
	)
	_converted = WebSocketFromJS(_returned)
	_result = _converted
	return
}

// Url returning attribute 'url' with
// type string (idl: USVString).
func (_this *WebSocket) Url() string {
	var ret string
	value := _this.Value_JS.Get("url")
	ret = (value).String()
	return ret
}

// ReadyState returning attribute 'readyState' with
// type int (idl: unsigned short).
func (_this *WebSocket) ReadyState() int {
	var ret int
	value := _this.Value_JS.Get("readyState")
	ret = (value).Int()
	return ret
}

// BufferedAmount returning attribute 'bufferedAmount' with
// type int (idl: unsigned long long).
func (_this *WebSocket) BufferedAmount() int {
	var ret int
	value := _this.Value_JS.Get("bufferedAmount")
	ret = (value).Int()
	return ret
}

// Onopen returning attribute 'onopen' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) Onopen() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onopen")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnopen setting attribute 'onopen' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) SetOnopen(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onopen", input)
}

// Onerror returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) Onerror() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnerror setting attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) SetOnerror(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onerror", input)
}

// Onclose returning attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) Onclose() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onclose")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnclose setting attribute 'onclose' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) SetOnclose(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onclose", input)
}

// Extensions returning attribute 'extensions' with
// type string (idl: DOMString).
func (_this *WebSocket) Extensions() string {
	var ret string
	value := _this.Value_JS.Get("extensions")
	ret = (value).String()
	return ret
}

// Protocol returning attribute 'protocol' with
// type string (idl: DOMString).
func (_this *WebSocket) Protocol() string {
	var ret string
	value := _this.Value_JS.Get("protocol")
	ret = (value).String()
	return ret
}

// Onmessage returning attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) Onmessage() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onmessage")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// SetOnmessage setting attribute 'onmessage' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *WebSocket) SetOnmessage(value *domcore.EventHandler) {
	var __callback0 js.Value
	if value != nil {
		__callback0 = (*value).Value
	} else {
		__callback0 = js.Null()
	}
	input := __callback0
	_this.Value_JS.Set("onmessage", input)
}

// BinaryType returning attribute 'binaryType' with
// type BinaryType (idl: BinaryType).
func (_this *WebSocket) BinaryType() BinaryType {
	var ret BinaryType
	value := _this.Value_JS.Get("binaryType")
	ret = BinaryTypeFromJS(value)
	return ret
}

// SetBinaryType setting attribute 'binaryType' with
// type BinaryType (idl: BinaryType).
func (_this *WebSocket) SetBinaryType(value BinaryType) {
	input := value.JSValue()
	_this.Value_JS.Set("binaryType", input)
}

func (_this *WebSocket) Close(code *int, reason *string) {
	var (
		_args [2]interface{}
		_end  int
	)
	if code != nil {
		_p0 := code
		_args[0] = _p0
		_end++
	}
	if reason != nil {
		_p1 := reason
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("close", _args[0:_end]...)
	return
}

func (_this *WebSocket) Send(data string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *WebSocket) Send2(data *file.Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *WebSocket) Send3(data *javascript.ArrayBuffer) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *WebSocket) Send4(data *Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := data.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}
