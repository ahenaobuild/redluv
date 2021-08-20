package propertyfx

import (
	"github.com/hellobuild/Luv-Go/vms/secp256k1fx"
)

// OwnedOutput ...
type OwnedOutput struct {
	secp256k1fx.OutputOwners `serialize:"true"`
}
