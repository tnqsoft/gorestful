package config

type Config struct {
	DB  *DBConfig
	APP *AppConfig
}

type AppConfig struct {
	Hostname string
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "12345678",
			Name:     "gorest",
			Charset:  "utf8",
		},
		APP: &AppConfig{
			Hostname: ":3000",
		},
	}
}
