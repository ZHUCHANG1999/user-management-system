package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Database DatabaseConf
	JWT      JWTConf
}

type DatabaseConf struct {
	DataSource string
}

type JWTConf struct {
	Secret string
	Expire int64
}
