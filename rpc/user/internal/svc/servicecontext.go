package svc

import (
	"dcs/rpc/user/dao"
	"dcs/rpc/user/internal/config"
	"fmt"
	"github.com/8treenet/gcache"
	"github.com/8treenet/gcache/option"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Dao    *dao.Instance
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open("mysql", c.Mysql.Dns)
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	opt := option.DefaultOption{}
	opt.Expires = 30               //缓存时间, 默认120秒。范围30-43200
	opt.Level = option.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = false         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存。 ps: affected如果未0，不触发更新。
	opt.PenetrationSafe = false    //开启防穿透, 默认false。 ps:防击穿强制全局开启。

	//缓存中间件附加到gorm.DB
	gcache.AttachDB(db, &opt, &option.RedisOption{Addr: c.Redis.Host})
	db.LogMode(true)
	//cachePlugin.FlushDB()
	//cachePlugin.Debug()

	return &ServiceContext{
		Config: c,
		DB:     db,
		Dao: &dao.Instance{
			User: dao.NewUserDao(db),
		},
	}
}
