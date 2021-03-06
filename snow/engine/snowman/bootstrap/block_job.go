// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrap

import (
	"errors"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/snow/choices"
	"github.com/hellobuild/Luv-Go/snow/consensus/snowman"
	"github.com/hellobuild/Luv-Go/snow/engine/common/queue"
	"github.com/hellobuild/Luv-Go/snow/engine/snowman/block"
	"github.com/hellobuild/Luv-Go/utils/logging"
)

type parser struct {
	log                     logging.Logger
	numAccepted, numDropped prometheus.Counter
	vm                      block.ChainVM
}

func (p *parser) Parse(blkBytes []byte) (queue.Job, error) {
	blk, err := p.vm.ParseBlock(blkBytes)
	if err != nil {
		return nil, err
	}
	return &blockJob{
		parser:      p,
		log:         p.log,
		numAccepted: p.numAccepted,
		numDropped:  p.numDropped,
		blk:         blk,
	}, nil
}

type blockJob struct {
	parser                  *parser
	log                     logging.Logger
	numAccepted, numDropped prometheus.Counter
	blk                     snowman.Block
}

func (b *blockJob) ID() ids.ID { return b.blk.ID() }
func (b *blockJob) MissingDependencies() (ids.Set, error) {
	missing := ids.Set{}
	if parent := b.blk.Parent(); parent.Status() != choices.Accepted {
		missing.Add(parent.ID())
	}
	return missing, nil
}

func (b *blockJob) HasMissingDependencies() (bool, error) {
	if parent := b.blk.Parent(); parent.Status() != choices.Accepted {
		return true, nil
	}
	return false, nil
}

func (b *blockJob) Execute() error {
	hasMissingDeps, err := b.HasMissingDependencies()
	if err != nil {
		return err
	}
	if hasMissingDeps {
		b.numDropped.Inc()
		return errors.New("attempting to accept a block with missing dependencies")
	}
	status := b.blk.Status()
	switch status {
	case choices.Unknown, choices.Rejected:
		b.numDropped.Inc()
		return fmt.Errorf("attempting to execute block with status %s", status)
	case choices.Processing:
		blkID := b.blk.ID()
		if err := b.blk.Verify(); err != nil {
			b.log.Error("block %s failed verification during bootstrapping due to %s", blkID, err)
			return fmt.Errorf("failed to verify block in bootstrapping: %w", err)
		}

		b.numAccepted.Inc()
		b.log.Trace("accepting block %s in bootstrapping", blkID)
		if err := b.blk.Accept(); err != nil {
			b.log.Debug("block %s failed to accept during bootstrapping due to %s", blkID, err)
			return fmt.Errorf("failed to accept block in bootstrapping: %w", err)
		}
	}
	return nil
}
func (b *blockJob) Bytes() []byte { return b.blk.Bytes() }
