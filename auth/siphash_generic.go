// Copyright (c) 2017 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package auth

import (
	"encoding/binary"
)

func siphashCoreGeneric(hVal *[4]uint64, msg []byte) {
	v0, v1, v2, v3 := hVal[0], hVal[1], hVal[2], hVal[3]

	for len(msg) >= BlockSize {
		m := binary.LittleEndian.Uint64(msg)
		msg = msg[BlockSize:]

		v3 ^= m

		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)
		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2
		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0
		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)

		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)
		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2
		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0
		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)

		v0 ^= m
	}
	hVal[0], hVal[1], hVal[2], hVal[3] = v0, v1, v2, v3
}

func siphashFinalizeGeneric(tag *[TagSize]byte, hVal *[4]uint64, buf *[8]byte) {
	v0, v1, v2, v3 := hVal[0], hVal[1], hVal[2], hVal[3]

	m := binary.LittleEndian.Uint64(buf[:])

	v3 ^= m

	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)
	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2
	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0
	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	v0 += v1
	v1 = v1<<13 | v1>>(64-13)
	v1 ^= v0
	v0 = v0<<32 | v0>>(64-32)
	v2 += v3
	v3 = v3<<16 | v3>>(64-16)
	v3 ^= v2
	v0 += v3
	v3 = v3<<21 | v3>>(64-21)
	v3 ^= v0
	v2 += v1
	v1 = v1<<17 | v1>>(64-17)
	v1 ^= v2
	v2 = v2<<32 | v2>>(64-32)

	v0 ^= m

	v2 ^= 0xee
	for i := 0; i < 4; i++ {
		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)
		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2
		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0
		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)
	}
	binary.LittleEndian.PutUint64(tag[:], v0^v1^v2^v3)

	v1 ^= 0xdd
	for i := 0; i < 4; i++ {
		v0 += v1
		v1 = v1<<13 | v1>>(64-13)
		v1 ^= v0
		v0 = v0<<32 | v0>>(64-32)
		v2 += v3
		v3 = v3<<16 | v3>>(64-16)
		v3 ^= v2
		v0 += v3
		v3 = v3<<21 | v3>>(64-21)
		v3 ^= v0
		v2 += v1
		v1 = v1<<17 | v1>>(64-17)
		v1 ^= v2
		v2 = v2<<32 | v2>>(64-32)
	}
	binary.LittleEndian.PutUint64(tag[8:], v0^v1^v2^v3)
}
