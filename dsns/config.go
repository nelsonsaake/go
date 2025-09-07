package dsns

// Config holds all configuration variables needed for DSN creation
type Config struct {
	DBName   string
	Username string
	Password string
	Port     string
	Host     string
}
