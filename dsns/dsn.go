package dsns

import "fmt"

type dsn struct {
	dbname   string
	username string
	password string
	port     string
	host     string
}

// String: format dsn into string like so
// "root:@tcp(127.0.0.1:3306)/rb_stthomas?charset=utf8mb4&parseTime=True&loc=Local"
func (d *dsn) String() string {

	// specify the format of the DSN string
	dnsf := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	// bind the values to the format
	dnsString := fmt.Sprintf(dnsf, d.username, d.password, d.host, d.port, d.dbname)

	return dnsString
}

// New: creates a new DSN string with the given username, password, host, port, and dbname
// default values are used if not provided
// default values: username="root", password="", host="127.0.0.1", port="3306"
func New(setters ...dsnSetter) string {

	// prototype very common DSN
	d := &dsn{username: "root", password: "", host: "127.0.0.1", port: "3306"}

	// apply custom setters, if any
	for _, setter := range setters {
		d = setter(d)
	}

	// return dns
	return d.String()
}
