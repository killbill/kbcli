package killbill

import "time"

type KillbillConfig interface {
	GetUrl() string
	GetUsername() string
	GetPassword() string
	GetApiKey() string
	GetApiSecret() string
	GetTimeout() time.Duration
}

type Config struct {
	Url        string
	Username   string
	Password   string
	ApiKey     string
	ApiSecret  string
	TimeoutSec int64
}

func (k *Config) GetUrl() string {
	return k.Url
}

func (k *Config) GetUsername() string {
	return k.Username
}

func (k *Config) GetPassword() string {
	return k.Password
}

func (k *Config) GetApiKey() string {
	return k.ApiKey
}

func (k *Config) GetApiSecret() string {
	return k.ApiSecret
}

func (k *Config) GetTimeout() time.Duration {
	return time.Duration(k.TimeoutSec) * time.Second
}
