package main

import (
	service "github.com/aymane-smi/spring-resource/services"
	"github.com/aymane-smi/spring-resource/structs"
)

func main() {
	shared := &structs.Shared{
		SharedEntity: structs.Entity{
			Name:     "User",
			TypeId:   2,
			RepoType: "JpaRepository",
		},
		SharedPom: structs.Pom{
			GroupId:    "com.example",
			ArtifactId: "demo",
		},
		SharedException: structs.Exception{
			Name:           "ResourceNotFound",
			ResponseStatus: "NOT_FOUND",
		},
	}
	service.GenerateEntity(*shared)
	service.GenerateDto(*shared)
	service.GenerateRepository(*shared)
	service.GenerateException(*shared)
	service.GenerateService(*shared)
	service.GenerateServiceImpl(*shared)

}
