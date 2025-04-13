package bech32_test

import (
	"testing"

	"github.com/sonr-io/crypto/bech32"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		data     []byte
		expected string
	}{
		{
			name:     "empty prefix",
			prefix:   "",
			data:     []byte{},
			expected: "1",
		},
		{
			name:     "empty data",
			prefix:   "prefix",
			data:     []byte{},
			expected: "prefix1",
		},
		{
			name:     "single data",
			prefix:   "prefix",
			data:     []byte{0x00},
			expected: "prefix1p",
		},
		{
			name:     "multiple data",
			prefix:   "prefix",
			data:     []byte{0x00, 0x01},
			expected: "prefix1q7",
		},
		{
			name:     "multiple data with padding",
			prefix:   "prefix",
			data:     []byte{0x00, 0x01, 0x02},
			expected: "prefix1qr8",
		},
		{
			name:     "multiple data with padding and no padding",
			prefix:   "prefix",
			data:     []byte{0x00, 0x01, 0x02, 0x03},
			expected: "prefix1qr8w",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := bech32.Encode(test.prefix, test.data)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("unexpected result: expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		bech     string
		expected string
		err      error
	}{
		{
			name:     "empty bech",
			bech:     "",
			expected: "",
			err:      bech32.ErrInvalidLength,
		},
		{
			name:     "invalid bech",
			bech:     "invalid",
			expected: "",
			err:      bech32.ErrInvalidCharacter,
		},
		{
			name:     "invalid checksum",
			bech:     "prefix1p",
			expected: "",
			err:      bech32.ErrInvalidChecksum,
		},
		{
			name:     "invalid separator",
			bech:     "prefix1pq",
			expected: "",
			err:      bech32.ErrInvalidSeparator,
		},
		{
			name:     "invalid prefix",
			bech:     "invalid1p",
			expected: "",
			err:      bech32.ErrInvalidPrefix,
		},
		{
			name:     "valid bech",
			bech:     "prefix1p",
			expected: "prefix",
			err:      nil,
		},
		{
			name:     "valid bech with padding",
			bech:     "prefix1qr8",
			expected: "prefix",
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _, err := bech32.Decode(test.bech)
			if err != nil && err != test.err {
				t.Errorf("unexpected error: %v", err)
			}
			if err != nil && err == test.err {
				return
			}
			if result != test.expected {
				t.Errorf("unexpected result: expected %s, got %s", test.expected, result)
			}
		})
	}
}
