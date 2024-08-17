package model

import (
	"crypto/sha1"
	"fmt"
	"net/url"

	"code.gopub.tech/bencode"
)

// Info
type Info bencode.Dict

// Name 建议保存为的文件(夹)名 实际保存时用户也可另行指定
func (i Info) Name() string {
	return bencode.AsStr(i["name"])
}

// PieceLength 一个分片的长度(字节数) 通常是 2 的幂次方
func (i Info) PieceLength() int64 {
	return bencode.AsInt(i["piece length"])
}

// Pieces 长度是 20 的整数倍; 每 20 位表示一个分片的 SHA1 散列值
func (i Info) Pieces() []byte {
	return []byte(bencode.AsString(i["pieces"]))
}

// IsPrivate 是否是私有种子(只能通过追踪器获取对等方信息)
// see bep_0027
func (i Info) IsPrivate() bool {
	return bencode.AsInt(i["private"]) == 1
}

// IsSingleFile 是否是单文件
func (i Info) IsSingleFile() bool {
	_, ok := i["length"]
	return ok
}

// Length 如果是单文件 表示文件大小(字节数)
func (i Info) Length() int64 {
	return bencode.AsInt(i["length"])
}

// IsMultiFile 是否是多文件
func (i Info) IsMultiFile() bool {
	_, ok := i["files"]
	return ok
}

// Files 如果是多文件 表示目录下的文件
func (i Info) Files() (files []File) {
	for _, file := range bencode.AsList(i["files"]) {
		if f, ok := file.(bencode.Dict); ok {
			files = append(files, File(f))
		}
	}
	return
}

func (i Info) TotalSize() int64 {
	if i.IsSingleFile() {
		return i.Length()
	}
	var sum int64
	for _, file := range i.Files() {
		sum += file.Length()
	}
	return sum
}

// Hash 种子文件的 info_hash 用这个哈希值来识别一个种子
func (i Info) Hash() []byte {
	info := bencode.Dict(i).Encode()
	hash := sha1.Sum(info)
	return hash[:]
}

// HashStr 种子文件的 info_hash 的十六进制表示
func (i Info) HashStr() string {
	return fmt.Sprintf("%x", i.Hash())
}

// HashEscape 种子文件的 info_hash 用在 url 上的表示(使用百分号转义)
func (i Info) HashEscape() string {
	return url.QueryEscape(string(i.Hash()))
}
