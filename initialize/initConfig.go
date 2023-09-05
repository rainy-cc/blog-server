package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
}

var Mysql = MysqlConfig{}

func InitConfig() {
	v := viper.New()
	v.AddConfigPath("./conf")
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	readConfigErr := v.ReadInConfig() // Find and read the config file
	if readConfigErr != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("读取配置失败: %v", readConfigErr))
	}
	unmarshalErr := v.UnmarshalKey("mysql", &Mysql)
	if unmarshalErr != nil {
		panic(fmt.Errorf("转换配置失败: %v", unmarshalErr))
	}
}
