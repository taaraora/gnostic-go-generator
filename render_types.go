// Copyright 2019 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	surface "github.com/googleapis/gnostic/surface"
)

func (renderer *Renderer) RenderTypes() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	f.WriteLine(`// Types used by the API.`)
	for _, modelType := range renderer.Model.Types {
		f.WriteLine(`// ` + modelType.Description)
		if modelType.Kind == surface.TypeKind_STRUCT {
			f.WriteLine(`type ` + modelType.TypeName + ` struct {`)
			for _, field := range modelType.Fields {
				typ := field.NativeType
				if field.Kind == surface.FieldKind_REFERENCE {
					typ = "*" + typ
				} else if field.Kind == surface.FieldKind_ARRAY {
					typ = "[]" + typ
				} else if field.Kind == surface.FieldKind_MAP {
					typ = "map[string]" + typ
				} else if field.Kind == surface.FieldKind_ANY {
					typ = "interface{}"
				}
				f.WriteLine(field.FieldName + ` ` + typ + structTagGenerate(jsonTag(field), dbTag(field)))
			}
			f.WriteLine(`}`)
		} else if modelType.Kind == surface.TypeKind_OBJECT {
			f.WriteLine(`type ` + modelType.TypeName + ` map[string]` + modelType.ContentType)
		} else {
			f.WriteLine(`type ` + modelType.TypeName + ` interface {}`)
		}
	}
	return f.Bytes(), nil
}

func structTagGenerate(tags ...string) string {
	var name string
	for k,v := range tags{
		if tags[0] != "" {
			if k != len(tags)-1 && len(tags) != 1 && v[len(v)-1] == 96 && tags[len(tags)-1] != "" {
				name += v[:len(v)-1]
			} else {
				name += v
			}
		}
	}
	return name
}


func jsonTag(field *surface.Field) string {
	if field.Serialize && field.Name != "" {
		if field.Name[0] <= 65 || field.Name[0] <= 90{
			field.Name = string(field.Name[0]+32) + field.Name[1:]
		}
		return "`json:" + `"` + field.Name + `,omitempty"` + "`"
	}

	return ""
}

func dbTag(field *surface.Field) string {
	var tagComponents []rune
	var numberOfWords bool
	counter := 0
	for k, v := range field.Name {
		var lower = false
		if v <= 65 || v <= 90 {
			numberOfWords = true
			lower = true
			counter++
			if counter < 2 && k !=0 {
				tagComponents = append(tagComponents, 95)
			}
		}
		if lower {
			tagComponents = append(tagComponents, v+32)
		} else {
			counter = 0
			tagComponents = append(tagComponents, v)
		}
	}
	if !numberOfWords {
		return ""
	}
	return ` db:` + `"` + string(tagComponents) + `"`+"`"
}
