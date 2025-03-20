package main

import (
	"bytes"
	"compress/flate"
)

func compless(content []byte, compressed bytes.Buffer) error {
	w, err := flate.NewWriter(&compressed, flate.DefaultCompression)
	if err != nil {
		return err
	}
	w.Write(content)
	w.Close()

	return nil
}
