// dbc: database config
package dbc

import (
	"fmt"
)

type Config struct {
	UserName string
	Password string
	DataBase string
	Host     string
	Port     string
	Loc      string
}

func (cs Config) MySqlConnString() string {
	if len(cs.Loc) == 0 {
		cs.Loc = "Local"
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		cs.UserName,
		cs.Password,
		cs.Host,
		cs.Port,
		cs.DataBase,
		cs.Loc,
	)
}

type Setter func(cs Config) Config

func UserName(n string) Setter {
	return func(d Config) Config {
		d.UserName = n
		return d
	}
}

func Password(p string) Setter {
	return func(d Config) Config {
		d.Password = p
		return d
	}
}

func DataBase(n string) Setter {
	return func(d Config) Config {
		d.DataBase = n
		return d
	}
}

func Host(h string) Setter {
	return func(d Config) Config {
		d.Host = h
		return d
	}
}

func Port(p string) Setter {
	return func(d Config) Config {
		d.Port = p
		return d
	}
}

func New(setters ...Setter) Config {
	d := Config{}
	for _, setter := range setters {
		d = setter(d)
	}
	return d
}
