// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package modeldecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/apm-server/model"
)

func TestProcessDecode(t *testing.T) {
	pid, ppid, title, argv := 123, 456, "foo", []string{"a", "b"}
	for _, test := range []struct {
		input map[string]interface{}
		p     model.Process
	}{
		{input: nil},
		{
			input: map[string]interface{}{
				"pid": 123.0, "ppid": 456.0, "title": title, "argv": []interface{}{"a", "b"},
			},
			p: model.Process{Pid: pid, Ppid: &ppid, Title: title, Argv: argv},
		},
	} {
		var proc model.Process
		decodeProcess(test.input, &proc)
		assert.Equal(t, test.p, proc)
	}
}
