package main

func (renderer *Renderer) RenderEvent_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	f.WriteLine(`// Types used by the API.`)
	f.WriteLine(`const (`)
	for _, modelType := range renderer.Model.Types {
		if modelType.TypeSource != "" {
			f.WriteLine(modelType.TypeSource + `FactStreamType` + `=` + `"` + modelType.TypeSource + `_facts"`)
		}
	}
	f.WriteLine(`)`)
	f.WriteLine(`const (`)
	for _, modelType := range renderer.Model.Types {
		for _, x_generate := range modelType.SpecificationExtension {
			if x_generate.Name == "x-generate-command" {
				f.WriteLine(modelType.Name + `CommandsStreamType = "` + modelType.Name + `_commands"`)
			}
		}
	}
	f.WriteLine(`)`)

	return f.Bytes(), nil
}
