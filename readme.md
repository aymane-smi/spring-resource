# Spring Resources
```auto generation, happy developer```

## ⚠️ this project is under development do not use it in production solution ⚠️
## Introduction

Spring Resource is a cli tool introduced to help spring developers to generate boilerplate for their resources (repository, entity, service and controller) without any headache within a second.

## Examples
 - add docker file to your project using the following flag: `-d` or `--docker`
 - the show all available command you can use the flag: `-h` or `--help`

## Project structure
- `cmd`: folder that contains the logic behind the cli select ui
- `services`: contain all files that are responsible about creating `docker`, `dto`, `entity`, `repository` and `service`
- `static`: contain all template file(.tmpl) for docker, java and kotlin files
- `structs`: folder that contain the structure definition like `docker`, `entity`, `pom` and `shard`
- `utils`: contain all files for helper situations.

### Supported Docker images for java

- Openjdk
     - 23-jdk-oracle
     - 17-jdk-oracle 
- Eclipse Temurin
     - 17-jdk-focal
     - 8-jdk-focal

## Contributors

@Creator: aymane-smi

others:

![contributors](https://contrib.rocks/image?repo=aymane-smi/spring-resource)
