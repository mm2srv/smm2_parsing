package smm2_parsing

import (
	"bytes"
	"encoding/binary"
)

type Random struct {
	s0 uint32
	s1 uint32
	s2 uint32
	s3 uint32
}

func (r *Random) u32() uint32 {
	// https://github.com/kinnay/NintendoClients/blob/d41d394065900c7214906b5a87da35f561b16763/nintendo/sead/random.py
	temp := r.s0
	temp = (temp ^ (temp << 11)) & 0xFFFFFFFF
	temp ^= temp >> 8
	temp ^= r.s3
	temp ^= r.s3 >> 19
	r.s0 = r.s1
	r.s1 = r.s2
	r.s2 = r.s3
	r.s3 = temp
	return temp
}

func (r *Random) uint(max uint32) uint64 {
	return (uint64(r.u32()) * uint64(max)) >> 32
}

func createKey(r *Random, table []uint32, size uint32, writer *bytes.Buffer) {
	for i := uint32(0); i < size/4; i++ {
		value := uint32(0)
		for e := 0; e < 4; e++ {
			index := r.uint(uint32(len(table)))
			shift := r.uint(4) * 8
			b := (table[index] >> shift) & 0xFF
			value = (value << 8) | b
		}
		binary.Write(writer, binary.LittleEndian, uint32(value))
	}
}
