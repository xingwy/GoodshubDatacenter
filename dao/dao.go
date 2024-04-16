package dao

import (
	"context"
	"database/sql/driver"
	"goodshub-datacenter/conf"
	"goodshub-datacenter/constants"
	"goodshub-datacenter/model/base"

	"github.com/go-redis/redis/v8"

	gorm "gorm.io/gorm"
)

type Dao struct {
	db    *gorm.DB
	redis *redis.Client
	ch    driver.Conn

	// Redis
	redisMangementInstance *RedisMangementInstance
}

// NewDao 创建连接
func NewDao(cfg *conf.Config) (d *Dao, err error) {
	// db, err := DBConnect(cfg.DB)
	// if err != nil {
	// 	return nil, err
	// }

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:         cfg.Redis.Addr,
	// 	DB:           0,
	// 	Password:     cfg.Redis.Auth,
	// 	WriteTimeout: time.Duration(time.Millisecond * 600),
	// 	ReadTimeout:  time.Duration(time.Millisecond * 300),
	// 	DialTimeout:  time.Duration(time.Minute * 3),
	// })

	// d = &Dao{
	// 	db:    db,
	// 	redis: rdb,
	// }

	// d.redisMangementInstance = &RedisMangementInstance{D: d}

	return
}

func (d *Dao) _getDBFromContext(ctx context.Context) (res *gorm.DB) {
	// 从上下文中获取事务对象
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				res = d.db
			}
		}()
	}()

	if ctx == nil {
		return d.db
	}

	tx := ctx.Value(constants.TRANSACTION_KEY)
	if tx == nil {
		return d.db
	}
	// 提交事务
	return tx.(*gorm.DB)
}

func (d *Dao) RecordCount(ctx context.Context, tableName string, cond *base.DBConditions) (int64, error) {
	var totalCount int64
	if cond != nil {
		err := cond.Fill(d._getDBFromContext(ctx).Table(tableName)).Count(&totalCount).Error
		if err != nil {
			return 0, err
		}
	} else {
		if err := d._getDBFromContext(ctx).Table(tableName).Count(&totalCount).Error; err != nil {
			return 0, err
		}
	}

	return totalCount, nil
}

func (d *Dao) Redis() *redis.Client {
	return d.redis
}

func (d *Dao) DB() *gorm.DB {
	return d.db
}

func (d *Dao) CH() driver.Conn {
	return d.ch
}

func (d *Dao) RedisMangement() RedisMangement {
	return d.redisMangementInstance
}
