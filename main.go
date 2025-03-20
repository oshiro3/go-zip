package main

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"os"
	"time"
)

func main() {
	path := "test.txt"
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	h := header{
		id1:   0x1f,
		id2:   0x8b,
		cm:    0x08,
		flag:  flag{fname: true}.toByte(),
		mtime: uint32(time.Now().UnixNano()),
		xfl:   0x00,
		os:    0x03,
	}
	fname := []byte("test.txt")
	fname = append(fname, 0x00)

	var compressed bytes.Buffer
	if err := compless(content, compressed); err != nil {
		panic(err)
	}

	crc32q := crc32.IEEETable
	f := footer{
		crc32: crc32.Checksum([]byte(content), crc32q),
		isize: uint32(len(content)),
	}

	var b bytes.Buffer
	if err := binary.Write(&b, binary.LittleEndian, h); err != nil {
		panic(err)
	}

	binary.Write(&b, binary.LittleEndian, fname)
	binary.Write(&b, binary.LittleEndian, compressed.Bytes())
	binary.Write(&b, binary.LittleEndian, f)

	output_path := "gen_test.txt.gz"
	if err := os.WriteFile(output_path, b.Bytes(), 0644); err != nil {
		panic(err)
	}
}
