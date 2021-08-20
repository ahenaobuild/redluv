// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/hellobuild/Luv-Go/codec"
	"github.com/hellobuild/Luv-Go/codec/linearcodec"
	"github.com/hellobuild/Luv-Go/codec/reflectcodec"
)

const (
	maxPackerSize  = 1 << 30 // max size, in bytes, of something being marshalled by Marshal()
	maxSliceLength = 1 << 18

	codecVersion = 0
)

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLength)
	c = codec.NewManager(maxPackerSize)
	if err := c.RegisterCodec(codecVersion, lc); err != nil {
		panic(err)
	}
}
