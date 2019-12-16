package epoch

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testValueStruct struct {
	Time Seconds `json:"time"`
}

type testPointerStruct struct {
	Time *Seconds `json:"time"`
}

const timestamp = int64(1136239445)

func TestNewSeconds(t *testing.T) {
	got := NewSeconds(time.Unix(timestamp, 0))
	assert.NotEqual(t, Seconds{}, got)
}

func TestSeconds_Unmarshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    testValueStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"time":%d}`, timestamp),
				want: testValueStruct{
					Time: Seconds{time.Unix(timestamp, 0)},
				},
			},
			"not_int": {
				v:       `{"time":"text"}`,
				wantErr: errors.New("failed to parse int: strconv.ParseInt: parsing \"\\\"text\\\"\": invalid syntax"),
			},
		}

		for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				var got testValueStruct
				err := json.Unmarshal([]byte(tc.v), &got)
				if tc.wantErr == nil {
					require.NoError(t, err)
					assert.Equal(t, tc.want, got)

					return
				}

				require.EqualError(t, err, tc.wantErr.Error())
			})
		}
	})

	t.Run("pointer", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    testPointerStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"time":%d}`, timestamp),
				want: testPointerStruct{
					Time: &Seconds{Time: time.Unix(timestamp, 0)},
				},
			},
			"nil": {
				v: `{"time":null}`,
				want: testPointerStruct{
					Time: nil,
				},
			},
		}

		for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				var got testPointerStruct
				err := json.Unmarshal([]byte(tc.v), &got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}

func TestSeconds_Marshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       testValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testValueStruct{
					Time: Seconds{Time: time.Unix(timestamp, 0)},
				},
				want: fmt.Sprintf(`{"time":%d}`, timestamp),
			},
		}

		for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				got, err := json.Marshal(tc.v)
				require.NoError(t, err)
				assert.Equal(t, tc.want, string(got))
			})
		}
	})

	t.Run("pointer", func(t *testing.T) {
		tests := map[string]struct {
			v       testPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testPointerStruct{
					Time: &Seconds{Time: time.Unix(timestamp, 0)},
				},
				want: fmt.Sprintf(`{"time":%d}`, timestamp),
			},
			"nil": {
				v: testPointerStruct{
					Time: nil,
				},
				want: `{"time":null}`,
			},
		}

		for name, tc := range tests {
			t.Run(name, func(t *testing.T) {
				got, err := json.Marshal(tc.v)
				require.NoError(t, err)
				assert.Equal(t, tc.want, string(got))
			})
		}
	})
}
