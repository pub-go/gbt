package model

// Trackers 追踪器列表
type Trackers struct {
	AnnounceList [][]string
	// [ [a, b, c], [back1, back2] ]
	// tiers 表示等级索引 当前使用第一等级 [a,b,c] 还是第二等级[back1, back2]
	// index 表示当前等级追踪器索引
	tiers, index int
}

func (t *Trackers) Reset() {
	t.tiers, t.index = 0, 0
}

func (t *Trackers) Tiers() int {
	return t.tiers
}

func (t *Trackers) Next() string {
	return t.AnnounceList[t.tiers][t.index]
}

func (t *Trackers) MarkCurrentFail() {
	tiers := t.AnnounceList[t.tiers]
	t.index++
	if t.index >= len(tiers) {
		t.tiers++
		t.tiers = t.tiers % len(t.AnnounceList)
		t.index = 0
	}
}
