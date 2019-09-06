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

const (
	xFactsStreamTypes    = "x-facts-stream-types"
	xCommandsStreamTypes = "x-command-stream-types"
	xCommandType         = "x-command-type"
	xFactType            = "x-fact-type"
)
const (
	typeConstantStreamType  = "StreamType"
	typeConstantFactType    = "FactType"
	typeConstantCommandType = "CommandType"
)

func (renderer *Renderer) Render_event_consts() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)

	modelSpecExtensions := make(map[string]string)
	for _, nah := range renderer.Model.Methods {
		for _, namedAny := range nah.SpecificationExtension {
			if namedAny.Name == "" {
				continue
			}
			if namedAny.Value != nil {
				modelSpecExtensions[namedAny.Name] = namedAny.Value.Yaml
			} else {
				modelSpecExtensions[namedAny.Name] = ""
			}
		}
	}

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
		lineConstant, err := splitLineConstants(factsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range generateTypeConstants(typeConstantStreamType, lineConstant) {
			f.WriteLine(v)
		}
	}

	//generate modelSpecExtensions CommandsStreamType

	if commandsTypes, exists := modelSpecExtensions[xCommandsStreamTypes]; exists {
		lineConstant, err := splitLineConstants(commandsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range generateTypeConstants(typeConstantStreamType, lineConstant) {
			f.WriteLine(v)
		}
	}

	//generate modelSpecExtensions FactType

	if commandsTypes, exists := modelSpecExtensions[xFactType]; exists {
		lineConstant, err := splitLineConstants(commandsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range generateTypeConstants(typeConstantFactType, lineConstant) {
			f.WriteLine(v)
		}
	}

	//generate modelSpecExtensions CommandsType

	if commandsTypes, exists := modelSpecExtensions[xCommandType]; exists {
		lineConstant, err := splitLineConstants(commandsTypes)
		if err != nil {
			return nil, err
		}
		for _, v := range generateTypeConstants(typeConstantCommandType, lineConstant) {
			f.WriteLine(v)
		}
	}
	return f.Bytes(), nil
}

func generateTypeConstants(constType string, constants []string) []string {
	var blockOfConstants []string
	blockOfConstants = append(blockOfConstants, `const (`)
	for _, constant := range constants {
		camelCaseConstant := snakeCaseToCamelCaseWithCapitalizedFirstLetter(constant)
		blockOfConstants = append(blockOfConstants, fmt.Sprintf(`%v%v = "%v"`, camelCaseConstant, constType, constant))
	}
	blockOfConstants = append(blockOfConstants, `)`)
	return blockOfConstants
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
