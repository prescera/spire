//go:build spire_minimal

// Minimal catalog for a SPIRE server running in a confidential guest: only the
// built-in plugins such a deployment uses. Everything else (cloud attestors,
// key managers, publishers) is left out of the binary -- smaller, so it fits
// an in-guest image unpack, and with less code inside the trust domain. In
// particular no built-in node attestor is compiled in, so external, attested
// node attestation is the only way an agent can be admitted.

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
)

func (repo *credentialComposerRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{}
}
