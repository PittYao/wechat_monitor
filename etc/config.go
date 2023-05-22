package etc

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	C *Config
)

type Config struct {
	Database Database `yaml:"database" validate:"required"`
	Port     string   `yaml:"port" validate:"required"`
}

type Database struct {
	Host     string `yaml:"host" validate:"required"`
	Port     uint16 `yaml:"port" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

func NewConfig() *Config {
	// 读取YAML配置文件
	data, err := os.ReadFile("etc.yaml")
	if err != nil {
		fmt.Println("Failed to read config file:", err)
		panic(err)
	}

	// 解析YAML配置
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Failed to parse YAML:", err)
		panic(err)
	}

	return &config
}
