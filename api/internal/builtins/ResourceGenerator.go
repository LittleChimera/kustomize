// Code generated by pluginator on ResourceGenerator; DO NOT EDIT.
// pluginator {(devel)  unknown   }

package builtins

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

type ResourceGeneratorPlugin struct {
	h        *resmap.PluginHelpers
	Resource string `json:"resource" yaml:"resource"`
}

func (p *ResourceGeneratorPlugin) Config(h *resmap.PluginHelpers, config []byte) error {
	p.h = h
	if err := yaml.Unmarshal(config, p); err != nil {
		return fmt.Errorf("failed to unmarshal ResourceGenerator config: %w", err)
	}
	return nil
}

func (p *ResourceGeneratorPlugin) Generate() (resmap.ResMap, error) {
	resmap, err := p.h.AccumulateResource(p.Resource)
	if err != nil {
		return nil, fmt.Errorf("failed to Accumulate: %w", err)
	}
	return resmap, nil
}

func NewResourceGeneratorPlugin() resmap.GeneratorPlugin {
	return &ResourceGeneratorPlugin{}
}
