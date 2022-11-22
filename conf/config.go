package conf

import "sync"

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
	log      sync.Mutex
}
