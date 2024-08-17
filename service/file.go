package service

import (
	"io"

	"code.gopub.tech/bencode"
	"code.gopub.tech/gbt/model"
)

func ReadMeta(r io.Reader) (meta model.MetaInfo, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return
	}
	val, err := bencode.Decode(data)
	if err != nil {
		return
	}
	meta = model.MetaInfo(bencode.AsDict(val))
	return
}
