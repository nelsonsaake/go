package envs

import "os"

type Mode string

const (
	ModeProduction  Mode = "production"
	ModeDevelopment Mode = "development"
	ModeTest        Mode = "test"
	ModeLocal       Mode = "local"
)

func GetMode() string {
	return os.Getenv("GO_ENV")
}

func IsProduction() bool {
	return GetMode() == string(ModeProduction)
}

func IsLocal() bool {
	return GetMode() == string(ModeLocal)
}

func IsDevelopment() bool {
	return GetMode() == string(ModeDevelopment)
}

func IsTest() bool {
	return GetMode() == string(ModeTest)
}
