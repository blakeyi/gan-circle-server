package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mongo struct{
		Url string
		Db  string
		Collections struct{
			User string
		}
	}
}
