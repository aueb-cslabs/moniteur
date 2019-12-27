package types

type redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	Ann_DB   int    `yaml:"Ann_DB"`
	Users_DB int    `yaml:"Users_DB"`
}
