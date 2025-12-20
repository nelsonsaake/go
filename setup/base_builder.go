package setup

import (
	"github.com/nelsonsaake/go/cfgs"
	"github.com/nelsonsaake/go/crypt"
	"github.com/nelsonsaake/go/envs"
	"github.com/nelsonsaake/go/kernel"
	"github.com/nelsonsaake/go/kvs"
	"github.com/nelsonsaake/go/logs"
	"github.com/nelsonsaake/go/meta"
	"github.com/nelsonsaake/go/settings"
)

type Builder interface {
	Run()
}

type BaseBuilder struct {
	EnvPath             []string
	EnvContent          string
	ConfigEnabled       bool
	ConfigPath          string
	KernelEnabled       bool
	KernelName          string
	CryptEnabled        bool
	CryptSecret         string
	SettingsEnabled     bool
	SettingsFilePath    string
	MetaEnabled         bool
	MetaFilePath        string
	KvsEnabled          bool
	KvsFilePath         string
	SpatieEnabled       bool
	PassportEnabled     bool
	SeedersEnabled      bool
	LogsEnabled         bool
	CliRoutesEnabled    bool
	StaticRoutesEnabled bool
	APIRoutesEnabled    bool
	WebRoutesEnabled    bool
	Builder             Builder
}

func NewBaseBuilder(builder Builder) *BaseBuilder {
	return &BaseBuilder{
		EnvPath:             []string{DefaultEnvPath},
		ConfigEnabled:       DefaultConfigEnabled,
		ConfigPath:          DefaultConfigPath,
		KernelEnabled:       DefaultKernelEnabled,
		KernelName:          DefaultKernelName,
		CryptEnabled:        DefaultCryptEnabled,
		CryptSecret:         DefaultCryptSecret,
		SettingsEnabled:     DefaultSettingsEnabled,
		SettingsFilePath:    DefaultSettingsPath,
		MetaEnabled:         DefaultMetaEnabled,
		MetaFilePath:        DefaultMetaPath,
		KvsEnabled:          DefaultKvsEnabled,
		KvsFilePath:         DefaultKvsPath,
		StaticRoutesEnabled: DefaultStaticRoutesEnabled,
		APIRoutesEnabled:    DefaultAPIRoutesEnabled,
		WebRoutesEnabled:    DefaultWebRoutesEnabled,
		Builder:             builder,
	}
}

func (b *BaseBuilder) WithEnvPath(path ...string) *BaseBuilder {
	b.EnvPath = path
	return b
}

func (b *BaseBuilder) WithEnvContent(content string) *BaseBuilder {
	b.EnvContent = content
	return b
}

func (b *BaseBuilder) WithConfigPath(path string) *BaseBuilder {
	b.ConfigPath = path
	return b
}

func (b *BaseBuilder) WithKernelName(name string) *BaseBuilder {
	b.KernelName = name
	return b
}

func (b *BaseBuilder) WithCryptSecret(secret string) *BaseBuilder {
	b.CryptSecret = secret
	return b
}

func (b *BaseBuilder) WithSettingsEnabled() *BaseBuilder {
	b.SettingsEnabled = true
	return b
}

func (b *BaseBuilder) WithSettingsDisabled() *BaseBuilder {
	b.SettingsEnabled = false
	return b
}

func (b *BaseBuilder) WithSettingsFilePath(path string) *BaseBuilder {
	b.SettingsFilePath = path
	return b
}

func (b *BaseBuilder) WithMetaEnabled() *BaseBuilder {
	b.MetaEnabled = true
	return b
}

func (b *BaseBuilder) WithMetaDisabled() *BaseBuilder {
	b.MetaEnabled = false
	return b
}

func (b *BaseBuilder) WithMetaFilePath(path string) *BaseBuilder {
	b.MetaFilePath = path
	return b
}

func (b *BaseBuilder) WithKvsEnabled() *BaseBuilder {
	b.KvsEnabled = true
	return b
}

func (b *BaseBuilder) WithKvsDisabled() *BaseBuilder {
	b.KvsEnabled = false
	return b
}

func (b *BaseBuilder) WithKvsFilePath(path string) *BaseBuilder {
	b.KvsFilePath = path
	return b
}

func (b *BaseBuilder) DisableCLI() *BaseBuilder {
	b.CliRoutesEnabled = false
	return b
}

func (b *BaseBuilder) DisableConfig() *BaseBuilder {
	b.ConfigEnabled = false
	return b
}

func (b *BaseBuilder) DisableKernel() *BaseBuilder {
	b.KernelEnabled = false
	return b
}

func (b *BaseBuilder) DisableSpatie() *BaseBuilder {
	b.SpatieEnabled = false
	return b
}

func (b *BaseBuilder) DisablePassport() *BaseBuilder {
	b.PassportEnabled = false
	return b
}

func (b *BaseBuilder) DisableCrypt() *BaseBuilder {
	b.CryptEnabled = false
	return b
}

func (b *BaseBuilder) DisableSeeders() *BaseBuilder {
	b.SeedersEnabled = false
	return b
}

func (b *BaseBuilder) WithLogs() *BaseBuilder {
	b.LogsEnabled = true
	return b
}

func (b *BaseBuilder) DisableLogs() *BaseBuilder {
	b.LogsEnabled = false
	return b
}

func (b *BaseBuilder) WithStaticRoutes() *BaseBuilder {
	b.StaticRoutesEnabled = true
	return b
}

func (b *BaseBuilder) WithAPIRoutes() *BaseBuilder {
	b.APIRoutesEnabled = true
	return b
}

func (b *BaseBuilder) WithWebRoutes() *BaseBuilder {
	b.WebRoutesEnabled = true
	return b
}

func (b *BaseBuilder) DisableStaticRoutes() *BaseBuilder {
	b.StaticRoutesEnabled = false
	return b
}

func (b *BaseBuilder) DisableAPIRoutes() *BaseBuilder {
	b.APIRoutesEnabled = false
	return b
}

func (b *BaseBuilder) DisableWebRoutes() *BaseBuilder {
	b.WebRoutesEnabled = false
	return b
}

func (b *BaseBuilder) Run() {

	// Setup environment
	if b.EnvContent != "" {
		envs.LoadContent(b.EnvContent)
	} else {
		envs.Setup(envs.Config{
			Paths: b.EnvPath,
		})
	}

	// Silence logs early unless enabled on the builder
	if !b.LogsEnabled {
		logs.Silent()
	}

	// Setup config if enabled
	if b.ConfigEnabled {
		cfgs.Load(b.ConfigPath)
	}

	// Setup kernel if enabled
	if b.KernelEnabled {
		kernel.Setup(kernel.Config{
			Name: b.KernelName,
		})
	}

	// Setup crypt if enabled
	if b.CryptEnabled {
		crypt.Setup(crypt.Config{
			Secret: b.CryptSecret,
		})
	}

	// Setup settings if enabled
	if b.SettingsEnabled {
		settings.Setup(settings.Config{
			FilePath: b.SettingsFilePath,
		})
	}

	// Setup meta if enabled
	if b.MetaEnabled {
		meta.Setup(meta.Config{
			FilePath: b.MetaFilePath,
		})
	}

	// Setup kvs if enabled
	if b.KvsEnabled {
		kvs.Setup(kvs.Config{
			FilePath: b.KvsFilePath,
		})
	}

	b.Builder.Run()
}
