package main

import (
	"fmt"
	"os"

	"github.com/aymane-smi/spring-resource/cmd"
	service "github.com/aymane-smi/spring-resource/services"
	"github.com/aymane-smi/spring-resource/structs"
	"github.com/aymane-smi/spring-resource/utils"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	newFigure := figure.NewFigure("Spring Resource", "", true)
	newFigure.Print()
	var path string
	var templType int = 1
	shared := &structs.Shared{}
	osArgs := os.Args
	if len(osArgs) > 1 {
		if osArgs[1] == "-d" || osArgs[1] == "--docker" {
			docker := &structs.Docker{
				Image: "",
				Tag:   "",
			}
			image, _ := cmd.Execute(3, 0)
			var tagType int
			if image == "openjdk" {
				tagType = 1
			} else {
				tagType = 2
			}
			tag, _ := cmd.Execute(4, tagType)
			docker.Image = image.(string)
			docker.Tag = tag.(string)
			fmt.Printf("project path:")
			fmt.Scanf("%s", &path)
			service.GenerateDocker(path, *docker)
		}
		if osArgs[1] == "-h" || osArgs[1] == "--help" {
			fmt.Println(`
Usage: spring-resource [FLAG]
			
Available flags in spring resource cli

Options:

-d, --docker generate Dockerfile for you spring project
-h, --help display list of commands available
`)
		}
	} else {
		entityType, errEntity := cmd.Execute(1, 0)
		repositoryType, errRepository := cmd.Execute(2, 0)
		shared.SharedEntity.TypeId = entityType.(int)
		shared.SharedEntity.RepoType = repositoryType.(string)
		if errRepository != nil || errEntity != nil {
			fmt.Println("error during list generation")
		}
		fmt.Print("Entity name:")
		fmt.Scanf("%s", &shared.SharedEntity.Name)
		fmt.Printf("project path:")
		fmt.Scanf("%s", &path)
		isProject, typeOfManager := utils.IsJavaOrKotlinProject(path)
		if !isProject {
			fmt.Println("the path is not for java nor kotlin project")
			os.Exit(1)
		}
		if typeOfManager == 1 {
			shared.SharedPom = utils.GenerateProjectInfoMaven(path)
		} else {
			shared.SharedPom = utils.GenerateProjectInfoGradle(path)
			result, language := utils.GradleLanguage(path)
			if !result {
				fmt.Println("error during the detection of the project language")
				os.Exit(1)
			} else {
				templType = language
			}
		}
		service.GenerateRepository(*shared, path, templType)
		service.GenerateEntity(*shared, path, templType)
		service.GenerateService(*shared, path)
		service.GenerateServiceImpl(*shared, path)
		fmt.Println("\nall file generated✅ please check the folder =>", path)
	}
}
