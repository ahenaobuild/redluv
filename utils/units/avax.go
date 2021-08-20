// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package units

// Denominations of value
const (
	NanoLuv   uint64 = 1
	MicroLuv  uint64 = 1000 * NanoLuv
	Schmeckle uint64 = 49*MicroLuv + 463*NanoLuv
	MilliLuv  uint64 = 1000 * MicroLuv
	Luv       uint64 = 1000 * MilliLuv
	KiloLuv   uint64 = 1000 * Luv
	MegaLuv   uint64 = 1000 * KiloLuv
)
