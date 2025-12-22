package utils

var config *Configs

type Configs struct {
	AppName        string
	AppEnvironment string
	Redis          DB
}
type DB struct {
	Host               string
	DbName             string
	StreamSubject      string
	StreamConsumeGroup string
}

type Api struct {
	Url    string
	APIKey string
}

func NewConfig(conf Configs) (err error) {
	config = &conf
	return
}

func GetConfig() *Configs {
	return config
}
