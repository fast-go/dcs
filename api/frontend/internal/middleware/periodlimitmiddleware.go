package middleware

import (
	"dcs/api/frontend/internal/config"
	"dcs/common/define"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
	"time"
)

type PeriodLimitMiddleware struct {
	Config   config.Config
	Seconds  int //窗口时间
	Quota    int //配额
	RedisKey string
	Store    *redis.Redis
}

func NewPeriodLimitMiddleware(c config.Config) *PeriodLimitMiddleware {
	store := redis.New("localhost:6379", redis.WithPass(""))
	if !store.Ping() {
		panic("redis err")
	}
	return &PeriodLimitMiddleware{
		Config:   c,
		Seconds:  cast.ToInt(time.Hour * 24),
		Quota:    50,
		RedisKey: define.PeriodLimitOrder,
		Store:    store,
	}
}

func (m *PeriodLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lmt := limit.NewPeriodLimit(m.Seconds, m.Quota, m.Store, m.RedisKey)
		if v, err := lmt.Take(strconv.FormatInt(cast.ToInt64(r.Context().Value("userId")), 10)); err == nil && (v == limit.Allowed || v == limit.HitQuota) {
			next(w, r)
		}
		return
	}
}
