package test

type sconf struct {
	App   appConfig   `maspstructure:"app"`
	Other otherConfig `mapstructure:"other"`
}
type appConfig struct {
	Name string `mapstructure:"name"`
}
type otherConfig struct {
	Some  string `mapstructure:"some"`
	Value string `mapstructure:"value"`
}

var conf sconf
