package service_test

import (
	"os"
	"testing"

	"code.gopub.tech/assert"
	"code.gopub.tech/gbt/service"
)

func TestReadMeta(t *testing.T) {
	f, err := os.Open("./testdata/KNOPPIX_V9.1CD-2021-01-25-EN.torrent")
	assert.Nil(t, err)
	meta, err := service.ReadMeta(f)
	assert.Nil(t, err)
	// t.Logf("%v", meta)
	t.Logf("%v", meta.Announce())
	t.Logf("%v", meta.AnnounceList())
	t.Logf("%v", meta.Comment())
	t.Logf("%v", meta.CreationDateTime())
	info:= meta.Info()
	t.Logf("%v",info.Name())
	t.Logf("%v",info.PieceLength())
	t.Logf("%v",info.Files())

}
