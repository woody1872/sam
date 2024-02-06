package checksum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyHashAlgorithm(t *testing.T) {
	tests := map[string]struct {
		got  string
		want error
	}{
		"valid algorithm":           {got: "sha256", want: nil},
		"valid algorithm uppercase": {got: "SHA256", want: nil},
		"invalid algorithm":         {got: "xxxxxxx", want: ErrInvalidHashAlgorithm},
		"empty string":              {got: "", want: ErrInvalidHashAlgorithm},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			var err error
			_, err = NewHashAlgorithm(tc.got)
			assert.Equal(t, tc.want, err)
		})
	}
}
