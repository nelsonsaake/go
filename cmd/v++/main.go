package main

import (
	"strconv"
	"strings"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/envs"
)

var die = do.Die

func main() {

	// efp = env file path
	efp, err := afs.Path(".env")
	die(err)

	// load envs from file
	err = envs.Load(efp)
	die(err)

	// get version from env
	v := envs.Get("VERSION")

	// vs = version slice
	vs := strings.Split(v, ".")

	// last index
	li := len(vs) - 1

	// mvn = minor version number
	mvn := vs[li]

	// mvni = minor version number as int
	mvni, _ := strconv.Atoi(mvn)

	// increment mvni
	mvni++

	// convert back to string
	mvn = strconv.Itoa(mvni)

	// update value in slice
	vs[li] = mvn

	// generate a new version
	v = strings.Join(vs, ".")

	// set version in .env
	err = envs.Set(efp, "VERSION", v)
	die(err)
}
