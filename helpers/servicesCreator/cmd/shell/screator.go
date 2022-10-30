package main

import (
	"flag"

	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/domain"
	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/services/folders"
	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/services/pipeline"
)

var (
	servicePath     = flag.String("path", "", "Relative path from where the program is being called")
	serviceName     = flag.String("name", "", "Service name, only one prhase")
	githubFlowsPath = flag.String("pipeline", "", "Github flow path")
)

func main() {
	flag.Parse()

	serviceInfo := domain.NewServiceInformation(*servicePath, domain.ServiceName(*serviceName))
	pipeline.CreatePipeline(serviceInfo, githubFlowsPath)
	folders.CreateFolderStructure(serviceInfo)
}
