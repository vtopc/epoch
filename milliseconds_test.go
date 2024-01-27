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

type testMillisecondsValueStruct struct {
	Timestamp Milliseconds `json:"timestamp"`
}

type testMillisecondsPointerStruct struct {
	Timestamp *Milliseconds `json:"timestamp"`
}

const tms = int64(1136239445999)

func TestNewMilliseconds(t *testing.T) {
	const ns = 123 * nsPerMs
	t.Run("seconds", func(t *testing.T) {
		got := NewMilliseconds(time.Unix(tms/msPerS, ns))
		assert.Equal(t, got.Unix(), tms/msPerS)
	})

	t.Run("nanoseconds", func(t *testing.T) {
		got := NewMilliseconds(time.Unix(0, ns))
		assert.Equal(t, got.UnixNano(), ns)
	})
}

func TestMilliseconds_Unmarshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    testMillisecondsValueStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"timestamp":%d}`, tms),
				want: testMillisecondsValueStruct{
					Timestamp: Milliseconds{Time: tmsTime},
				},
			},
			"not_int": {
				v:       `{"timestamp":"text"}`,
				wantErr: errors.New("failed to parse epoch.Milliseconds: strconv.ParseInt: parsing \"\\\"text\\\"\": invalid syntax"),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testMillisecondsValueStruct
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
			want    testMillisecondsPointerStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"timestamp":%d}`, tms),
				want: testMillisecondsPointerStruct{
					Timestamp: &Milliseconds{Time: tmsTime},
				},
			},
			"nil": {
				v: `{"timestamp":null}`,
				want: testMillisecondsPointerStruct{
					Timestamp: nil,
				},
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testMillisecondsPointerStruct
				err := json.Unmarshal([]byte(tc.v), &got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			})
		}
	})
}

func TestMilliseconds_Marshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       testMillisecondsValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testMillisecondsValueStruct{
					Timestamp: Milliseconds{Time: tmsTime},
				},
				want: fmt.Sprintf(`{"timestamp":%d}`, tms),
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
			v       testMillisecondsPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testMillisecondsPointerStruct{
					Timestamp: &Milliseconds{Time: tmsTime},
				},
				want: fmt.Sprintf(`{"timestamp":%d}`, tms),
			},
			"nil": {
				v: testMillisecondsPointerStruct{
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
