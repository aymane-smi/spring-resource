package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aymane-smi/spring-resource/structs"
)

// check if the given path is for java project (maven/gradle)
// return boolean in case of success or failure
// also return int value(0,1,2)
// 0 for non java project
// 1 for maven
// 2 for gradle
func IsJavaOrKotlinProject(path string) (bool, int) {
	maven, errMaven := os.Stat(path + "/pom.xml")
	gradleJava, errGradleJava := os.Stat(path + "/build.gradle")
	gradleKotlin, errGradleKotlin := os.Stat(path + "/build.gradle.kts")
	if errMaven != nil && errGradleJava != nil && errGradleKotlin != nil {
		return false, 0
	} else if maven != nil {
		return true, 1
	} else if gradleJava != nil {
		return true, 2
	} else if gradleKotlin != nil {
		return true, 2
	}
	return false, 0
}

// generate Pom structer with given groupId and artifactId from the given path of Maven pom
func GenerateProjectInfoMaven(path string) structs.Pom {
	argsGroup := []string{"help:evaluate", "-Dexpression=project.groupId", "-q", "-f", path, "-DforceStdout"}
	argsArtifact := []string{"help:evaluate", "-Dexpression=project.artifactId", "-q", "-f", path, "-DforceStdout"}
	outputGroup, errGroup := exec.Command("mvn", argsGroup...).Output()
	outputArtifact, errArtifact := exec.Command("mvn", argsArtifact...).Output()
	if errGroup != nil || errArtifact != nil {
		fmt.Println(errArtifact, errGroup)
	}
	return structs.Pom{GroupId: string(outputGroup), ArtifactId: string(outputArtifact)}
}

// generate Pom structer with given groupId and artifactId from the given path of Gradle settings
func GenerateProjectInfoGradle(path string) structs.Pom {
	var gradle bytes.Buffer
	fmt.Fprintf(&gradle, "-q -p %s properties", path)
	groupCmds := []string{
		gradle.String(),
		"^group:",
		"{print $2}",
	}
	nameCmds := []string{
		gradle.String(),
		"^name:",
		"{print $2}",
	}
	outputName, errName := pipeHelper(nameCmds...)
	outputGroup, errGroup := pipeHelper(groupCmds...)
	if errName != nil || errGroup != nil {
		fmt.Println("error during gradle info extraction")
		os.Exit(1)
	}
	return structs.Pom{
		GroupId:    strings.TrimSpace(string(outputGroup)),
		ArtifactId: strings.TrimSpace(string(outputName)),
	}
}

// write a helper for multiple pipe commands (gradle case)
// assuming that len(commands) = 3
func pipeHelper(commands ...string) ([]byte, error) {
	var buffer bytes.Buffer
	var err bytes.Buffer
	//init
	gradle := exec.Command("gradle", strings.Split(commands[0], " ")...)
	grep := exec.Command("grep", commands[1])
	awk := exec.Command("awk", commands[2])
	//swap rw
	grep.Stdin, _ = gradle.StdoutPipe()
	awk.Stdin, _ = grep.StdoutPipe()
	awk.Stdout = &buffer
	awk.Stderr = &err
	//run and wait
	_ = awk.Start()
	_ = grep.Start()
	_ = gradle.Run()
	_ = awk.Wait()
	_ = grep.Wait()
	if err.Len() != 0 {
		return nil, errors.New(err.String())
	}
	return buffer.Bytes(), nil
}

// check if the gardle project use java or kotlin
// return boolean in case of success or failure
// true, 1 for java
// true,2 for kotlin
// in case of ailure false, 0

func GradleLanguage(path string) (bool, int) {
	pathJava, errorJava := os.Stat(path + "/src/main/java")
	pathKotlin, errorKotlin := os.Stat(path + "/src/main/kotlin")
	if errorJava != nil && errorKotlin != nil {
		return false, 0
	}
	if pathJava != nil {
		return true, 1
	} else if pathKotlin != nil {
		return true, 2
	}
	return false, 0
}
