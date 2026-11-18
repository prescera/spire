//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/credentialcomposer/uniqueid"
)

func (repo *credentialComposerRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{uniqueid.BuiltIn()}
}
