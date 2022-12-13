package config

type Conf struct {
	App struct {
		Name       string `env:"APP_NAME"`
		Port       string `env:"APP_PORT"`
		Mode       string `env:"APP_MODE"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET"`
		Tmp        string `env:"APP_TMP"`
	}
	Db struct {
		Host string `env:"DB_HOST"`
		Name string `env:"DB_NAME"`
		User string `env:"DB_USER"`
		Pass string `env:"DB_PASSWORD"`
		Port string `env:"DB_PORT"`
	}
	BasicAuth struct {
		Username string `env:"BASIC_AUTH_USER"`
		Password string `env:"BASIC_AUTH_PASSWORD"`
	}
	Blockhain struct {
		Host        string `env:"BLOCKCHAIN_HOST"`
		Key         string `env:"BLOCKCHAIN_KEY"`
		Secret_key  string `env:"BLOCKCHAIN_SECRET"`
		Account     string `env:"BLOCKCHAIN_ACCOUNT"`
		Public      string `env:"BLOCKCHAIN_PUB"`
		Secret_base string `env:"BLOCKCHAIN_BASE"`
		//Port       string `env:"BLOCKCHAIN_PORT"`
		//Keystore   string `env:"BLOCKCHAIN_KEYSTORE"`
	}
	Contract struct {
		Smart_contract string `env:"SMART_CONTRACT"`
	}
	IPFS struct {
		Host string `env:"IPFS_HOST"`
		Port string `env:"IPFS_PORT"`
	}
	Email struct {
		Host string `env:"EMAIL_HOST"`
		Port string `env:"EMAIL_PORT"`
		User string `env:"EMAIL_USER"`
		Pass string `env:"EMAIL_PASSWORD"`
	}
	Signature struct {
		Key string `env:"SIGNATURE_KEY"`
	}
}
