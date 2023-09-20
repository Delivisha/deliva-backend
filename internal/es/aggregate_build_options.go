package es

import (
	"coolmarket/internal/registry"
	"fmt"
)

type VersionSetter interface {
	setVersion(int)
}

func setVersion(version int) registry.BuildOption {
	return func(v interface{}) error {
		if agg, ok := v.(VersionSetter); ok {
			agg.setVersion(version)
			return nil
		}
		return fmt.Errorf("%T does not have method SetVersion(int)")
	}
}
