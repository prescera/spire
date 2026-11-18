package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type nodeAttestorRepository struct {
	nodeattestor.Repository
}

func (repo *nodeAttestorRepository) Binder() any {
	return repo.SetNodeAttestor
}

func (repo *nodeAttestorRepository) Constraints() catalog.Constraints {
	return catalog.ExactlyOne()
}

func (repo *nodeAttestorRepository) Versions() []catalog.Version {
	return []catalog.Version{
		nodeAttestorV1{},
	}
}

type nodeAttestorV1 struct{}

func (nodeAttestorV1) New() catalog.Facade        { return new(nodeattestor.V1) }
func (nodeAttestorV1) Deprecated() (bool, string) { return false, "" }
