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
	t.Logf("%v", meta)
}
