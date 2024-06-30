package main

import (
	"flag"
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
	flagI := flag.String("i", "CONST", "default is constructor injection")
	flagInject := flag.String("inject", "CONST", "default is constructor injection")

	if len(osArgs) > 1 {
		flag.Parse()
		handleArgs(osArgs[1], flagI, flagInject, shared, &path)
	} else {
		executeMainFunctionality(shared, path, templType)
	}
}

func handleArgs(arg string, flagI, flagInject *string, shared *structs.Shared, path *string) {
	switch arg {
	case "-d", "--docker":
		handleDockerGeneration(path)
	case "-h", "--help":
		displayHelp()
	default:
		if utils.CheckInjectValue(*flagI) || utils.CheckInjectValue(*flagInject) {
			shared.Injection = utils.OrStr(*flagI, *flagInject)
			executeMainFunctionality(shared, *path, 1)
		}
	}
}

func handleDockerGeneration(path *string) {
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
	fmt.Scanf("%s", path)
	service.GenerateDocker(*path, *docker)
}

func displayHelp() {
	fmt.Println(`
Usage: spring-resource [FLAG]
	
Available flags in spring resource CLI:

Options:

-d, --docker generate Dockerfile for your Spring project
-h, --help display list of commands available
-i, --injection choose the DI method, ex: -i=[AUTO]
`)
}

func executeMainFunctionality(shared *structs.Shared, path string, templType int) {
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
		fmt.Println("the path is not for a Java or Kotlin project")
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
	service.GenerateService(*shared, path, templType)
	service.GenerateServiceImpl(*shared, path, templType)
	fmt.Println("\nAll files generated ✅ please check the folder =>", path)
}
