// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timestampvm

import (
	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/snow"
)

// ID is a unique identifier for this VM
var (
	ID = ids.ID{'t', 'i', 'm', 'e', 's', 't', 'a', 'm', 'p'}
)

// Factory ...
type Factory struct{}

// New ...
func (f *Factory) New(*snow.Context) (interface{}, error) { return &VM{}, nil }
