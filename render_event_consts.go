/*
	The specification can be expanded, for the generation of constants, only a slice of strings.

   		example:
   		x-constants:
   		- constant_1
   		- constant_2
		...
*/

package main

import (
	"errors"
	"fmt"
	"strings"
)

const xFactsStreamTypes = "x-facts-stream-types"
const xCommandsStreamTypes = "x-command-stream-types"

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)

	modelSpecExtensions := make(map[string]string)

	for _, namedAny := range renderer.Model.ComponentsSpecificationExtension {
		if namedAny.Name == "" {
			continue
		}
		if namedAny.Value != nil {
			modelSpecExtensions[namedAny.Name] = namedAny.Value.Yaml
		} else {
			modelSpecExtensions[namedAny.Name] = ""
		}
	}

	//generate modelSpecExtensions FactStreamType

	if factsTypes, exists := modelSpecExtensions[xFactsStreamTypes]; exists {
		blockOfConstants, err := generateStreamTypeConstants(factsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range blockOfConstants {
			f.WriteLine(v)
		}
	}

	//generate modelSpecExtensions CommandsStreamType

	if commandsTypes, exists := modelSpecExtensions[xCommandsStreamTypes]; exists {
		blockOfConstants, err := generateStreamTypeConstants(commandsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range blockOfConstants {
			f.WriteLine(v)
		}
	}
	return f.Bytes(), nil
}
func generateStreamTypeConstants(yamlValue string) ([]string, error) {
	constants, err := splitLineConstants(yamlValue)
	if err != nil {
		return nil, err
	}
	var blockOfConstants []string
	blockOfConstants = append(blockOfConstants, `const (`)
	for _, constant := range constants {
		camelCaseConstant := snakeCaseToCamelCaseWithCapitalizedFirstLetter(constant)
		blockOfConstants = append(blockOfConstants, fmt.Sprintf(`%vStreamType = "%v"`, camelCaseConstant, constant))
	}
	blockOfConstants = append(blockOfConstants, `)`)
	return blockOfConstants, err
}

func splitLineConstants(yaml string) ([]string, error) {
	result := make([]string, 0)
	const yamlArrayItemPrefix = "- "
	lines := strings.Split(yaml, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, yamlArrayItemPrefix) {
			return nil, errors.New("can't parse yaml list")
		}
		result = append(result, strings.TrimPrefix(line, yamlArrayItemPrefix))
	}
	return result, nil
}
