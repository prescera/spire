package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/notifier"
)

type notifierRepository struct {
	notifier.Repository
}

func (repo *notifierRepository) Binder() any {
	return repo.AddNotifier
}

func (repo *notifierRepository) Constraints() catalog.Constraints {
	return catalog.ZeroOrMore()
}

func (repo *notifierRepository) Versions() []catalog.Version {
	return []catalog.Version{
		notifierV1{},
	}
}

type notifierV1 struct{}

func (notifierV1) New() catalog.Facade { return new(notifier.V1) }
func (notifierV1) Deprecated() (bool, string) {
	return true, "use the BundlePublisher plugin type instead"
}
