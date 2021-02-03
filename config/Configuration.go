package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

var defaultPath = "config.yml"

type Config struct {
	FilePath string
}

func GetCfg() *Config {
	return &Config{
		FilePath: defaultPath,
	}
}

//设置配置路径
func (cfg *Config) SetConfigPath(path string) {
	cfg.FilePath = path
}

//加载配置
func (cfg *Config) LoadConfig() {
	viper.SetConfigType("yml")
	viper.SetConfigFile(cfg.FilePath)
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件失败, 异常信息 : ", err)
	}
}

//获取字符串类型配置
func (cfg *Config) GetString(key string) string {
	return fmt.Sprintf("%v", viper.Get(key))
}

//获取Uint64类型的配置
func (cfg *Config) GetUInt64(key string) uint64 {
	intervalue, err := strconv.Atoi(fmt.Sprintf("%v", viper.Get(key)))
	if err != nil {
		panic(err)
	}
	return uint64(intervalue)
}

//获取int类型配置
func (cfg *Config) GetInt(key string) int {
	intervalue, err := strconv.Atoi(fmt.Sprintf("%v", viper.Get(key)))
	if err != nil {
		panic(err)
	}
	return intervalue
}
