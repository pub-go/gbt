package model_test

import (
	"testing"

	"code.gopub.tech/assert"
	"code.gopub.tech/gbt/model"
)

func TestTrackers(t *testing.T) {
	var trackers = &model.Trackers{
		AnnounceList: [][]string{
			{"tracker1", "tracker2"},
			{"backup1", "backup2"},
		},
	}
	assert.Equal(t, trackers.Next(), "tracker1")
	trackers.MarkCurrentFail()
	assert.Equal(t, trackers.Next(), "tracker2")
	trackers.MarkCurrentFail()
	assert.Equal(t, trackers.Next(), "backup1")
	trackers.MarkCurrentFail()
	assert.Equal(t, trackers.Next(), "backup2")
	trackers.MarkCurrentFail()
	assert.Equal(t, trackers.Next(), "tracker1")
}
