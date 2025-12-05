package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixRange(t *testing.T) {
	ranges := []Range{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   14,
		},
		{
			start: 16,
			end:   20,
		},
		{
			start: 12,
			end:   18,
		},
	}
	ranges = mergeRanges(ranges)

	expected := []Range{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   20,
		},
	}

	assert.Equal(t, expected, ranges)
	total := countRanges(ranges)
	assert.Equal(t, 14, total)
}

func TestMergeRanges2(t *testing.T) {
	ranges := []Range{
		{
			start: 5,
			end:   10,
		},
		{
			start: 5,
			end:   9,
		},
	}

	ranges = mergeRanges(ranges)

	expected := []Range{
		{
			start: 5,
			end:   10,
		},
	}

	assert.Equal(t, expected, ranges)

	total := countRanges(ranges)
	assert.Equal(t, 6, total)
}

func TestMergeRanges3(t *testing.T) {
	ranges := []Range{
		{
			start: 5,
			end:   10,
		},
		{
			start: 5,
			end:   9,
		},
		{
			start: 6,
			end:   8,
		},
	}

	ranges = mergeRanges(ranges)

	expected := []Range{
		{
			start: 5,
			end:   10,
		},
	}

	assert.Equal(t, expected, ranges)

	total := countRanges(ranges)
	assert.Equal(t, 6, total)
}

func TestMergeRanges4(t *testing.T) {
	ranges := []Range{
		{
			start: 5,
			end:   10,
		},
		{
			start: 5,
			end:   9,
		},
		{
			start: 5,
			end:   8,
		},
		{
			start: 5,
			end:   6,
		},
	}

	ranges = mergeRanges(ranges)

	expected := []Range{
		{
			start: 5,
			end:   10,
		},
	}

	assert.Equal(t, expected, ranges)

	total := countRanges(ranges)
	assert.Equal(t, 6, total)
}
