package rank

import (
	"redis-example/cache"

	"github.com/go-redis/redis"
)

type Order int

const (
	ASC Order = iota
	DESC
)

type Rank struct {
	rKey string
}

type RankItem struct {
	Member string
	Score  float64
}

func NewRank(key string) *Rank {
	return &Rank{rKey: key}
}

// 添加排行item
func (r *Rank) AddItem(member string, score float64) bool {
	return r.AddItemX(&RankItem{Member: member, Score: score})
}

//添加排行item
func (r *Rank) AddItemX(rankItem *RankItem) bool {
	if rankItem == nil {
		return false
	}
	_, err := cache.Cache().ZAdd(r.rKey, redis.Z{Score: rankItem.Score, Member: rankItem.Member}).Result()
	if err != nil {
		return false
	}
	return true
}

// 获拍行榜所有数据
// order 正/逆 序
func (r *Rank) GetAllRanks(order Order) []RankItem {
	if order == DESC {
		return r.GetRanksDesc(0, -1)
	}
	return r.GetRanksAsc(0, -1)
}

// 获取排行榜部分数据 正序
func (r *Rank) GetRanksAsc(start, stop int64) []RankItem {
	zInfo, err := cache.Cache().ZRangeWithScores(r.rKey, start, stop).Result()
	if err != nil {
		return nil
	}
	var rankLst []RankItem
	for i := 0; i < len(zInfo); i++ {
		rankLst = append(rankLst, RankItem{Member: zInfo[i].Member.(string), Score: zInfo[i].Score})
	}
	return rankLst
}

// 获取排行榜部分数据 逆序
func (r *Rank) GetRanksDesc(start, stop int64) []RankItem {
	zInfo, err := cache.Cache().ZRevRangeWithScores(r.rKey, start, stop).Result()
	if err != nil {
		return nil
	}
	var rankLst []RankItem
	for i := 0; i < len(zInfo); i++ {
		rankLst = append(rankLst, RankItem{Member: zInfo[i].Member.(string), Score: zInfo[i].Score})
	}
	return rankLst
}

// 获取自己的排行 正序
func (r *Rank) GetSelfRankAsc(member string) int64 {
	rankNo, err := cache.Cache().ZRank(r.rKey, member).Result()
	if err != nil {
		return -1
	}
	return rankNo
}

// 获取自己的排行 逆序
func (r *Rank) GetSelfRankDesc(member string) int64 {
	rankNo, err := cache.Cache().ZRevRank(r.rKey, member).Result()
	if err != nil {
		return -1
	}
	return rankNo
}
