package service

import (
	"io"

	"code.gopub.tech/bencode"
)

func ReadMeta(r io.Reader) (meta bencode.Value, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return
	}
	meta, err = bencode.Decode(data)
	return
}
