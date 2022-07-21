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

package log_test

import (
	"fmt"
	"testing"

	"github.com/acmestack/log4g/assert"
	"github.com/acmestack/log4g/log"
)

func TestNewErrNo(t *testing.T) {
	assert.Panic(t, func() {
		log.NewErrno(200, 0, "")
	}, "project invalid, should be >= 1000")
	assert.Panic(t, func() {
		log.NewErrno(1000, 0, "")
	}, "code invalid, should be 1~999")
	fmt.Println(log.NewErrno(1000, 1, "").Code())
	fmt.Println(log.OK.Code(), log.OK.Msg())
	fmt.Println(log.ERROR.Code(), log.ERROR.Msg())
}
