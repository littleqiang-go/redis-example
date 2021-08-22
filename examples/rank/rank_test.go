package rank

import (
	"fmt"
	"redis-example/cache"
	"testing"

	"github.com/go-redis/redis"
)

func TestGetRankData(t *testing.T) {
	cache.InitCache(&redis.Options{Addr: "localhost:6379"})

	key := "ranks"
	rank := NewRank(key)
	for i := 1; i <= 10; i++ {
		member := fmt.Sprintf("userid-%d", i)
		score := i * 10
		rank.AddItem(member, float64(score))
	}

	t.Logf("%+v", rank.GetAllRanks(DESC))
}

func TestGetMyRank(t *testing.T) {
	cache.InitCache(&redis.Options{Addr: "localhost:6379"})

	key := "ranks"
	rank := NewRank(key)
	member := "userid-7"
	t.Log(rank.GetSelfRankAsc(member))
	if rank.GetSelfRankAsc(member)+1 == 7 {
		t.Log("suc")
	} else {
		t.Log("fail")
	}

	if rank.GetSelfRankDesc(member)+1 == 4 {
		t.Log("suc")
	} else {
		t.Log("fail")
	}
}
