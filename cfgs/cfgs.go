package cfgs

import "github.com/nelsonsaake/go/objs"

var defaultCfg *Config

func Load(dirs ...string) {
	defaultCfg = New(dirs...)
}

func Get(key ...string) any {
	return defaultCfg.Get(key...)
}

func GetString(key ...string) string {
	return defaultCfg.GetString(key...)
}

func GetInt(key ...string) int {
	return defaultCfg.GetInt(key...)
}

func GetBool(key ...string) bool {
	return defaultCfg.GetBool(key...)
}

func GetFloat64(key ...string) float64 {
	return defaultCfg.GetFloat64(key...)
}

func GetMap(key ...string) map[string]any {
	return defaultCfg.GetMap(key...)
}

func GetStringMap(key ...string) map[string]string {
	return defaultCfg.GetStringMap(key...)
}

func GetSlice(key ...string) []any {
	return defaultCfg.GetSlice(key...)
}

func GetStringSlice(key ...string) []string {
	return defaultCfg.GetStringSlice(key...)
}

func GetObj(key ...string) *objs.Obj {
	return defaultCfg.GetObj(key...)
}

func GetAs(dest any, key ...string) {
	defaultCfg.GetAs(dest, key...)
}

// Aliases
var (
	GS  = GetString
	GSI = GetStringSlice
	GI  = GetInt
	GF  = GetFloat64
	GB  = GetBool
	GM  = GetMap
	GSM = GetStringMap
	GSL = GetSlice
	GO  = GetObj
	GA  = Get
)
