package nftfx

import (
	"github.com/hellobuild/Luv-Go/vms/secp256k1fx"
)

// MintOutput ...
type MintOutput struct {
	GroupID                  uint32 `serialize:"true" json:"groupID"`
	secp256k1fx.OutputOwners `serialize:"true"`
}
