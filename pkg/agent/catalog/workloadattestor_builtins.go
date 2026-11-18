//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/docker"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/k8s"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/slurm"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/systemd"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/unix"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/windows"
	"github.com/spiffe/spire/pkg/common/catalog"
)

func (repo *workloadAttestorRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		docker.BuiltIn(),
		k8s.BuiltIn(),
		slurm.BuiltIn(),
		systemd.BuiltIn(),
		unix.BuiltIn(),
		windows.BuiltIn(),
	}
}
