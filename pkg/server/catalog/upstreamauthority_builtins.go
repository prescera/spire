//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/awspca"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/awssecret"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/certmanager"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/disk"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/ejbca"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/gcpcas"
	spireplugin "github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/spire"
	"github.com/spiffe/spire/pkg/server/plugin/upstreamauthority/vault"
)

func (repo *upstreamAuthorityRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		awssecret.BuiltIn(),
		awspca.BuiltIn(),
		gcpcas.BuiltIn(),
		vault.BuiltIn(),
		spireplugin.BuiltIn(),
		disk.BuiltIn(),
		certmanager.BuiltIn(),
		ejbca.BuiltIn(),
	}
}
