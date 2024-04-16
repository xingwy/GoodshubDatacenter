package define

import (
	"fmt"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Config 服务器配置
type GinConfig struct {
	Name             string        `yaml:"name" json:"name"`       // 服务名称，必填
	Version          string        `yaml:"version" json:"version"` // 服务版本，必填
	Host             string        `yaml:"host" json:"host"`       // 域名主机
	Port             int64         `yaml:"port" json:"port"`
	BroadcastIP      string        `yaml:"broadcastIP" json:"broadcastIP"`               // 广播的运行地址，默认为：IP
	BroadcastPort    int           `yaml:"broadcastPort" json:"broadcastPort"`           // 广播的运行端口，默认为：Port
	Timeout          int           `yaml:"timeout" json:"timeout"`                       // 优雅退出时的超时机制
	Debug            bool          `yaml:"debug" json:"debug"`                           // 是否开启调试
	Pprof            bool          `yaml:"pprof" json:"pprof"`                           // 是否监控性能
	ReadTimeout      time.Duration `yaml:"readtimeout" json:"readtimeout"`               // 读超时
	WriteTimeout     time.Duration `yaml:"writetimeout" json:"writetimeout"`             // 写超时
	DisableAccessLog bool          `yaml:"disable_access_log" json:"disable_access_log"` // disable_access_log
	OpenLimit        bool          `yaml:"openLimit" json:"openLimit"`                   //打开限流器
	UseMicroRandom   bool          `yaml:"useMicroRandom" json:"useMicroRandom"`         //使用轮询选择器
	UseOldSelected   bool          `yaml:"useOldSelected" json:"useOldSelected"`         //使用旧的micro缓存器，可能会增加带宽
}

// Addr 运行地址
func (i *GinConfig) Addr() string {
	return fmt.Sprintf("%s:%d", i.Host, i.Port)
}

// BroadcastAddr 广播的运行地址
func (i *GinConfig) BroadcastAddr() string {
	return fmt.Sprintf("%s:%d", i.BroadcastIP, i.BroadcastPort)
}

type GormConfig struct {
	SkipDefaultTransaction                   bool             `yaml:"skipDefaultTransaction"`
	NamingStrategy                           schema.Namer     `yaml:"namingStrategy"`
	FullSaveAssociations                     bool             `yaml:"fullSaveAssociations"`
	Logger                                   logger.Interface `yaml:"logger"`
	NowFunc                                  func() time.Time `yaml:"nowFunc"`
	DryRun                                   bool             `yaml:"dryRun"`
	PrepareStmt                              bool             `yaml:"prepareStmt"`
	DisableAutomaticPing                     bool             `yaml:"disableAutomaticPing"`
	DisableForeignKeyConstraintWhenMigrating bool             `yaml:"disableForeignKeyConstraintWhenMigrating"`
	DisableNestedTransaction                 bool             `yaml:"disableNestedTransaction"`
	AllowGlobalUpdate                        bool             `yaml:"allowGlobalUpdate"`
	QueryFields                              bool             `yaml:"queryFields"`
	CreateBatchSize                          int              `yaml:"createBatchSize"`
}

type DBConfig struct {
	Alias        string        `yaml:"alias" json:"alias"`
	Type         string        `yaml:"type" json:"type"`
	Server       string        `yaml:"server" json:"server"`
	Port         int           `yaml:"port" json:"port"`
	Database     string        `yaml:"database" json:"database"`
	User         string        `yaml:"user" json:"user"`
	Password     string        `yaml:"password" json:"password"`
	MaxIdleConn  int           `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConn  int           `yaml:"maxOpenConns" json:"maxOpenConns"`
	Charset      string        `yaml:"charset" json:"charset"`
	TimeZone     string        `yaml:"timezone" json:"timezone"`
	MaxLeftTime  time.Duration `yaml:"maxLeftTime" json:"maxLeftTime"`
	ReadTimeout  time.Duration `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout" json:"writeTimeout"`
	Timeout      time.Duration `yaml:"timeout" json:"timeout"`
	Option       *GormConfig   `yaml:"option" json:"option"`
}

type RedisConfig struct {
	Active       int           `yaml:"active" json:"active"`
	Idle         int           `yaml:"idle" json:"idle"`
	IdleTimeout  time.Duration `yaml:"idleTimeout" json:"idleTimeout"`
	WaitTimeout  time.Duration `yaml:"waitTimeout" json:"waitTimeout"`
	Wait         bool          `yaml:"wait" json:"wait"`
	Name         string        `yaml:"name" json:"name"`
	Proto        string        `yaml:"proto" json:"proto"`
	Addr         string        `yaml:"addr" json:"addr"`
	Auth         string        `yaml:"auth" json:"auth"`
	DialTimeout  time.Duration `yaml:"dialTimeout" json:"dialTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout" json:"writeTimeout"`
	DB           int           `yaml:"db" json:"db"`
	SlowLog      time.Duration `yaml:"slowLog" json:"slowLog"`
}

type OpenTaoBao struct {
	AppKey          string `yaml:"appKey" json:"appKey"`
	AppSecret       string `yaml:"appSecret" json:"appSecret"`
	ServerUrl       string `yaml:"serverUrl" json:"serverUrl"`
	ConnectTimeount int64  `yaml:"connectTimeout" json:"connectTimeout"`
	ReadTimeout     int64  `yaml:"readTimeout" json:"readTimeout"`
}

type Memory struct {
	Flag   bool `yaml:"flag" json:"flag"`
	LruMax int  `yaml:"lruMax" json:"lruMax"`
}
