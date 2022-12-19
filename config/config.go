package config

type DbConfig struct {
	Name     string `envconfig:"postgres_name" default:"postgres"`
	User     string `envconfig:"postgres_user" default:"postgres"`
	Password string `envconfig:"postgres_password" default:"password"`
	Host     string `envconfig:"postgres_host" default:"host.docker.internal"`
	Port     string `envconfig:"postgres_port" default:"5432"`
}
