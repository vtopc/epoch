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

type testStrSecondsValueStruct struct {
	Timestamp StrSeconds `json:"timestamp"`
}

type testStrSecondsPointerStruct struct {
	Timestamp *StrSeconds `json:"timestamp"`
}

func TestNewStrSeconds(t *testing.T) {
	got := NewStrSeconds(time.Unix(ts, 0))
	assert.NotEqual(t, StrSeconds{}, got)
}

func TestStrSeconds_Unmarshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			j       string
			want    testStrSecondsValueStruct
			wantErr error
		}{
			"positive": {
				j: fmt.Sprintf(`{"timestamp":"%d"}`, ts),
				want: testStrSecondsValueStruct{
					Timestamp: StrSeconds{Time: time.Unix(ts, 0)},
				},
			},
			"not_int": {
				j:       `{"timestamp":"text"}`,
				wantErr: errors.New(`failed to parse StrSeconds: strconv.ParseInt: parsing "text": invalid syntax`),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Log("json:", tc.j)

				var got testStrSecondsValueStruct
				err := json.Unmarshal([]byte(tc.j), &got)
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
			j       string
			want    testStrSecondsPointerStruct
			wantErr error
		}{
			"positive": {
				j: fmt.Sprintf(`{"timestamp":"%d"}`, ts),
				want: testStrSecondsPointerStruct{
					Timestamp: &StrSeconds{Time: time.Unix(ts, 0)},
				},
			},
			"nil": {
				j: `{"timestamp":null}`,
				want: testStrSecondsPointerStruct{
					Timestamp: nil,
				},
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testStrSecondsPointerStruct
				err := json.Unmarshal([]byte(tc.j), &got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}

func TestStrSeconds_Marshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       testStrSecondsValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testStrSecondsValueStruct{
					Timestamp: StrSeconds{Time: time.Unix(ts, 0)},
				},
				want: fmt.Sprintf(`{"timestamp":"%d"}`, ts),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				got, err := json.Marshal(tc.v)
				require.NoError(t, err)
				assert.Equal(t, tc.want, string(got))
			})
		}
	})

	t.Run("pointer", func(t *testing.T) {
		tests := map[string]struct {
			v       testStrSecondsPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testStrSecondsPointerStruct{
					Timestamp: &StrSeconds{Time: time.Unix(ts, 0)},
				},
				want: fmt.Sprintf(`{"timestamp":"%d"}`, ts),
			},
			"nil": {
				v: testStrSecondsPointerStruct{
					Timestamp: nil,
				},
				want: `{"timestamp":null}`,
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				got, err := json.Marshal(tc.v)
				require.NoError(t, err)
				assert.Equal(t, tc.want, string(got))
			})
		}
	})
}
