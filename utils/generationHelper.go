package utils

import (
	"os"
)

func GenerateTree(path string) bool {
	//prototype function
	//works only on array of two ["path", "subfolder1"]
	// slices := strings.Split(path, "/")
	// var treeBuilder string
	// for i:=0;i<len(slices);i++{
	// 	treeBuilder += slices[i]

	// }
	// if err := os.MkdirAll(slices[0], 0755); err != nil {
	// 	return false
	// } else if err := os.MkdirAll(slices[0]+"/"+slices[1], 0755); err != nil {
	// 	return false
	// }
	if err := os.MkdirAll(path, 0755); err != nil {
		return false
	} else {
		os.Chmod(path, 0755|os.ModeSetuid)
		return true
	}

}
