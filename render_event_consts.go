package main

import (
	"errors"
	"strings"
)

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	var constants struct {
		name  []string
		value []string
	}
	for _, namedAny := range renderer.Model.ComponentsSpecificationExtension {
		constants.name = append(constants.name, namedAny.Name)
		constants.value = append(constants.value, namedAny.Value.Yaml)
	}

	//generate constants FactStreamType
	if len(constants.value) > 0 {
		constant, err := splitLineConstants(constants.value[0], constants.name[0])
		if err != nil {
			return f.Bytes(), err
		}
		f.WriteLine(`const (`)
		for _, v := range constant {
			f.WriteLine(v + `FactStreamType` + `=` + `"` + v + `_facts"`)
		}
	}
	f.WriteLine(`)`)

	//generate constants CommandsStreamType
	if len(constants.value) >= 1 {
		constant, err := splitLineConstants(constants.value[1], constants.name[1])
		if err != nil {
			return f.Bytes(), err
		}
		f.WriteLine(`const (`)
		for _, v := range constant {
			f.WriteLine(v + `CommandsStreamType` + `=` + `"` + v + `_commands"`)
		}
	}
	f.WriteLine(`)`)
	return f.Bytes(), nil
}
func splitLineConstants(constant string, name string) ([]string, error) {
	constants := strings.Split(constant, "\n")
	constants = constants[:len(constants)-1]
	for i := 0; i < len(constants); i++ {
		if constants[i][:2] != "- " {
			return constants, errors.New(`Specification Extensions : ` + name + `
the specification can be expanded, for the generation of constants, only a slice of strings.

example:
         x-constants:
           - constant_1
           - constant_2  
           ...`)
		}
		constants[i] = constants[i][2:]
	}
	return constants, nil
}
