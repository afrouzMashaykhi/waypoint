package config

import (
	"github.com/mitchellh/mapstructure"
)

// Operation is something in the Waypoint configuration that is executed
// using some underlying plugin. This is a general shared structure that is
// used by internal/core to initialize all the proper plugins.
type Operation struct {
	Labels map[string]string `hcl:"labels,optional"`
	Hooks  []*Hook           `hcl:"hook,block"`
	Use    *Use              `hcl:"use,block"`

	// set internally to note an operation is required for validation
	required bool
}

func (b *Build) Operation() *Operation {
	return mapoperation(b, true)
}

func (b *Registry) Operation() *Operation {
	return mapoperation(b, false)
}

func (b *Deploy) Operation() *Operation {
	return mapoperation(b, true)
}

func (b *Release) Operation() *Operation {
	return mapoperation(b, false)
}

// mapoperation takes a struct that is a superset of Operation and
// maps it down to an Operation. This will panic if this fails.
func mapoperation(input interface{}, req bool) *Operation {
	if input == nil {
		return nil
	}

	var op Operation
	if err := mapstructure.Decode(input, &op); err != nil {
		panic(err)
	}
	op.required = req

	return &op
}
