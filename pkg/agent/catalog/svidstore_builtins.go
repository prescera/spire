//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/agent/plugin/svidstore/awssecretsmanager"
	"github.com/spiffe/spire/pkg/agent/plugin/svidstore/gcpsecretmanager"
	"github.com/spiffe/spire/pkg/common/catalog"
)

func (repo *svidStoreRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		awssecretsmanager.BuiltIn(),
		gcpsecretmanager.BuiltIn(),
	}
}
