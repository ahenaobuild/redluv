package avm

import (
	"bytes"
	"testing"

	"github.com/hellobuild/Luv-Go/pubsub"
	"github.com/hellobuild/Luv-Go/vms/components/avax"
	"github.com/hellobuild/Luv-Go/vms/secp256k1fx"
	"github.com/stretchr/testify/assert"

	"github.com/hellobuild/Luv-Go/ids"
)

type mockFilter struct {
	addr []byte
}

func (f *mockFilter) Check(addr []byte) bool {
	return bytes.Equal(addr, f.addr)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	addrID := ids.ShortID{1}
	tx := Tx{UnsignedTx: &BaseTx{BaseTx: avax.BaseTx{
		Outs: []*avax.TransferableOutput{
			{
				Out: &secp256k1fx.TransferOutput{
					OutputOwners: secp256k1fx.OutputOwners{
						Addrs: []ids.ShortID{addrID},
					},
				},
			},
		},
	}}}
	addrBytes := addrID[:]

	fp := pubsub.NewFilterParam()
	err := fp.Add(addrBytes)
	assert.NoError(err)

	parser := NewPubSubFilterer(&tx)
	fr, _ := parser.Filter([]pubsub.Filter{&mockFilter{addr: addrBytes}})
	assert.Equal([]bool{true}, fr)
}
