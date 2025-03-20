package main

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"hash/crc32"
	"os"
	"time"
)

type header struct {
	id1   uint8
	id2   uint8
	cm    uint8
	flag  uint8
	mtime uint32
	xfl   uint8
	os    uint8
}

type footer struct {
	crc32 uint32
	isize uint32
}

type flag struct {
	ftext     bool
	fhcrc     bool
	fextra    bool
	fname     bool
	fcomment  bool
	reserved1 bool
	reserved2 bool
	reserved3 bool
}

func (f flag) toByte() uint8 {
	var b byte = 0
	if f.ftext {
		b |= 1 << 0
	}
	if f.fhcrc {
		b |= 1 << 1
	}
	if f.fextra {
		b |= 1 << 2
	}
	if f.fname {
		b |= 1 << 3
	}
	if f.fcomment {
		b |= 1 << 4
	}
	if f.reserved1 {
		b |= 1 << 5
	}
	if f.reserved2 {
		b |= 1 << 6
	}
	if f.reserved3 {
		b |= 1 << 7
	}
	return b
}

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
	w, err := flate.NewWriter(&compressed, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}
	w.Write(content)
	w.Close()

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
