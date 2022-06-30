/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha

import (
	"fmt"

	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugins/optional/grafana/v1alpha/scaffolds"
)

var _ plugin.EditSubcommand = &editSubcommand{}

type editSubcommand struct {
	config config.Config
}

func (p *editSubcommand) UpdateMetadata(cliMeta plugin.CLIMetadata, subcmdMeta *plugin.SubcommandMetadata) {
	subcmdMeta.Description = MetaDataDescription

	subcmdMeta.Examples = fmt.Sprintf(`  # Edit a common project with this plugin
  %[1]s edit --plugins=grafana.kubebuilder.io/v1-alpha
`, cliMeta.CommandName)
}

func (p *editSubcommand) InjectConfig(c config.Config) error {
	p.config = c
	return nil
}

func (p *editSubcommand) Scaffold(fs machinery.Filesystem) error {
	if err := InsertPluginMetaToConfig(p.config, pluginConfig{}); err != nil {
		return err
	}

	scaffolder := scaffolds.NewEditScaffolder()
	scaffolder.InjectFS(fs)
	return scaffolder.Scaffold()
}