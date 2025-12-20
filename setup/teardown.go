package setup

import (
	"github.com/nelsonsaake/go/kvs"
	"github.com/nelsonsaake/go/meta"
	"github.com/nelsonsaake/go/settings"
)

func (b *BaseBuilder) Teardown() error {

	if b.KvsEnabled {
		if err := kvs.Teardown(); err != nil {
			return err
		}
	}

	if b.SettingsEnabled {
		if err := settings.Teardown(); err != nil {
			return err
		}
	}

	if b.MetaEnabled {
		if err := meta.Teardown(); err != nil {
			return err
		}
	}

	return nil
}
