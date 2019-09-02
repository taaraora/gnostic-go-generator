package main

import (
	"strings"
)

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	var constants []string
	for _, namedAny := range renderer.Model.ComponentsSpecificationExtension {
		constants = append(constants, namedAny.Value.Yaml)
	}
	f.WriteLine(`const (`)
	//generate constants FactStreamType
	if len(constants) > 0 {
		for _, constant := range splitLineConstants(constants[0]) {
			f.WriteLine(constant + `FactStreamType` + `=` + `"` + constant + `_facts"`)
		}
	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	//generate constants CommandsStreamType
	if len(constants) >= 1 {
		for _, constant := range splitLineConstants(constants[1]) {
			f.WriteLine(constant + `CommandsStreamType` + `=` + `"` + constant + `_commands"`)
		}
	}
	f.WriteLine(`)`)
	return f.Bytes(), nil
}
func splitLineConstants(constant string) []string {
	constants := strings.Split(constant, "\n")
	constants = constants[:len(constants)-1]
	for i := 0; i < len(constants); i++ {
		constants[i] = constants[i][2:]
	}
	return constants
}
