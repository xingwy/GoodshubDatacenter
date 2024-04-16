package conf

import (
	"goodshub-datacenter/model/base"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	//Conf 配置项
	Conf = &Config{}
)

type Config struct {
	ServiceName string            `yaml:"serviceName"`
	GinServer   *base.GinConfig   `yaml:"ginserver"`
	DB          *base.DBConfig    `yaml:"db"`
	Redis       *base.RedisConfig `yaml:"redis"`
	OpenTaoBao  *base.OpenTaoBao  `yaml:"openTaobao"`
}

func init() {

}

func Init() (err error) {
	// 读取 YAML 文件
	yamlFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	// 解析 YAML 数据
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}
	Conf = &config

	return nil
}
