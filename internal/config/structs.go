package config

type Config struct {
	Version  int      `json:"version"`
	Env      string   `json:"env"`
	Network  Network  `json:"network"`
	Postgres Postgres `json:"postgres"`
}

type Network struct {
	Address      string `json:"address"`
	Port         string `json:"port"`
	WriteTimeout int    `json:"writeTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
}

type Postgres struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	Sslmode      string `json:"sslmode"`
}
