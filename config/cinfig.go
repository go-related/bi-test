package config

type Config struct {
	ConnectionString       string
	Port                   string
	FactoryOwnerPrivateKey string
}

func LoadEnvVariables() Config {
	return Config{
		ConnectionString: "host=localhost port=5432 user=blockinvest dbname=test password=Test123 sslmode=disable",
		Port:             ":8080",
	}
}
