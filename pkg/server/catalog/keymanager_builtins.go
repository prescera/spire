//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/awskms"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/azurekeyvault"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/disk"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/gcpkms"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/hashicorpvault"
	kmipkey "github.com/spiffe/spire/pkg/server/plugin/keymanager/kmip"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/memory"
)

func (repo *keyManagerRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		awskms.BuiltIn(),
		disk.BuiltIn(),
		gcpkms.BuiltIn(),
		azurekeyvault.BuiltIn(),
		kmipkey.BuiltIn(),
		hashicorpvault.BuiltIn(),
		memory.BuiltIn(),
	}
}
