package model

import "code.gopub.tech/bencode"

type File bencode.Dict

// Length 文件大小(字节数)
func (f File) Length() int64 {
	return bencode.AsInt(f["length"])
}

// Path 路径
func (f File) Path() (path []string) {
	list := bencode.AsList(f["path"])
	for _, item := range list {
		if p, ok := item.(bencode.String); ok {
			path = append(path, string(p))
		}
	}
	return
}
