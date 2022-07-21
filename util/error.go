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
	"errors"
	"fmt"
)

var WrapFormat = func(err error, fileline string, format string, a ...interface{}) error {
	if err == nil {
		if format != "" {
			return fmt.Errorf(fileline+" "+format, a...)
		}
		return errors.New(fileline + " " + fmt.Sprint(a...))
	}
	if format == "" {
		return fmt.Errorf("%s %s; %w", fileline, fmt.Sprint(a...), err)
	}
	return fmt.Errorf("%s %s; %w", fileline, fmt.Sprintf(format, a...), err)
}

// Error 创建携带文件信息的 error 对象。文件信息未来也许可以在编译期计算。
func Error(fileline string, text string) error {
	return WrapFormat(nil, fileline, "", text)
}

// Errorf 创建携带文件信息的 error 对象。文件信息未来也许可以在编译期计算。
func Errorf(fileline string, format string, a ...interface{}) error {
	return WrapFormat(nil, fileline, format, a...)
}

// Wrap 创建携带文件信息的 error 对象。文件信息未来也许可以在编译期计算。
func Wrap(err error, fileline string, text string) error {
	return WrapFormat(err, fileline, "", text)
}

// Wrapf 创建携带文件信息的 error 对象。文件信息未来也许可以在编译期计算。
func Wrapf(err error, fileline string, format string, a ...interface{}) error {
	return WrapFormat(err, fileline, format, a...)
}
