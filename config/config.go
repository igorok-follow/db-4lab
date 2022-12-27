package config

type Config struct {
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
	Token    *Token    `yaml:"token"`
	Logger   *Logger   `yaml:"logger"`
	Gateway  *Gateway  `yaml:"gateway"`
	Redis    *Redis    `yaml:"redis"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Gateway struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Uri string `yaml:"uri"`
}

type Token struct {
	SecretKey string `yaml:"secret_key"`
	PublicKey string `yaml:"public_key"`
	Salt      string `yaml:"salt"`
}

type Logger struct {
	LogLevel string `yaml:"log_level"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}
