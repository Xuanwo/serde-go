package serde

import (
	"math"
	"math/bits"
)

const (
	UintSize = bits.UintSize

	MinInt   = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxInt   = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt8  = math.MinInt8
	MaxInt8  = math.MaxInt8
	MinInt16 = math.MinInt16
	MaxInt16 = math.MaxInt16
	MinInt32 = math.MinInt32
	MaxInt32 = math.MaxInt32
	MinInt64 = math.MinInt64
	MaxInt64 = math.MaxInt64

	MinUint   = 0
	MaxUint   = 1<<UintSize - 1 // 1<<32 - 1 or 1<<64 - 1
	MinUint8  = 0
	MaxUint8  = math.MaxUint8
	MinUint16 = 0
	MaxUint16 = math.MaxUint16
	MinUint32 = 0
	MaxUint32 = math.MaxUint32
	MinUint64 = 0
	MaxUint64 = math.MaxUint64
)
