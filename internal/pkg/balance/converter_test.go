package balance

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ConvertToCents(t *testing.T) {
	for _, tc := range []struct {
		name string
		val  float64

		error    error
		expected int64
	}{
		{
			name:     "11 to 1100",
			val:      11,
			expected: 1100,
		},
		{
			name:     "11.2 to 1120",
			val:      11.2,
			expected: 1120,
		},
		{
			name:     "11.11 to 1111",
			val:      11.11,
			expected: 1111,
		},
		{
			name:     "11.99 to 1199",
			val:      11.99,
			expected: 1199,
		},
		{
			name:     "11.00 to 1100",
			val:      11.00,
			expected: 1100,
		},
		{
			name:     "11.100; to 1110",
			val:      11.100,
			expected: 1110,
		},
		{
			name:     "11.000 to 1100",
			val:      11.000,
			expected: 1100,
		},
		{
			name:  "11.111; expect error",
			val:   11.111,
			error: ErrorIncorrectValue,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ConvertToCents(tc.val)
			if tc.error != nil {
				require.Error(t, err)
				require.Zero(t, actual)
				return
			}

			require.NoError(t, err)
			require.NotZero(t, actual)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func Test_ConvertFromCents(t *testing.T) {
	for _, tc := range []struct {
		name     string
		val      int64
		expected float64
	}{
		{
			name:     "1100 to 11.00",
			val:      11,
			expected: .11,
		},
		{
			name:     "1120 to 11.2",
			val:      1120,
			expected: 11.2,
		},
		{
			name:     "1111 to 11.11",
			val:      1111,
			expected: 11.11,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ConvertFromCents(tc.val)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
