package providers

import (
	"crypto/rand"
	"math"
	"math/big"
)

func Uint() uint {
	const maxRandInt = 100_000_000
	num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))

	return uint(num.Int64()) + 1
}

func Uint64() uint64 {
	const maxRandInt = math.MaxInt64
	num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))

	return uint64(num.Int64()) + 1
}

func Uint32() uint32 {
	const maxRandInt = math.MaxUint32
	num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))

	return uint32(num.Int64()) + 2
}

func Uint16() uint16 {
	const maxRandInt = math.MaxUint16
	num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))

	return uint16(num.Int64()) + 1
}

func Uint8() uint8 {
	const maxRandInt = math.MaxUint8
	num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))

	return uint8(num.Int64()) + 1
}

func UintSlice() []uint {
	const maxRandInt = math.MaxInt

	var (
		result           = []uint{}
		maxIterations, _ = rand.Int(rand.Reader, big.NewInt(10))
	)
	for maxIterations.Int64() > 0 {
		maxIterations.Sub(maxIterations, big.NewInt(1))
		num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))
		result = append(result, uint(num.Int64())+1)
	}

	return result
}

func Uint64Slice() []uint64 {
	const maxRandInt = math.MaxInt64

	var (
		result           = []uint64{}
		maxIterations, _ = rand.Int(rand.Reader, big.NewInt(10))
	)
	for maxIterations.Int64() > 0 {
		maxIterations.Sub(maxIterations, big.NewInt(1))

		num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))
		result = append(result, uint64(num.Int64())+1)
	}

	return result
}

func Uint32Slice() []uint32 {
	const maxRandInt = math.MaxUint32

	var (
		result           = []uint32{}
		maxIterations, _ = rand.Int(rand.Reader, big.NewInt(10))
	)
	for maxIterations.Int64() > 0 {
		maxIterations.Sub(maxIterations, big.NewInt(1))

		num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))
		result = append(result, uint32(num.Int64())+1)
	}

	return result
}

func Uint16Slice() []uint16 {
	const maxRandInt = math.MaxUint16

	var (
		result           = []uint16{}
		maxIterations, _ = rand.Int(rand.Reader, big.NewInt(3))
	)
	for maxIterations.Int64() > 0 {
		maxIterations.Sub(maxIterations, big.NewInt(1))

		num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))
		result = append(result, uint16(num.Int64())+1)
	}

	return result
}

func Uint8Slice() []uint8 {
	const maxRandInt = math.MaxUint8

	var (
		result           = []uint8{}
		maxIterations, _ = rand.Int(rand.Reader, big.NewInt(2))
	)
	for maxIterations.Int64() > 0 {
		maxIterations.Sub(maxIterations, big.NewInt(1))

		num, _ := rand.Int(rand.Reader, big.NewInt(maxRandInt))
		result = append(result, uint8(num.Int64())+1)
	}

	return result
}
