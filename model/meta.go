package model

import (
	"time"

	"code.gopub.tech/bencode"
)

type Meta bencode.Dict

func (m Meta) Announce() string {
	return bencode.AsStr(m["announce"])
}

func (m Meta) AnnounceList() (list []string) {
	for _, item := range bencode.AsList(m["announce-list"]) {
		if l := bencode.AsList(item); len(l) == 1 {
			if s, ok := l[0].(bencode.String); ok {
				list = append(list, string(s))
			}
		}
	}
	return
}

func (m Meta) Comment() string {
	return bencode.AsStr(m["comment"])
}

func (m Meta) CreationDate() int64 {
	return bencode.AsInt(m["creation date"])
}

func (m Meta) CreationDateTime() time.Time {
	sec := m.CreationDate()
	return time.Unix(sec, 0)
}

func (m Meta) Info() Info {
	return Info(bencode.AsDict(m["info"]))
}

type Info bencode.Dict

func (i Info) Name() string {
	return bencode.AsStr(i["name"])
}

func (i Info) PieceLength() int64 {
	return bencode.AsInt(i["piece length"])
}

func (i Info) Files() (files []File) {
	for _, file := range bencode.AsList(i["files"]) {
		if f, ok := file.(bencode.Dict); ok {
			files = append(files, File(f))
		}
	}
	return
}

type File bencode.Dict

func (f File) Length() int64 {
	return bencode.AsInt(f["length"])
}

func (f File) Path() (path []string) {
	list := bencode.AsList(f["path"])
	for _, item := range list {
		if p, ok := item.(bencode.String); ok {
			path = append(path, string(p))
		}
	}
	return
}
