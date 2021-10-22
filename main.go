package main

import (
	"os"

	"sigs.k8s.io/kustomize/kyaml/fn/framework"
	"sigs.k8s.io/kustomize/kyaml/fn/framework/command"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	var config struct {
		Data map[string]string `yaml:"data"`
	}
	fn := func(items []*yaml.RNode) ([]*yaml.RNode, error) {
		for i := range items {
			meta, err := items[i].GetMeta()

			if err != nil {
				return nil, err
			}

			if meta.Kind == "ClusterIssuer" {
				items[i].SetNamespace("")
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
