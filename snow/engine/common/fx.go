// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/hellobuild/Luv-Go/ids"
)

// Fx wraps an instance of a feature extension
type Fx struct {
	ID ids.ID
	Fx interface{}
}
