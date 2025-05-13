package config

// Config 配置结构体
type Config struct {
	MySQL  MySQLConfig  `yaml:"mysql"`
	Server ServerConfig `yaml:"server"`
}

// MySQLConfig MySQL配置结构体
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// ServerConfig 服务器配置结构体
type ServerConfig struct {
	Port int `yaml:"port"`
}
