package config

type Config struct {
	Token string   `yaml:"token"`
	Host  string   `yaml:"host"`
	DB    Database `yaml:"db"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
