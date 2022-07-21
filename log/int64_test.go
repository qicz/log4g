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

package log

import (
	"reflect"
	"sync"
	"testing"
	"unsafe"

	"github.com/acmestack/log4g/assert"
)

func TestInt64(t *testing.T) {

	// atomic.Int64 和 int64 占用的空间大小一样
	assert.Equal(t, unsafe.Sizeof(Int64{}), uintptr(8))

	var i Int64
	assert.Equal(t, i.Load(), int64(0))

	i.Store(1)
	assert.Equal(t, i.Load(), int64(1))
}

func TestReflectInt64(t *testing.T) {

	// s 必须分配在堆上
	s := new(struct {
		I Int64
	})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		addr := reflect.ValueOf(s).Elem().Field(0).Addr()
		v, ok := addr.Interface().(*Int64)
		assert.True(t, ok)
		for i := 0; i < 10; i++ {
			v.Add(1)
		}
	}()
	go func() {
		defer wg.Done()
		addr := reflect.ValueOf(s).Elem().Field(0).Addr()
		v, ok := addr.Interface().(*Int64)
		assert.True(t, ok)
		for i := 0; i < 10; i++ {
			v.Add(2)
		}
	}()
	wg.Wait()
	assert.Equal(t, int64(30), s.I.Load())
}
