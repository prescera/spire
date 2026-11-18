//go:build spire_minimal

// Minimal catalog for a SPIRE agent in a confidential guest: the memory key
// manager and the unix workload attestor, nothing else built in. Node
// attestation is external (the attested cnap plugin) and is the only way in.

package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager/memory"
	"github.com/spiffe/spire/pkg/common/catalog"
)

func (repo *keyManagerRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		memory.BuiltIn(),
	}
}
