package protocol100r5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testData struct {
	name  string
	input []byte
	res   []byte
	err   error
}

func TestCmdBuildRequest(t *testing.T) {
	tests := []testData{
		{
			name:  "SetZero Comand",
			input: []byte{cmd_set_zero},
			res:   []byte{0xF8, 0x55, 0xCE, 0x01, 0x00, 0x72, 0x25, 0xBF},
			err:   nil,
		},
		{
			name:  "SetTare Comand",
			input: []byte{cmd_set_tare, 0x00, 0x00, 0x00, 0x50},
			res:   []byte{0xF8, 0x55, 0xCE, 0x05, 0x00, 0xA3, 0x00, 0x00, 0x00, 0x50, 0x4F, 0x8F},
			err:   nil,
		},

		/*
		   {
		   			name:  "data len < 1",
		   			input: []byte{},
		   			res:   nil,
		   			err:   fmt.Errorf("empty message"),
		   		},
		*/
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			res, err := cmdBuildRequest(test.input)
			require.Equal(t, test.res, res)
			require.ErrorIs(t, test.err, err)
		})
	}

}

func TestCmdParseResponse(t *testing.T) {
	tests := []testData{
		{
			name:  "Ack Massa",
			input: []byte{0xF8, 0x55, 0xCE, 0x09, 0x00, 0x24, 0xE8, 0x03, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x84, 0xD4},
			res:   []byte{0x24, 0xE8, 0x03, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00},
			err:   nil,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			res, err := cmdParseResponse(test.input)
			require.ErrorIs(t, test.err, err)
			require.Equal(t, test.res, res)

		})
	}

}
