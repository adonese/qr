package main

import (
	"hash/crc32"

	"github.com/npat-efault/crc16"
)

func computeCRC(data []byte, poly uint32) uint32 {
	tab := crc32.MakeTable(poly)
	return crc32.Checksum(data, tab)
}

func computeCRC16(data []byte) uint16 {
	conf := &crc16.Conf{
		Poly: 0x1021, BitRev: true,
		IniVal: 0xffff, FinVal: 0xffff,
		BigEnd: true,
	}
	return crc16.Checksum(conf, data)
}
