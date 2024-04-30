package main

import (
	service "github.com/aymane-smi/spring-resource/services"
	"github.com/aymane-smi/spring-resource/structs"
)

func main() {
	shared := &structs.Shared{
		SharedEntity: structs.Entity{
			Name:     "User",
			TypeId:   3,
			RepoType: "JpaRepository",
		},
		SharedPom: structs.Pom{
			GroupId:    "com.example",
			ArtifactId: "demo",
		},
	}
	service.GenerateEntity(*shared)
	service.GenerateDto(*shared)
	service.GenerateRepository(*shared)
	service.GenerateService(*shared)
	service.GenerateServiceImpl(*shared)

}
