package dsc

import "github.com/nelsonsaake/go/tr"

func Tr(d tr.Subject, id, t string, rsrc ...string) any {
	return getDefaultDescriptorService().Tr(d, id, t, rsrc...)
}

func dotr(rsrc, fromName, fromValue, to string) string {

	res, ok := New().
		WithValue(fromName, fromValue).
		// create a unique id,
		// this will be used to cache the calculation
		// this way we don't do it twice
		WithID(rsrc, fromValue).
		// rules to use
		WithRSrc(rsrc, "rules").
		// translate
		Tr(to).(string)

	if !ok {
		return "<nil>"
	}

	return res
}

func TrDBTypeToGoType(dbType string) string {

	return dotr(
		"tr.db_type.go_type",
		"dbType", dbType,
		"goType",
	)
}

func TrDBTypeToPhpType(dbType string) string {

	return dotr(
		"tr.db_type.php_type",
		"dbType", dbType,
		"phpType",
	)
}

func TrGoTypeToLaravelTypeValidation(goType string) string {

	return dotr(
		"tr.go_type.lar_vld",
		"goType", goType,
		"vld",
	)
}
