package main

import (
	"os"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type Config struct {
	ClusterResource []string `yaml:"clusterResources" json:"clusterResources"`
}

func main() {
	var config = new(Config)

	fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		clusterResourceSet := make(map[string]struct{}, len(config.ClusterResource))
		for _, s := range config.ClusterResource {
			clusterResourceSet[s] = struct{}{}
		}

		for i := range items {
			meta, err := items[i].GetMeta()

			if err != nil {
				return nil, err
			}

			_, ok := clusterResourceSet[meta.Kind]

			if ok {
				err = items[i].SetNamespace("")

				if err != nil {
					return nil, err
				}
			}
		}
		return items, nil
	}
	p := framework.SimpleProcessor{Filter: kio.FilterFunc(fn), Config: &config}
	cmd := command.Build(p, command.StandaloneDisabled, false)

	command.AddGenerateDockerfile(cmd)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
