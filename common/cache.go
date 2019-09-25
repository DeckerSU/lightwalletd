package common

import (
	"bytes"

	"github.com/adityapk00/lightwalletd/parser"
	"github.com/pkg/errors"
)

type BlockCache struct {
	MaxEntries int

	FirstBlock int
	LastBlock  int

	m map[int]*parser.Block
}

func New(maxEntries int) *BlockCache {
	return &BlockCache{
		MaxEntries: maxEntries,
		FirstBlock: -1,
		LastBlock:  -1,
		m:          make(map[int]*parser.Block),
	}
}

func (c *BlockCache) Add(height int, block *parser.Block) error {
	println("Cache add", height)
	if c.FirstBlock == -1 && c.LastBlock == -1 {
		// If this is the first block, prep the data structure
		c.FirstBlock = height
		c.LastBlock = height - 1
	} else if height >= c.FirstBlock && height <= c.LastBlock {
		// Overwriting an existing entry. If so, then remove all
		// subsequent blocks, since this might be a reorg
		for i := height; i <= c.LastBlock; i++ {
			delete(c.m, i)
		}
		c.LastBlock = height - 1
	}

	if height != c.LastBlock+1 {
		return errors.New("Blocks need to be added sequentially")
	}

	if c.m[height-1] != nil && !bytes.Equal(block.GetPrevHash(), c.m[height-1].GetEncodableHash()) {
		return errors.New("Prev hash of the block didn't match")
	}

	// Add the entry and update the counters
	c.m[height] = block

	c.LastBlock = height

	// If the cache is full, remove the oldest block
	if c.LastBlock-c.FirstBlock+1 > c.MaxEntries {
		delete(c.m, c.FirstBlock)
		c.FirstBlock = c.FirstBlock + 1
	}

	return nil
}

func (c *BlockCache) Get(height int) *parser.Block {
	println("Cache get", height)
	if c.LastBlock == -1 || c.FirstBlock == -1 {
		return nil
	}

	if height < c.FirstBlock || height > c.LastBlock {
		return nil
	}

	println("Cache returned")
	return c.m[height]
}
