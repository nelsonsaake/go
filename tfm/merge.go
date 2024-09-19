// tfm: allows merging template.FuncMap
package tfm

import "html/template"

// Function to merge multiple FuncMaps
func MergeFuncMaps(funcMaps ...template.FuncMap) template.FuncMap {
	merged := template.FuncMap{}
	for _, fm := range funcMaps {
		for key, value := range fm {
			merged[key] = value
		}
	}
	return merged
}
