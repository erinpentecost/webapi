// Code generated by webidl-bind. DO NOT EDIT.

package webidl

import "syscall/js"

// using following types:

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

// callback: Function
type FunctionFunc func(arguments []js.Value) interface{}

// Function is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type Function js.Func

func FunctionToJS(callback FunctionFunc) *Function {
	if callback == nil {
		return nil
	}
	ret := Function(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 []js.Value // javascript: any arguments
		)
		_p0 = make([]js.Value, 0, len(args[0:]))
		for _, __in := range args[0:] {
			var __out js.Value
			__out = __in
			_p0 = append(_p0, __out)
		}
		_returned := callback(_p0)
		_converted := _returned
		return _converted
	}))
	return &ret
}

func FunctionFromJS(_value js.Value) FunctionFunc {
	return func(arguments []js.Value) (_result interface{}) {
		var (
			_args []interface{} = make([]interface{}, 0+len(arguments))
			_end  int
		)
		for _, __in := range arguments {
			__out := __in
			_args[_end] = __out
			_end++
		}
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted js.Value // javascript: any
		)
		_converted = _returned
		_result = _converted
		return
	}
}

// callback: VoidFunction
type VoidFunctionFunc func()

// VoidFunction is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type VoidFunction js.Func

func VoidFunctionToJS(callback VoidFunctionFunc) *VoidFunction {
	if callback == nil {
		return nil
	}
	ret := VoidFunction(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var ()
		callback()
		// returning no return value
		return nil
	}))
	return &ret
}

func VoidFunctionFromJS(_value js.Value) VoidFunctionFunc {
	return func() {
		var (
			_args [0]interface{}
			_end  int
		)
		_value.Invoke(_args[0:_end]...)
		return
	}
}