//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/awsiid"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/azureimds"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/azuremsi"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/gcpiit"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/httpchallenge"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/jointoken"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/k8spsat"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/sshpop"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/tpmdevid"
	"github.com/spiffe/spire/pkg/server/plugin/nodeattestor/x509pop"
)

func (repo *nodeAttestorRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		awsiid.BuiltIn(),
		azuremsi.BuiltIn(),
		azureimds.BuiltIn(),
		gcpiit.BuiltIn(),
		httpchallenge.BuiltIn(),
		jointoken.BuiltIn(),
		k8spsat.BuiltIn(),
		sshpop.BuiltIn(),
		tpmdevid.BuiltIn(),
		x509pop.BuiltIn(),
	}
}
