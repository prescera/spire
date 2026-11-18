package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"

	"github.com/spiffe/spire/pkg/agent/plugin/keymanager"
)

type keyManagerRepository struct {
	keymanager.Repository
}

func (repo *keyManagerRepository) Binder() any {
	return repo.SetKeyManager
}

func (repo *keyManagerRepository) Constraints() catalog.Constraints {
	return catalog.ExactlyOne()
}

func (repo *keyManagerRepository) Versions() []catalog.Version {
	return []catalog.Version{keyManagerV1{}}
}

type keyManagerV1 struct{}

func (keyManagerV1) New() catalog.Facade        { return new(keymanager.V1) }
func (keyManagerV1) Deprecated() (bool, string) { return false, "" }
