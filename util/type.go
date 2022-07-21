/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"reflect"
)

// errorType error 的反射类型。
var errorType = reflect.TypeOf((*error)(nil)).Elem()

// IsNil 返回 v 的值是否为 nil，但是不会 panic 。
func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Ptr,
		reflect.Slice,
		reflect.UnsafePointer:
		return v.IsNil()
	}
	return !v.IsValid()
}

// IsErrorType t 是否是 error 类型。
func IsErrorType(t reflect.Type) bool {
	return t == errorType
}

// IsPrimitiveValueType returns whether the input type is the primitive value
// type which only is int, unit, float, bool, string and complex.
func IsPrimitiveValueType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Complex64, reflect.Complex128:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.String:
		return true
	case reflect.Bool:
		return true
	}
	return false
}

// IsValueType returns whether the input type is the value type which is the
// primitive value type and their one dimensional composite type including array,
// slice, map and struct, such as [3]string, []string, []int, map[int]int, etc.
func IsValueType(t reflect.Type) bool {
	fn := func(t reflect.Type) bool {
		return IsPrimitiveValueType(t) || t.Kind() == reflect.Struct
	}
	switch t.Kind() {
	case reflect.Map, reflect.Slice, reflect.Array:
		return fn(t.Elem())
	default:
		return fn(t)
	}
}

// Converter converts string value into user-defined value. It should
// be function type, and its prototype is func(string)(type,error).
type Converter interface{}

// IsValidConverter 返回是否是合法的转换器类型。
func IsValidConverter(t reflect.Type) bool {
	return t.Kind() == reflect.Func &&
		t.NumIn() == 1 &&
		t.In(0).Kind() == reflect.String &&
		t.NumOut() == 2 &&
		IsValueType(t.Out(0)) &&
		IsErrorType(t.Out(1))
}
