package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/envs"
)

func mainv1() {

	// vfp = version file path
	vfp, err := afs.Path("VERSION.md")
	die(err)

	// read version from file
	bs, err := os.ReadFile(vfp)
	die(err)

	// v = version
	v := fmt.Sprintf("%s", bs)

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

	// convert back to bytes
	bs = []byte(v)

	// read file info
	info, err := os.Stat(vfp)
	die(err)

	// get file perm
	perm := info.Mode().Perm()

	// write version back to file
	err = os.WriteFile(vfp, bs, perm)
	die(err)

	// efp = env file path
	efp, err := afs.Path(".env")
	die(err)

	// set version in .env
	err = envs.Set(efp, "VERSION", v)
	die(err)
}
