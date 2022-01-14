//
// Copyright 2021-2022 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// From https://github.com/redhat-developer/kam/tree/master/pkg/pipelines/resources
package resources

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMerge(t *testing.T) {
	mergeTests := []struct {
		src  Resources
		dest Resources
		want Resources
	}{
		{
			src:  Resources{"test1": "val1"},
			dest: Resources{},
			want: Resources{"test1": "val1"},
		},
		{
			src:  Resources{"test1": "val1"},
			dest: Resources{"test2": "val2"},
			want: Resources{"test1": "val1", "test2": "val2"},
		},
		{
			src:  Resources{"test1": "val1"},
			dest: Resources{"test1": "val2"},
			want: Resources{"test1": "val1"},
		},
	}

	for _, tt := range mergeTests {
		result := Merge(tt.src, tt.dest)

		if diff := cmp.Diff(tt.want, result); diff != "" {
			t.Fatalf("failed merge: %s\n", diff)
		}
	}
}
