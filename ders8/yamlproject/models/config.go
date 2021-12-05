package models

type Config struct {
	User     string   `yaml:"User"`
	Database string   `yaml:"Database"`
	Port     string   `yaml:"Port"`
	Server   string   `yaml:"Server"`
	Settings []string `yaml:"Settings"`
}
