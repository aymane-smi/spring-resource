package service

import (
	"os"
	"text/template"

	"github.com/aymane-smi/spring-resource/structs"
)

func GenerateDocker(path string, docker structs.Docker) (bool, error) {
	tmpl, _ := template.ParseFiles("static/docker.tmpl")
	file, errFile := os.Create(path + "/Dockerfile")
	if errFile != nil {
		return false, errFile
	}
	defer file.Close()
	tmpl.Execute(file, docker)
	return true, nil
}
