package epoch

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStrMillisecondsValueStruct struct {
	Timestamp StrMilliseconds `json:"timestamp"`
}

type testStrMillisecondsPointerStruct struct {
	Timestamp *StrMilliseconds `json:"timestamp"`
}

func TestNewStrMilliseconds(t *testing.T) {
	const ns = 123 * nsPerMs
	t.Run("seconds", func(t *testing.T) {
		got := NewStrMilliseconds(time.Unix(tms/msPerS, ns))
		assert.Equal(t, got.Unix(), tms/msPerS)
	})

	t.Run("nanoseconds", func(t *testing.T) {
		got := NewStrMilliseconds(time.Unix(0, ns))
		assert.Equal(t, got.UnixNano(), ns)
	})
}

func TestStrMilliseconds_Unmarshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			j       string
			want    testStrMillisecondsValueStruct
			wantErr error
		}{
			"positive": {
				j: fmt.Sprintf(`{"timestamp":"%d"}`, tms),
				want: testStrMillisecondsValueStruct{
					Timestamp: StrMilliseconds{Time: tmsTime},
				},
			},
			"not_int": {
				j:       `{"timestamp":"text"}`,
				wantErr: errors.New(`failed to parse epoch.StrMilliseconds: strconv.ParseInt: parsing "text": invalid syntax`),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testStrMillisecondsValueStruct
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
			want    testStrMillisecondsPointerStruct
			wantErr error
		}{
			"positive": {
				j: fmt.Sprintf(`{"timestamp":"%d"}`, tms),
				want: testStrMillisecondsPointerStruct{
					Timestamp: &StrMilliseconds{Time: tmsTime},
				},
			},
			"nil": {
				j: `{"timestamp":null}`,
				want: testStrMillisecondsPointerStruct{
					Timestamp: nil,
				},
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testStrMillisecondsPointerStruct
				err := json.Unmarshal([]byte(tc.j), &got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}

func TestStrMilliseconds_Marshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       testStrMillisecondsValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testStrMillisecondsValueStruct{
					Timestamp: StrMilliseconds{Time: tmsTime},
				},
				want: fmt.Sprintf(`{"timestamp":"%d"}`, tms),
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
			v       testStrMillisecondsPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testStrMillisecondsPointerStruct{
					Timestamp: &StrMilliseconds{Time: tmsTime},
				},
				want: fmt.Sprintf(`{"timestamp":"%d"}`, tms),
			},
			"nil": {
				v: testStrMillisecondsPointerStruct{
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
