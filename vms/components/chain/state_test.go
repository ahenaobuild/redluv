package chain

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/hellobuild/Luv-Go/database"
	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/snow/choices"
	"github.com/hellobuild/Luv-Go/snow/consensus/snowman"
	"github.com/hellobuild/Luv-Go/utils/hashing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

type TestBlock struct {
	*snowman.TestBlock
}

// SetStatus sets the status of the Block. Implements the BlockWrapper interface.
func (b *TestBlock) SetStatus(status choices.Status) { b.TestBlock.TestDecidable.StatusV = status }

// NewTestBlock returns a new test block with height, bytes, and ID derived from [i]
// and using [parent] as the parent block
func NewTestBlock(i uint64, parent *TestBlock) *TestBlock {
	b := []byte{byte(i)}
	id := hashing.ComputeHash256Array(b)
	return &TestBlock{
		TestBlock: &snowman.TestBlock{
			TestDecidable: choices.TestDecidable{
				IDV:     id,
				StatusV: choices.Unknown,
			},
			HeightV: i,
			ParentV: parent,
			BytesV:  b,
		},
	}
}

// NewTestBlocks generates [numBlocks] consecutive blocks starting
// at height [parent.Height() + 1].
func NewTestBlocks(numBlocks uint64, parent *TestBlock) []*TestBlock {
	blks := make([]*TestBlock, 0, numBlocks)
	var startHeight uint64
	if parent != nil {
		startHeight = parent.HeightV + 1
	}
	for i := startHeight; i < numBlocks; i++ {
		blks = append(blks, NewTestBlock(i, parent))
		parent = blks[len(blks)-1]
	}

	return blks
}

func createInternalBlockFuncs(t *testing.T, blks []*TestBlock) (func(id ids.ID) (snowman.Block, error), func(b []byte) (snowman.Block, error), func(height uint64) (ids.ID, error)) {
	blkMap := make(map[ids.ID]*TestBlock)
	blkByteMap := make(map[byte]*TestBlock)
	for _, blk := range blks {
		blkMap[blk.ID()] = blk
		blkBytes := blk.Bytes()
		if len(blkBytes) != 1 {
			t.Fatalf("Expected block bytes to be length 1, but found %d", len(blkBytes))
		}
		blkByteMap[blkBytes[0]] = blk
	}

	getBlock := func(id ids.ID) (snowman.Block, error) {
		blk, ok := blkMap[id]
		if !ok || !blk.Status().Fetched() {
			return nil, database.ErrNotFound
		}

		return blk, nil
	}

	parseBlk := func(b []byte) (snowman.Block, error) {
		if len(b) != 1 {
			return nil, fmt.Errorf("expected block bytes to be length 1, but found %d", len(b))
		}

		blk, ok := blkByteMap[b[0]]
		if !ok {
			return nil, fmt.Errorf("parsed unexpected block with bytes %x", b)
		}
		if blk.Status() == choices.Unknown {
			blk.SetStatus(choices.Processing)
		}
		blkMap[blk.ID()] = blk

		return blk, nil
	}
	getAcceptedBlockIDAtHeight := func(height uint64) (ids.ID, error) {
		for _, blk := range blks {
			if blk.Height() != height {
				continue
			}

			if blk.Status() == choices.Accepted {
				return blk.ID(), nil
			}
		}

		return ids.ID{}, fmt.Errorf("could not find accepted block at height %d", height)
	}

	return getBlock, parseBlk, getAcceptedBlockIDAtHeight
}

func cantBuildBlock() (snowman.Block, error) {
	return nil, errors.New("can't build new block")
}

// checkProcessingBlock checks that [blk] is of the correct type and is
// correctly uniquified when calling GetBlock and ParseBlock.
func checkProcessingBlock(t *testing.T, s *State, blk snowman.Block) {
	if _, ok := blk.(*BlockWrapper); !ok {
		t.Fatalf("Expected block to be of type (*BlockWrapper)")
	}

	parsedBlk, err := s.ParseBlock(blk.Bytes())
	if err != nil {
		t.Fatalf("Failed to parse verified block due to %s", err)
	}
	if parsedBlk.ID() != blk.ID() {
		t.Fatalf("Expected parsed block to have the same ID as the requested block")
	}
	if !bytes.Equal(parsedBlk.Bytes(), blk.Bytes()) {
		t.Fatalf("Expected parsed block to have the same bytes as the requested block")
	}
	if status := parsedBlk.Status(); status != choices.Processing {
		t.Fatalf("Expected parsed block to have status Processing, but found %s", status)
	}
	if parsedBlk != blk {
		t.Fatalf("Expected parsed block to return a uniquified block")
	}

	getBlk, err := s.GetBlock(blk.ID())
	if err != nil {
		t.Fatalf("Unexpected error during GetBlock for processing block %s", err)
	}
	if getBlk != parsedBlk {
		t.Fatalf("Expected GetBlock to return the same unique block as ParseBlock")
	}
}

// checkDecidedBlock asserts that [blk] is returned with the correct status by ParseBlock
// and GetBlock.
func checkDecidedBlock(t *testing.T, s *State, blk snowman.Block, expectedStatus choices.Status, cached bool) {
	if _, ok := blk.(*BlockWrapper); !ok {
		t.Fatalf("Expected block to be of type (*BlockWrapper)")
	}

	parsedBlk, err := s.ParseBlock(blk.Bytes())
	if err != nil {
		t.Fatalf("Unexpected error parsing decided block %s", err)
	}
	if parsedBlk.ID() != blk.ID() {
		t.Fatalf("ParseBlock returned block with unexpected ID %s, expected %s", parsedBlk.ID(), blk.ID())
	}
	if !bytes.Equal(parsedBlk.Bytes(), blk.Bytes()) {
		t.Fatalf("Expected parsed block to have the same bytes as the requested block")
	}
	if status := parsedBlk.Status(); status != expectedStatus {
		t.Fatalf("Expected parsed block to have status %s, but found %s", expectedStatus, status)
	}
	// If the block should be in the cache, assert that the returned block is identical to [blk]
	if cached && parsedBlk != blk {
		t.Fatalf("Expected parsed block to have been cached, but retrieved non-unique decided block")
	}

	getBlk, err := s.GetBlock(blk.ID())
	if err != nil {
		t.Fatalf("Unexpected error during GetBlock for decided block %s", err)
	}
	if getBlk.ID() != blk.ID() {
		t.Fatalf("GetBlock returned block with unexpected ID %s, expected %s", getBlk.ID(), blk.ID())
	}
	if !bytes.Equal(getBlk.Bytes(), blk.Bytes()) {
		t.Fatalf("Expected block from GetBlock to have the same bytes as the requested block")
	}
	if status := getBlk.Status(); status != expectedStatus {
		t.Fatalf("Expected block from GetBlock to have status %s, but found %s", expectedStatus, status)
	}

	// Since ParseBlock should have triggered a cache hit, assert that the block is identical
	// to the parsed block.
	if getBlk != parsedBlk {
		t.Fatalf("Expected block returned by GetBlock to have been cached, but retrieved non-unique decided block")
	}
}

func checkAcceptedBlock(t *testing.T, s *State, blk snowman.Block, cached bool) {
	checkDecidedBlock(t, s, blk, choices.Accepted, cached)
}

func checkRejectedBlock(t *testing.T, s *State, blk snowman.Block, cached bool) {
	checkDecidedBlock(t, s, blk, choices.Rejected, cached)
}

func TestState(t *testing.T) {
	testBlks := NewTestBlocks(3, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)
	blk1 := testBlks[1]
	blk2 := testBlks[2]
	// Need to create a block with a different bytes and hash here
	// to generate a conflict with blk2
	blk3Bytes := []byte{byte(3)}
	blk3ID := hashing.ComputeHash256Array(blk3Bytes)
	blk3 := &TestBlock{
		TestBlock: &snowman.TestBlock{
			TestDecidable: choices.TestDecidable{
				IDV:     blk3ID,
				StatusV: choices.Processing,
			},
			HeightV: uint64(2),
			BytesV:  blk3Bytes,
			ParentV: blk1,
		},
	}
	testBlks = append(testBlks, blk3)

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	lastAccepted, err := chainState.LastAccepted()
	if err != nil {
		t.Fatal(err)
	}
	if lastAccepted != genesisBlock.ID() {
		t.Fatal("Expected last accepted block to be the genesis block")
	}

	wrappedGenesisBlk, err := chainState.GetBlock(genesisBlock.ID())
	if err != nil {
		t.Fatalf("Failed to get genesis block due to: %s", err)
	}

	// Check that a cache miss on a block is handled correctly
	if _, err := chainState.GetBlock(blk1.ID()); err == nil {
		t.Fatal("expected GetBlock to return an error for blk1 before it's been parsed")
	}
	if _, err := chainState.GetBlock(blk1.ID()); err == nil {
		t.Fatal("expected GetBlock to return an error for blk1 before it's been parsed")
	}

	// Parse and verify blk1 and blk2
	parsedBlk1, err := chainState.ParseBlock(blk1.Bytes())
	if err != nil {
		t.Fatal("Failed to parse blk1 due to: %w", err)
	}
	if err := parsedBlk1.Verify(); err != nil {
		t.Fatal("Parsed blk1 failed verification unexpectedly due to %w", err)
	}
	parsedBlk2, err := chainState.ParseBlock(blk2.Bytes())
	if err != nil {
		t.Fatalf("Failed to parse blk2 due to: %s", err)
	}
	if err := parsedBlk2.Verify(); err != nil {
		t.Fatalf("Parsed blk2 failed verification unexpectedly due to %s", err)
	}

	// Check that the verified blocks have been placed in the processing map
	if numProcessing := len(chainState.verifiedBlocks); numProcessing != 2 {
		t.Fatalf("Expected chain state to have 2 processing blocks, but found: %d", numProcessing)
	}

	parsedBlk3, err := chainState.ParseBlock(blk3.Bytes())
	if err != nil {
		t.Fatalf("Failed to parse blk3 due to %s", err)
	}
	getBlk3, err := chainState.GetBlock(blk3.ID())
	if err != nil {
		t.Fatalf("Failed to get blk3 due to %s", err)
	}
	assert.Equal(t, parsedBlk3.ID(), getBlk3.ID(), "State GetBlock returned the wrong block")

	// Check that parsing blk3 does not add it to processing blocks since it has
	// not been verified.
	if numProcessing := len(chainState.verifiedBlocks); numProcessing != 2 {
		t.Fatalf("Expected State to have 2 processing blocks, but found: %d", numProcessing)
	}

	if err := parsedBlk3.Verify(); err != nil {
		t.Fatalf("Parsed blk3 failed verification unexpectedly due to %s", err)
	}
	// Check that blk3 has been added to processing blocks.
	if numProcessing := len(chainState.verifiedBlocks); numProcessing != 3 {
		t.Fatalf("Expected chain state to have 3 processing blocks, but found: %d", numProcessing)
	}

	// Decide the blocks and ensure they are removed from the processing blocks map
	if err := parsedBlk1.Accept(); err != nil {
		t.Fatal(err)
	}
	if err := parsedBlk2.Accept(); err != nil {
		t.Fatal(err)
	}
	if err := parsedBlk3.Reject(); err != nil {
		t.Fatal(err)
	}

	if numProcessing := len(chainState.verifiedBlocks); numProcessing != 0 {
		t.Fatalf("Expected chain state to have 0 processing blocks, but found: %d", numProcessing)
	}

	// Check that the last accepted block was updated correctly
	lastAcceptedID, err := chainState.LastAccepted()
	if err != nil {
		t.Fatal(err)
	}
	if lastAcceptedID != blk2.ID() {
		t.Fatal("Expected last accepted block to be blk2")
	}
	if lastAcceptedID := chainState.LastAcceptedBlock().ID(); lastAcceptedID != blk2.ID() {
		t.Fatal("Expected last accepted block to be blk2")
	}

	// Flush the caches to ensure decided blocks are handled correctly on cache misses.
	chainState.Flush()
	checkAcceptedBlock(t, chainState, wrappedGenesisBlk, false)
	checkAcceptedBlock(t, chainState, parsedBlk1, false)
	checkAcceptedBlock(t, chainState, parsedBlk2, false)
	checkRejectedBlock(t, chainState, parsedBlk3, false)
}

func TestBuildBlock(t *testing.T) {
	testBlks := NewTestBlocks(2, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)
	blk1 := testBlks[1]

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	buildBlock := func() (snowman.Block, error) {
		// Once the block is built, mark it as processing
		blk1.SetStatus(choices.Processing)
		return blk1, nil
	}

	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          buildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	builtBlk, err := chainState.BuildBlock()
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, chainState.verifiedBlocks, 0)

	if err := builtBlk.Verify(); err != nil {
		t.Fatalf("Built block failed verification due to %s", err)
	}
	assert.Len(t, chainState.verifiedBlocks, 1)

	checkProcessingBlock(t, chainState, builtBlk)

	if err := builtBlk.Accept(); err != nil {
		t.Fatalf("Unexpected error while accepting built block %s", err)
	}

	checkAcceptedBlock(t, chainState, builtBlk, true)
}

func TestStateDecideBlock(t *testing.T) {
	testBlks := NewTestBlocks(4, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)
	badAcceptBlk := testBlks[1]
	badAcceptBlk.AcceptV = errors.New("this block should fail on Accept")
	badVerifyBlk := testBlks[2]
	badVerifyBlk.VerifyV = errors.New("this block should fail verification")
	badRejectBlk := testBlks[3]
	badRejectBlk.RejectV = errors.New("this block should fail on reject")
	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	// Parse badVerifyBlk (which should fail verification)
	badBlk, err := chainState.ParseBlock(badVerifyBlk.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if err := badBlk.Verify(); err == nil {
		t.Fatal("Bad block should have failed verification")
	}
	// Ensure a block that fails verification is not marked as processing
	assert.Len(t, chainState.verifiedBlocks, 0)

	// Ensure that an error during block acceptance is propagated correctly
	badBlk, err = chainState.ParseBlock(badAcceptBlk.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if err := badBlk.Verify(); err != nil {
		t.Fatal(err)
	}
	assert.Len(t, chainState.verifiedBlocks, 1)

	if err := badBlk.Accept(); err == nil {
		t.Fatal("Block should have errored on Accept")
	}

	// Ensure that an error during block reject is propagated correctly
	badBlk, err = chainState.ParseBlock(badRejectBlk.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	if err := badBlk.Verify(); err != nil {
		t.Fatal(err)
	}
	// Note: an error during block Accept/Reject is fatal, so it is undefined whether
	// the block that failed on Accept should be removed from processing or not. We allow
	// either case here to make this test more flexible.
	if numProcessing := len(chainState.verifiedBlocks); numProcessing > 2 || numProcessing == 0 {
		t.Fatalf("Expected number of processing blocks to be either 1 or 2, but found %d", numProcessing)
	}

	if err := badBlk.Reject(); err == nil {
		t.Fatal("Block should have errored on Reject")
	}
}

func TestStateParent(t *testing.T) {
	testBlks := NewTestBlocks(3, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)
	blk1 := testBlks[1]
	blk2 := testBlks[2]

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	parsedBlk2, err := chainState.ParseBlock(blk2.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	missingBlk1 := parsedBlk2.Parent()
	if status := missingBlk1.Status(); status != choices.Unknown {
		t.Fatalf("Expected status of parent of blk2 to be %s, but found %s", choices.Unknown, status)
	}

	parsedBlk1, err := chainState.ParseBlock(blk1.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	genesisBlkParent := parsedBlk1.Parent()
	checkAcceptedBlock(t, chainState, genesisBlkParent, true)

	parentBlk1 := parsedBlk2.Parent()
	checkProcessingBlock(t, chainState, parentBlk1)
}

func TestGetBlockInternal(t *testing.T) {
	testBlks := NewTestBlocks(1, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	genesisBlockInternal := chainState.LastAcceptedBlockInternal()
	if _, ok := genesisBlockInternal.(*TestBlock); !ok {
		t.Fatalf("Expected LastAcceptedBlockInternal to return a block of type *snowman.TestBlock, but found %T", genesisBlockInternal)
	}
	if genesisBlockInternal.ID() != genesisBlock.ID() {
		t.Fatalf("Expected LastAcceptedBlockInternal to be blk %s, but found %s", genesisBlock.ID(), genesisBlockInternal.ID())
	}

	blk, err := chainState.GetBlockInternal(genesisBlock.ID())
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := blk.(*TestBlock); !ok {
		t.Fatalf("Expected retrieved block to return a block of type *snowman.TestBlock, but found %T", blk)
	}
	if blk.ID() != genesisBlock.ID() {
		t.Fatalf("Expected GetBlock to be blk %s, but found %s", genesisBlock.ID(), blk.ID())
	}
}

func TestGetBlockError(t *testing.T) {
	testBlks := NewTestBlocks(2, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)
	blk1 := testBlks[1]

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	wrappedGetBlock := func(id ids.ID) (snowman.Block, error) {
		blk, err := getBlock(id)
		if err != nil {
			return nil, fmt.Errorf("wrapping error to prevent caching miss: %w", err)
		}
		return blk, nil
	}
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            wrappedGetBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	_, err := chainState.GetBlock(blk1.ID())
	if err == nil {
		t.Fatal("Expected GetBlock to return an error for unknown block")
	}

	// Update the status to Processing, so that it will be returned by the internal get block
	// function.
	blk1.SetStatus(choices.Processing)
	blk, err := chainState.GetBlock(blk1.ID())
	if err != nil {
		t.Fatal(err)
	}
	if blk.ID() != blk1.ID() {
		t.Fatalf("Expected GetBlock to retrieve %s, but found %s", blk1.ID(), blk.ID())
	}
	checkProcessingBlock(t, chainState, blk)
}

func TestParseBlockError(t *testing.T) {
	testBlks := NewTestBlocks(1, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	blk, err := chainState.ParseBlock([]byte{255})
	if err == nil {
		t.Fatalf("Expected ParseBlock to return an error parsing an invalid block but found block of type %T", blk)
	}
}

func TestBuildBlockError(t *testing.T) {
	testBlks := NewTestBlocks(1, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	chainState := NewState(&Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	})

	blk, err := chainState.BuildBlock()
	if err == nil {
		t.Fatalf("Expected BuildBlock to return an error but found block of type %T", blk)
	}
}

func TestMeteredCache(t *testing.T) {
	registry := prometheus.NewRegistry()
	namespace1 := "Joe"
	namespace2 := "Namath"

	testBlks := NewTestBlocks(1, nil)
	genesisBlock := testBlks[0]
	genesisBlock.SetStatus(choices.Accepted)

	getBlock, parseBlock, getCanonicalBlockID := createInternalBlockFuncs(t, testBlks)
	config := &Config{
		DecidedCacheSize:    2,
		MissingCacheSize:    2,
		UnverifiedCacheSize: 2,
		LastAcceptedBlock:   genesisBlock,
		GetBlock:            getBlock,
		UnmarshalBlock:      parseBlock,
		BuildBlock:          cantBuildBlock,
		GetBlockIDAtHeight:  getCanonicalBlockID,
	}
	_, err := NewMeteredState(registry, namespace1, config)
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewMeteredState(registry, namespace1, config)
	if err == nil {
		t.Fatal("Expected creating a second NewMeteredState with the same namespace to error due to a registry conflict")
	}
	_, err = NewMeteredState(registry, namespace2, config)
	if err != nil {
		t.Fatal(err)
	}
}
