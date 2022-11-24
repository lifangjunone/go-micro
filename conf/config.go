package conf

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var (
	db *sql.DB
)

type Config struct {
	App   *app   `toml:"app"`
	MySQL *mysql `toml:"mysql"`
}

type app struct {
	Name string `toml:"name" env:"APP_NAME"`
	HTTP *http  `toml:"http"`
	GRPC *grpc  `toml:"grpc"`
}

type http struct {
	Host string `toml:"host" env:"HTTP_HOST"`
	Port string `toml:"port" env:"HTTP_PORT"`
}

type grpc struct {
	Host string `toml:"host" env:"GRPC_HOST"`
	Port string `toml:"port" env:"GRPC_PORT"`
}

type mysql struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	lock     sync.Mutex
}

func (h *http) GetAddr() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

func (g *grpc) GetAddr() string {
	return fmt.Sprintf("%s:%s", g.Host, g.Port)
}

func (m *mysql) getDBConn() (*sql.DB, error) {
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dbConnStr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func (m *mysql) GetDB() (*sql.DB, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		dbObj, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = dbObj
	}
	return db, nil
}

func newDefaultHttp() *http {
	return &http{
		Host: "127.0.0.1",
		Port: "8050",
	}
}

func newDefaultGrpc() *grpc {
	return &grpc{
		Host: "127.0.0.1",
		Port: "18050",
	}
}

func newDefaultMysql() *mysql {
	return &mysql{
		Host:     "127.0.0.1",
		Port:     "3306",
		UserName: "root",
		Password: "123456",
		Database: "go-micro",
	}
}

func newDefaultApp() *app {
	return &app{
		Name: "go-micro",
		HTTP: newDefaultHttp(),
		GRPC: newDefaultGrpc(),
	}
}

func newConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		MySQL: newDefaultMysql(),
	}
}
