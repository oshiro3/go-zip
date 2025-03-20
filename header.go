package main

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
