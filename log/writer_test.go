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
	"context"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/acmestack/log4g/assert"
	"github.com/acmestack/log4g/log"
	"github.com/acmestack/log4g/util"
)

func TestWriters(t *testing.T) {

	var fileName string
	{
		file, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		if err = file.Close(); err != nil {
			t.Fatal(err)
		}
		fileName = file.Name()
	}

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			_, err := log.Writers.Get(fileName, func() (log.Writer, error) {
				return log.NewFileWriter(fileName)
			})
			util.Panic(err).When(err != nil)
		}()
	}
	wg.Wait()

	w, err := log.Writers.Get(fileName, nil)
	assert.Nil(t, err)
	assert.NotNil(t, w)

	ctx := context.Background()
	log.Writers.Release(ctx, w)
	assert.True(t, log.Writers.Has(fileName))
	log.Writers.Release(ctx, w)
	assert.True(t, log.Writers.Has(fileName))
	log.Writers.Release(ctx, w)
	assert.False(t, log.Writers.Has(fileName))
}
