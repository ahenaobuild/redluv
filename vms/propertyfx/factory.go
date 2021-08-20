package propertyfx

import (
	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/snow"
)

// ID that this Fx uses when labeled
var (
	ID = ids.ID{'p', 'r', 'o', 'p', 'e', 'r', 't', 'y', 'f', 'x'}
)

// Factory ...
type Factory struct{}

// New ...
func (f *Factory) New(*snow.Context) (interface{}, error) { return &Fx{}, nil }
