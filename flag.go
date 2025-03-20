package main

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
