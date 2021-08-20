package nftfx

import (
	"github.com/hellobuild/Luv-Go/vms/secp256k1fx"
)

// Credential ...
type Credential struct {
	secp256k1fx.Credential `serialize:"true"`
}
