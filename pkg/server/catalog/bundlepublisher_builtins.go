//go:build !spire_minimal

package catalog

import (
	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher/awsrolesanywhere"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher/awss3"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher/azureblob"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher/gcpcloudstorage"
	"github.com/spiffe/spire/pkg/server/plugin/bundlepublisher/k8sconfigmap"
)

func (repo *bundlePublisherRepository) BuiltIns() []catalog.BuiltIn {
	return []catalog.BuiltIn{
		awss3.BuiltIn(),
		azureblob.BuiltIn(),
		gcpcloudstorage.BuiltIn(),
		awsrolesanywhere.BuiltIn(),
		k8sconfigmap.BuiltIn(),
	}
}
