package main

import (
	"strings"
)

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	f.WriteLine(`// Types used by the API.`)
	f.WriteLine(`const (`)
	constantsFactStreamType := strings.Split(renderer.Model.FactsStreamTypes, "\n")
	for i := 0; i < (len(constantsFactStreamType) - 1); i++ {

		f.WriteLine(constantsFactStreamType[i][2:] + `FactStreamType` + `=` + `"` + constantsFactStreamType[i][2:] + `_facts"`)

	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	constantsCommandStreamType := strings.Split(renderer.Model.CommandStreamTypes, "\n")
	for i := 0; i < (len(constantsCommandStreamType) - 1); i++ {

		f.WriteLine(constantsCommandStreamType[i][2:] + `FactStreamType` + `=` + `"` + constantsCommandStreamType[i][2:] + `_facts"`)

	}
	f.WriteLine(`)`)

	return f.Bytes(), nil
}
