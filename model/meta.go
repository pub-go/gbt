package model

import (
	"math/rand"
	"time"

	"code.gopub.tech/bencode"
)

// MetaInfo 种子文件
// see bep_0003
type MetaInfo bencode.Dict

// Announce 追踪器的地址 (一定有的字段)
func (m MetaInfo) Announce() string {
	return bencode.AsStr(m["announce"])
}

// Info 描述种子文件对应的文件(夹) (一定有的字段)
func (m MetaInfo) Info() Info {
	return Info(bencode.AsDict(m["info"]))
}

// Comment 注释
func (m MetaInfo) Comment() string {
	return bencode.AsStr(m["comment"])
}

// CreationDate 创建时间 秒级时间戳
func (m MetaInfo) CreationDate() int64 {
	return bencode.AsInt(m["creation date"])
}

// CreationDateTime 创建时间
func (m MetaInfo) CreationDateTime() time.Time {
	sec := m.CreationDate()
	return time.Unix(sec, 0)
}

// AnnounceList 扩展后的追踪器列表
// 如果客户端兼容该字段 应当优先使用这个字段, 并且忽略 announce 字段.
//
// 各层级的通告将按顺序处理；在客户端进入下一层级之前，必须检查每一层级中的所有 URL。
// 每一层级内的 URL 将以随机选择的顺序进行处理；换句话说，首次读取时列表将被打乱，然后按顺序进行解析。
// 此外，如果与跟踪器的连接成功，它将被移到该层级的前面。
//
//	示例
//
//	d['announce-list'] = [ [tracker1], [backup1], [backup2] ]
//	每次通告时，首先尝试 追踪器 1, 如果无法联系到 追踪器 1, 则分别尝试 备份 1 和 备份 2.
//	在下一次通告时，按同样的顺序重复。这适用于标准跟踪器无法共享信息的情况。
//
//	d['announce-list'] = [[ tracker1, tracker2, tracker3 ]]
//	首先，打乱列表。（为便于讨论，我们假设列表已经被打乱。）
//	然后，如果无法访问 tracker1, 则尝试 tracker2.
//	如果可以访问 tracker2, 那么列表现在是：tracker2, tracker1, tracker3.
//	从那时起，这将是客户端尝试的顺序。
//	如果稍后 tracker2 和 tracker1 都无法访问，但 tracker3 有响应，
//	那么列表将更改为：tracker3, tracker2, tracker1, 并在未来按照该顺序进行尝试。
//	这种形式适用于可以交换对等体信息的跟踪器，并会使客户端帮助平衡跟踪器之间的负载。
//
//	d['announce-list'] = [ [ tracker1, tracker2 ], [backup1] ]
//	第一层由 tracker1 和 tracker2 组成，是随机排列的。
//	在客户端尝试连接 backup1 之前，每次通告都会尝试 tracker1 和 tracker2（尽管顺序可能不同）。
//
// see bep_0012
func (m MetaInfo) AnnounceList() (trackers [][]string) {
	list := bencode.AsList(m["announce-list"])
	trackers = make([][]string, 0, len(list))
	for _, tiers := range list {
		vals := bencode.AsList(tiers)
		urls := make([]string, 0, len(vals))
		for _, url := range vals {
			urls = append(urls, bencode.AsStr(url))
		}
		trackers = append(trackers, urls)
	}
	return
}

// Trackers 获取追踪器列表
func (m MetaInfo) Trackers() *Trackers {
	list := m.AnnounceList()
	announceList := make([][]string, 0, len(list))
	for _, tiers := range list {
		urls := make([]string, 0, len(tiers))
		for _, url := range tiers {
			if url != "" {
				urls = append(urls, url)
			}
		}
		if length := len(urls); length > 0 {
			rand.Shuffle(int(int32(length)), func(i, j int) {
				urls[i], urls[j] = urls[j], urls[i]
			})
			announceList = append(announceList, urls)
		}
	}
	if len(announceList) == 0 {
		announceList = append(announceList, []string{m.Announce()})
	}
	return &Trackers{
		AnnounceList: announceList,
	}
}
