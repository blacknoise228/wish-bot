package config

type Config struct {
	App        App        `yaml:"app"`
	Postgres   Postgres   `yaml:"postgres"`
	API        API        `yaml:"api"`
	Telegram   Telegram   `yaml:"token"`
	Migrations Migrations `yaml:"migrations"`
}

type App struct {
	Environment string `yaml:"environment"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	SSLmode  string `yaml:"sslMode"`
}

type API struct {
	Port string `yaml:"port"`
}

type Telegram struct {
	Token string `yaml:"token"`
	Admin string `yaml:"admin"`
}

type Migrations struct {
	Migrate bool `yaml:"migrate"`
}
