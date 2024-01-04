/*
 * Copyright 2023 cluetec GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package destination

import (
	"github.com/cluetec/lifeboat/internal/config"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		c config.DestinationConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Destination
		wantErr bool
	}{
		{
			name:    "Throw error if type is not known",
			args:    args{c: config.DestinationConfig{Type: "not-known-type"}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
