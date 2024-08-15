package service

import (
	"io"

	"code.gopub.tech/bencode"
	"code.gopub.tech/gbt/model"
)

func ReadMeta(r io.Reader) (meta model.Meta, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return
	}
	val, err := bencode.Decode(data)
	if err != nil {
		return
	}
	meta = model.Meta(bencode.AsDict(val))
	return
}
