package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go_web_app/settings"
)

var db *sqlx.DB

// Init 用来初始化 MySQL 连接
func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// Connect 方法执行了 Open 和 Ping
	// 也可以用 MustConnect 连接，若不成功，则直接 panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) // 设置最大连接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) // 设置空闲连接数
	return
}

// Close 用来结束与 MySQL 的连接
func Close() {
	_ = db.Close()
}
