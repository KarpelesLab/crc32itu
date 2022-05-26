package crc32itu

import "hash"

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code extracted from golang compress/bzip2

// This is a standard CRC32 like in hash/crc32 except that all the shifts are reversed,
// causing the bits in the input to be processed in the reverse of the usual order.

var crctab [256]uint32

func init() {
	const poly = 0x04C11DB7
	for i := range crctab {
		crc := uint32(i) << 24
		for j := 0; j < 8; j++ {
			if crc&0x80000000 != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
		crctab[i] = crc
	}
}

// updateCRC updates the crc value to incorporate the data in b.
// The initial value is 0.
func updateCRC(val uint32, b []byte) uint32 {
	crc := ^val
	for _, v := range b {
		crc = crctab[byte(crc>>24)^v] ^ (crc << 8)
	}
	return ^crc
}

func New() hash.Hash32 {
	return &crc32digest{}
}

type crc32digest struct {
	crc uint32
}

func (d *crc32digest) Size() int { return 4 }

func (d *crc32digest) BlockSize() int { return 1 }

func (d *crc32digest) Reset() { d.crc = 0 }

func (d *crc32digest) Sum32() uint32 { return d.crc }

func (d *crc32digest) Sum(in []byte) []byte {
	s := d.Sum32()
	return append(in, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func (d *crc32digest) Write(p []byte) (n int, err error) {
	d.crc = updateCRC(d.crc, p)
	return len(p), nil
}
