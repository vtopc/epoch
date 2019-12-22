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

type testSecondsValueStruct struct {
	Timestamp Seconds `json:"timestamp"`
}

type testSecondsPointerStruct struct {
	Timestamp *Seconds `json:"timestamp"`
}

const ts = int64(1136239445)

func TestNewSeconds(t *testing.T) {
	got := NewSeconds(time.Unix(ts, 0))
	assert.NotEqual(t, Seconds{}, got)
}

func TestSeconds_Unmarshal(t *testing.T) {
	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    testSecondsValueStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"timestamp":%d}`, ts),
				want: testSecondsValueStruct{
					Timestamp: Seconds{Time: time.Unix(ts, 0)},
				},
			},
			"not_int": {
				v:       `{"timestamp":"text"}`,
				wantErr: errors.New("failed to parse Seconds: strconv.ParseInt: parsing \"\\\"text\\\"\": invalid syntax"),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testSecondsValueStruct
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
			want    testSecondsPointerStruct
			wantErr error
		}{
			"positive": {
				v: fmt.Sprintf(`{"timestamp":%d}`, ts),
				want: testSecondsPointerStruct{
					Timestamp: &Seconds{Time: time.Unix(ts, 0)},
				},
			},
			"nil": {
				v: `{"timestamp":null}`,
				want: testSecondsPointerStruct{
					Timestamp: nil,
				},
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testSecondsPointerStruct
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
			v       testSecondsValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testSecondsValueStruct{
					Timestamp: Seconds{Time: time.Unix(ts, 0)},
				},
				want: fmt.Sprintf(`{"timestamp":%d}`, ts),
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
			v       testSecondsPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testSecondsPointerStruct{
					Timestamp: &Seconds{Time: time.Unix(ts, 0)},
				},
				want: fmt.Sprintf(`{"timestamp":%d}`, ts),
			},
			"nil": {
				v: testSecondsPointerStruct{
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
