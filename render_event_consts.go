package main

import "strings"

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)

	for _, v := range renderer.Model.ComponentsSpecificationExtension {
		if v.Value.Yaml != "" {

			constants := strings.Split(v.Value.Yaml, "\n")
			f.WriteLine(`const (`)
			for i := 0; i < (len(constants) - 1); i++ {

				f.WriteLine(constants[i][2:] + `FactStreamType` + `=` + `"` + constants[i][2:] + `_facts"`)

			}
			f.WriteLine(`)`)
		}

	}
	return f.Bytes(), nil
}
