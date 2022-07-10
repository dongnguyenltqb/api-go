package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Person struct {
	Name  string   `yaml:"name" json:"name"`
	Job   string   `yaml:"job" json:"job"`
	Skill []string `yaml:"skills" json:"skill"`
}

var cfg *Config

type Config struct {
	DSN     string              `mapstructure:"db_dsn"`
	Person  []map[string]Person `mapstructure:"persons"`
	Port    int                 `mapstructure:"port"`
	GinMode string              `mapstructure:"gin_mode"`
}

func defaut() {
	viper.SetDefault("port", 1234)
	viper.SetDefault("gin_mode", "release")
}

func Load() {
	// set default value
	defaut()

	fmt.Println("loading configuration....")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(".") // optionally look for config in the working directory
	if os.Getenv("ENV") == "test" {
		fmt.Println("adding test config path")
		viper.AddConfigPath(os.Getenv("CONFIG_PATH"))
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("loaded configuration")
	cfg = new(Config)
	viper.Unmarshal(cfg)
}

func Get() *Config {
	return cfg
}
