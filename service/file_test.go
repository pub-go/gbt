package service_test

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"code.gopub.tech/assert"
	"code.gopub.tech/gbt/service"
)

func TestReadMeta(t *testing.T) {
	// f, err := os.Open("./testdata/ubuntu-24.04-desktop-amd64.iso.torrent")
	f, err := os.Open("./testdata/KNOPPIX_V9.1CD-2021-01-25-EN.torrent")
	assert.Nil(t, err)
	meta, err := service.ReadMeta(f)
	assert.Nil(t, err)
	// t.Logf("%v", meta)
	t.Logf("Announce=%v", meta.Announce())
	t.Logf("AnnounceList=%v", meta.AnnounceList())
	t.Logf("Comment=%v", meta.Comment())
	t.Logf("CreationDateTime=%v", meta.CreationDateTime())
	info := meta.Info()
	t.Logf("Name=%v", info.Name())
	// https://torrent.ubuntu.com/file?info_hash=%2A%A4%F5%A7%E2%09%E5K2%80%3DCg%09q%C4%C8%CA%AA%05
	// %2A%A4%F5%A7%E2%09%E5K2%80%3DCg%09q%C4%C8%CA%AA%05
	t.Logf("info_hash=%v, %v, %v", info.Hash(), info.HashStr(), info.HashEscape())
	t.Logf("PieceLength=%v", info.PieceLength())
	var size int64
	for _, file := range info.Files() {
		length := file.Length()
		size += length
		t.Logf("-- Length=%v", length)
		t.Logf("-- Path=%v", file.Path())
	}
	t.Logf("length=%v, size=%v", info.Length(), size)
	t.Logf("IsSingleFile=%v, totalSize=%v", info.IsSingleFile(), info.TotalSize())
	t.Logf("AnnounceList=%v", meta.Trackers().AnnounceList)
	//t.Logf("%s", info.Pieces())
	t.Logf("%s", generateRandomString(20))
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
