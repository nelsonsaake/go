package dsns

type dsnSetter func(*dsn) *dsn

func WithHost(host string) dsnSetter {
	return func(d *dsn) *dsn {
		d.host = host
		return d
	}
}

func WithPort(port string) dsnSetter {
	return func(d *dsn) *dsn {
		d.port = port
		return d
	}
}

func WithUsername(username string) dsnSetter {
	return func(d *dsn) *dsn {
		d.username = username
		return d
	}
}

func WithPassword(password string) dsnSetter {
	return func(d *dsn) *dsn {
		d.password = password
		return d
	}
}

func WithDB(dbname string) dsnSetter {
	return func(d *dsn) *dsn {
		d.dbname = dbname
		return d
	}
}
