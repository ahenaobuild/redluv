// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/hellobuild/Luv-Go/database/memdb"
	"github.com/hellobuild/Luv-Go/snow/consensus/snowball"
	"github.com/hellobuild/Luv-Go/snow/consensus/snowman"
	"github.com/hellobuild/Luv-Go/snow/engine/common"
	"github.com/hellobuild/Luv-Go/snow/engine/common/queue"
	"github.com/hellobuild/Luv-Go/snow/engine/snowman/block"
	"github.com/hellobuild/Luv-Go/snow/engine/snowman/bootstrap"
)

func DefaultConfig() Config {
	blocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())
	return Config{
		Config: bootstrap.Config{
			Config:  common.DefaultConfigTest(),
			Blocked: blocked,
			VM:      &block.TestVM{},
		},
		Params: snowball.Parameters{
			Metrics:               prometheus.NewRegistry(),
			K:                     1,
			Alpha:                 1,
			BetaVirtuous:          1,
			BetaRogue:             2,
			ConcurrentRepolls:     1,
			OptimalProcessing:     100,
			MaxOutstandingItems:   1,
			MaxItemProcessingTime: 1,
		},
		Consensus: &snowman.Topological{},
	}
}
