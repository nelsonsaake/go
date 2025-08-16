package cfgs

var defaultCfg *Config

func init() {
	Load("src/configs")
}

func Load(dirs ...string) {
	defaultCfg = New(dirs...)
}

type cfgsSetup struct {
	dirs []string
}

func (c cfgsSetup) Setup() {
	Load(c.dirs...)
}

func Setup(dirs ...string) cfgsSetup {
	return cfgsSetup{dirs: dirs}
}

var (
	GetBool        = defaultCfg.GetBool
	GetFloat64     = defaultCfg.GetFloat64
	GetInt         = defaultCfg.GetInt
	GetString      = defaultCfg.GetString
	GetStringMap   = defaultCfg.GetStringMap
	GetMap         = defaultCfg.GetMap
	GetSlice       = defaultCfg.GetSlice
	GetStringSlice = defaultCfg.GetStringSlice
	GetObj         = defaultCfg.GetObj
	GetAs          = defaultCfg.GetAs
	Get            = defaultCfg.Get

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
