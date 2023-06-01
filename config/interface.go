package config

type config interface {
	ToInt() int
	ToString() string
	ToBool() bool
	envRequire(key string) *Cfg
}

type Cfg struct {
	key string
}
