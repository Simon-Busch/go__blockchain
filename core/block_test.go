package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version: 		1,
		PrevBlock: 	types.RandomHash(),
		Timestamp: 	time.Now().UnixNano(),
		Height: 		10,
		Nonce: 			9999,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h.Version, hDecode.Version)
	assert.Equal(t, h.PrevBlock, hDecode.PrevBlock)
	assert.Equal(t, h.Timestamp, hDecode.Timestamp)
	assert.Equal(t, h.Height, hDecode.Height)
	assert.Equal(t, h.Nonce, hDecode.Nonce)
}
