package dao

import (
	"fmt"
	"goodshub-datacenter/model/base"
	"goodshub-datacenter/struct/slog"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

var (
	defaultDatabase     = "mysql"
	ConnStrTmpl         = "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s"
	DefaultMaxOpenConns = 200
	DefaultMaxIdleConns = 60
	DefaultMaxLeftTime  = 300 * time.Second
	DefaultTimeout      = 5 * time.Second
	DefaultWriteTimeout = 5 * time.Second
	DefaultReadTimeout  = 10 * time.Second
	Charset             = "utf8mb4"
	MPort               = 3306
	TimeZone            = "Local"
)

// New 创建一个gorm.DB
func DBConnect(conf *base.DBConfig) (*gorm.DB, error) {
	err := authConfig(conf)
	if err != nil {
		panic(err)
	}
	connStr := fmt.Sprintf(
		ConnStrTmpl,
		conf.User,
		conf.Password,
		conf.Server,
		conf.Port,
		conf.Database,
		conf.Charset,
		conf.TimeZone,
		conf.Timeout,
		conf.ReadTimeout,
		conf.WriteTimeout)

	option := &gorm.Config{
		Logger: slog.NewLogger(),
	}

	if conf.Option != nil {
		option.SkipDefaultTransaction = conf.Option.SkipDefaultTransaction
		option.NamingStrategy = conf.Option.NamingStrategy
		option.FullSaveAssociations = conf.Option.FullSaveAssociations
		option.NowFunc = conf.Option.NowFunc
		option.DryRun = conf.Option.DryRun
		option.PrepareStmt = conf.Option.PrepareStmt
		option.DisableAutomaticPing = conf.Option.DisableAutomaticPing
		option.DisableForeignKeyConstraintWhenMigrating = conf.Option.DisableForeignKeyConstraintWhenMigrating
		option.DisableNestedTransaction = conf.Option.DisableNestedTransaction
		option.AllowGlobalUpdate = conf.Option.AllowGlobalUpdate
		option.QueryFields = conf.Option.QueryFields
		option.CreateBatchSize = conf.Option.CreateBatchSize
	}

	DBConnect, err := gorm.Open(mysql.Open(connStr), option)
	if err != nil {
		return nil, err
	}
	DBInstance, err := DBConnect.DB()
	if err != nil {
		return nil, err
	}
	DBInstance.SetConnMaxLifetime(conf.MaxLeftTime)
	DBInstance.SetMaxIdleConns(conf.MaxIdleConn)
	DBInstance.SetMaxOpenConns(conf.MaxOpenConn)

	return DBConnect, nil
}

func authConfig(conf *base.DBConfig) (err error) {

	if len(conf.Type) == 0 {
		conf.Type = defaultDatabase
	}

	if conf.Port == 0 {
		conf.Port = MPort
	}

	if len(conf.User) == 0 || len(conf.Password) == 0 {
		err = fmt.Errorf("User or  Password is empty")
		return
	}

	if len(conf.Server) == 0 {
		err = fmt.Errorf("server addr is empty")
		return
	}

	if len(conf.Database) == 0 {
		err = fmt.Errorf("database is empty")
		return
	}

	if conf.MaxIdleConn == 0 {
		conf.MaxIdleConn = DefaultMaxIdleConns
	}

	if conf.MaxLeftTime == 0 {
		conf.MaxLeftTime = DefaultMaxLeftTime
	}

	if conf.MaxOpenConn == 0 {
		conf.MaxOpenConn = DefaultMaxOpenConns
	}

	if strings.TrimSpace(conf.Charset) == "" {
		conf.Charset = Charset
	}

	if strings.TrimSpace(conf.TimeZone) == "" {
		conf.TimeZone = TimeZone
	}

	if conf.Timeout == 0 {
		conf.Timeout = DefaultTimeout
	}
	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = DefaultReadTimeout
	}
	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = DefaultWriteTimeout
	}

	return
}
