package utils

import (
	"os"
	"strings"
)

func GenerateTree(path string) bool {
	//prototype function
	//works only on array of two ["path", "subfolder1"]
	slices := strings.Split(path, "/")
	if err := os.MkdirAll(slices[0], 0755); err != nil {
		return false
	} else if err := os.MkdirAll(slices[0]+"/"+slices[1], 0755); err != nil {
		return false
	}
	os.Chmod(slices[0]+"/"+slices[1], 0755|os.ModeSetuid)
	return true
}
