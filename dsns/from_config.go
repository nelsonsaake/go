package dsns

// FromConfig creates a DSN string from a Config struct
func FromConfig(cfg Config) string {
	return New(
		WithHost(cfg.Host),
		WithPort(cfg.Port),
		WithUsername(cfg.Username),
		WithPassword(cfg.Password),
		WithDB(cfg.DBName),
	)
}
