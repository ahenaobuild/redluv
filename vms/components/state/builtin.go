// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"
	"time"

	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/snow/choices"
	"github.com/hellobuild/Luv-Go/utils/wrappers"
)

func marshalID(idIntf interface{}) ([]byte, error) {
	if id, ok := idIntf.(ids.ID); ok {
		return id[:], nil
	}
	return nil, errors.New("expected ids.ID but got unexpected type")
}

func unmarshalID(bytes []byte) (interface{}, error) {
	return ids.ToID(bytes)
}

func marshalStatus(statusIntf interface{}) ([]byte, error) {
	if status, ok := statusIntf.(choices.Status); ok {
		return status.Bytes(), nil
	}
	return nil, errors.New("expected choices.Status but got unexpected type")
}

func unmarshalStatus(bytes []byte) (interface{}, error) {
	p := wrappers.Packer{Bytes: bytes}
	status := choices.Status(p.UnpackInt())
	if err := status.Valid(); err != nil {
		return nil, err
	}
	return status, p.Err
}

func marshalTime(timeIntf interface{}) ([]byte, error) {
	if t, ok := timeIntf.(time.Time); ok {
		p := wrappers.Packer{MaxSize: wrappers.LongLen}
		p.PackLong(uint64(t.Unix()))
		return p.Bytes, p.Err
	}
	return nil, errors.New("expected time.Time but got unexpected type")
}

func unmarshalTime(bytes []byte) (interface{}, error) {
	p := wrappers.Packer{Bytes: bytes}
	unixTime := p.UnpackLong()
	return time.Unix(int64(unixTime), 0), nil
}
