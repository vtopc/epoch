package epoch

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testFloatMSValueStruct struct {
	Timestamp FloatMS `json:"timestamp"`
}

type testFloatMSPointerStruct struct {
	Timestamp *FloatMS `json:"timestamp"`
}

func TestFloatMS_Marshal(t *testing.T) {
	const js = `{"timestamp":1136239445.999}`

	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       testFloatMSValueStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testFloatMSValueStruct{
					Timestamp: FloatMS{Time: tmsTime},
				},
				want: js,
			},
			"rounding": {
				v: testFloatMSValueStruct{
					Timestamp: FloatMS{Time: time.Unix(1136239445, 500999000)},
				},
				want: `{"timestamp":1136239445.5}`,
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
			v       testFloatMSPointerStruct
			want    string
			wantErr error
		}{
			"positive": {
				v: testFloatMSPointerStruct{
					Timestamp: &FloatMS{Time: tmsTime},
				},
				want: js,
			},
			"nil": {
				v: testFloatMSPointerStruct{
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

func TestFloatMS_Unmarshal(t *testing.T) {
	const js = `{"timestamp":1136239445.999}`

	t.Run("value", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    FloatMS
			wantErr error
		}{
			"positive": {
				v:    js,
				want: FloatMS{Time: tmsTime},
			},
			"not_int": {
				v:       `{"timestamp":"text"}`,
				wantErr: errors.New("failed to parse epoch.FloatMS: strconv.ParseFloat: parsing \"\\\"text\\\"\": invalid syntax"),
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testFloatMSValueStruct
				err := json.Unmarshal([]byte(tc.v), &got)
				if tc.wantErr == nil {
					require.NoError(t, err)
					assert.Equal(t, tc.want, got.Timestamp)

					return
				}

				require.EqualError(t, err, tc.wantErr.Error())
			})
		}
	})

	t.Run("pointer", func(t *testing.T) {
		tests := map[string]struct {
			v       string
			want    *FloatMS
			wantErr error
		}{
			"positive": {
				v:    js,
				want: &FloatMS{Time: tmsTime},
			},
			"nil": {
				v:    `{"timestamp":null}`,
				want: nil,
			},
		}

		for name, tc := range tests {
			tc := tc
			t.Run(name, func(t *testing.T) {
				var got testFloatMSPointerStruct
				err := json.Unmarshal([]byte(tc.v), &got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, got.Timestamp)
			})
		}
	})
}
