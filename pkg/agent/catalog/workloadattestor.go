package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor"
	"github.com/spiffe/spire/pkg/common/catalog"
)

type workloadAttestorRepository struct {
	workloadattestor.Repository
}

func (repo *workloadAttestorRepository) Binder() any {
	return repo.AddWorkloadAttestor
}

func (repo *workloadAttestorRepository) Constraints() catalog.Constraints {
	return catalog.AtLeastOne()
}

func (repo *workloadAttestorRepository) Versions() []catalog.Version {
	return []catalog.Version{workloadAttestorV1{}}
}

type workloadAttestorV1 struct{}

func (workloadAttestorV1) New() catalog.Facade        { return new(workloadattestor.V1) }
func (workloadAttestorV1) Deprecated() (bool, string) { return false, "" }
