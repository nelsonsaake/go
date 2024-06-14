package do

import "github.com/joho/godotenv"

// LoadEnv: load .env and parse the env in working dir
func LoadEnv() error {
	return godotenv.Load(".env")
}
