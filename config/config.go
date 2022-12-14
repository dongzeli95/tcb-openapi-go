package config

import (
	"github.com/dongzeli95/tcb-openapi-go/sts"
	"github.com/gogf/gf/database/gredis"
	"github.com/spf13/viper"
	"strings"
	"time"
)

type Config struct {
	EnvId     string        //TCB 环境 ID
	TcbRegion string        //TCB 环境所属地域
	Timeout   time.Duration //请求超时设置
	LogPrefix string        //日志前缀
	Debug     bool          //debug
	SecretId  string        //访问管理密钥ID
	SecretKey string        //访问管理密钥KEY
	//Deprecated
	StsConfig sts.Config //cam config
	//Deprecated
	RedisConfig gredis.Config //redis config
}

//初始化
func init() {
	//fmt.Println(os.Getwd())
	//viper.SetConfigFile(".env")//默认
	viper.SetConfigFile("../../.env") //适配在component下各个文件夹中的test
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.ReadInConfig()

	//fmt.Println(viper.AllKeys())
	//fmt.Println(viper.AllSettings())
	//fmt.Println(os.Environ())
}
