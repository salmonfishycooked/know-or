package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 是全局变量，用来保存程序的所有配置信息
var Conf = &AppConfig{}

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Init 用来进行项目配置管理库 viper 的初始化
func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	//viper.SetConfigType("yaml")   // 指定配置文件类型（用于从远程获取配置信息时指定配置的类型）
	viper.AddConfigPath(".") // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置文件失败
		return
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		return
	}
	viper.WatchConfig()                            // 开启配置文件监视
	viper.OnConfigChange(func(in fsnotify.Event) { // 配置文件更新后执行的回调函数
		if err := viper.Unmarshal(Conf); err != nil {
			return
		}
		fmt.Println("配置文件已更新......")
	})
	return
}
