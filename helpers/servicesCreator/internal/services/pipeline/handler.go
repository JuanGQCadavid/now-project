package pipeline

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/domain"
)

//go:embed resources/servicepipeline.yaml
var servicePipelineTemplate string

func CreatePipeline(serviceInfo *domain.ServiceInformation, githubFlowsPath *string) {

	pipelineFile, err := os.OpenFile(*githubFlowsPath, os.O_CREATE|os.O_WRONLY, 0600)
	defer pipelineFile.Close()

	if err != nil {
		log.Fatal("An error happen while opening github pipeline file", githubFlowsPath, "\nError:", err.Error())
	}

	a := strings.ReplaceAll(servicePipelineTemplate, "{{.ServicePath}}", serviceInfo.ServicePath)
	a = strings.ReplaceAll(a, "{{.ServiceName}}", serviceInfo.ServiceName.ToLoweCase())
	a = strings.ReplaceAll(a, "{{.ServiceLambdaNale}}", serviceInfo.ServiceName.ToCapitalCase())

	log.Println(a)
	count, err := pipelineFile.WriteString(a)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("count: ", count)
	}

	// err = serviceTemplate.Execute(pipelineFile, serviceInfo)

	// if err != nil {
	// 	log.Fatal("An error happen while performing the pipeline template parsering", "\nError:", err.Error())
	// }
}
