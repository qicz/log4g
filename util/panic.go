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

// PanicCond 封装触发 panic 的条件。
type PanicCond struct {
	fn func() interface{}
}

// When 满足给定条件时抛出一个 panic 。
func (p *PanicCond) When(isPanic bool) {
	if isPanic {
		panic(p.fn())
	}
}

// NewPanicCond PanicCond 的构造函数。
func NewPanicCond(fn func() interface{}) *PanicCond {
	return &PanicCond{fn}
}

// Panic 抛出一个异常值。
func Panic(err error) *PanicCond {
	return NewPanicCond(func() interface{} { return err })
}
