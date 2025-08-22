// Copyright 2025 The OpenCmd Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package schema

type Schema struct {
	Version string `json:"version" yaml:"version"`
	Vendor  string `json:"vendor" yaml:"vendor"`
	Name    string `json:"name" yaml:"name"`

	Description []string            `json:"description,omitempty" yaml:"description,omitempty"`
	Options     map[string]*Option  `json:"options,omitempty" yaml:"options,omitempty"`
	SubCommands map[string]*Command `json:"commands,omitempty" yaml:"commands,omitempty"`
	Constraints []*Constraint       `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Extends     string              `json:"extends,omitempty" yaml:"extends,omitempty"`

	Definitions map[string]*TypeDefinition `json:"definitions,omitempty" yaml:"definitions,omitempty"`
}

type TypeDefinition struct {
	Type        string         `json:"type" yaml:"type"`
	Description []string       `json:"description,omitempty" yaml:"description,omitempty"`
	Properties  map[string]any `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type CommandDefinition struct {
	Name        string              `json:"name" yaml:"name"`
	Description []string            `json:"description,omitempty" yaml:"description,omitempty"`
	Options     map[string]*Option  `json:"options,omitempty" yaml:"options,omitempty"`
	SubCommands map[string]*Command `json:"commands,omitempty" yaml:"commands,omitempty"`
	Constraints []*Constraint       `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Extends     string              `json:"extends,omitempty" yaml:"extends,omitempty"`
}

type Option struct {
	Type        string       `json:"type" yaml:"type"`
	Required    bool         `json:"required,omitempty" yaml:"required,omitempty"`
	Default     any          `json:"default,omitempty" yaml:"default,omitempty"`
	Alias       []string     `json:"alias,omitempty" yaml:"alias,omitempty"`
	Description []string     `json:"description,omitempty" yaml:"description,omitempty"`
	Decorators  []*Decorator `json:"decorators,omitempty" yaml:"decorators,omitempty"`
}

type Command struct {
	Description []string            `json:"description,omitempty" yaml:"description,omitempty"`
	Options     map[string]*Option  `json:"options,omitempty" yaml:"options,omitempty"`
	SubCommands map[string]*Command `json:"commands,omitempty" yaml:"commands,omitempty"`
	Constraints []*Constraint       `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Extends     string              `json:"extends,omitempty" yaml:"extends,omitempty"`
}

type Decorator struct {
	Name   string         `json:"name" yaml:"name"`
	Params map[string]any `json:"params,omitempty" yaml:"params,omitempty"`
}

type Constraint struct {
	Type         string   `json:"type" yaml:"type"`
	Expression   string   `json:"expression" yaml:"expression"`
	Message      string   `json:"message,omitempty" yaml:"message,omitempty"`
	LeftOperand  *Operand `json:"left_operand" yaml:"left_operand"`
	RightOperand *Operand `json:"right_operand,omitempty" yaml:"right_operand,omitempty"`
	Operator     string   `json:"operator,omitempty" yaml:"operator,omitempty"`
}

type Operand struct {
	Type   string  `json:"type" yaml:"type"`
	Path   string  `json:"path,omitempty" yaml:"path,omitempty"`
	Value  any     `json:"value,omitempty" yaml:"value,omitempty"`
	Filter *Filter `json:"filter,omitempty" yaml:"filter,omitempty"`
}

type Filter struct {
	Type  string         `json:"type" yaml:"type"`
	Value map[string]any `json:"value" yaml:"value"`
}
