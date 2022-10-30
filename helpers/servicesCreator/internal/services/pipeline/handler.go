package pipeline

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/domain"
)

//go:embed resources/servicepipeline.yaml
var servicePipelineTemplate string

func CreatePipeline(serviceInfo *domain.ServiceInformation, githubFlowsPath *string) {

	filePathNamed := *githubFlowsPath + fmt.Sprintf("/deploy-go-%s.yml", serviceInfo.ServiceName.ToLoweCase())
	pipelineFile, err := os.OpenFile(filePathNamed, os.O_CREATE|os.O_WRONLY, 0600)
	defer pipelineFile.Close()
	if err != nil {
		log.Fatal("An error happen while opening github pipeline file", filePathNamed, "\nError:", err.Error())
	}

	a := strings.ReplaceAll(servicePipelineTemplate, "{{.ServicePath}}", fmt.Sprintf("%s/%s", serviceInfo.ServicePath, serviceInfo.ServiceName.ToLoweCase()))
	a = strings.ReplaceAll(a, "{{.ServiceName}}", serviceInfo.ServiceName.ToLoweCase())
	a = strings.ReplaceAll(a, "{{.ServiceLambdaNale}}", serviceInfo.ServiceName.ToCapitalCase())

	log.Println(a)
	count, err := pipelineFile.WriteString(a)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("count: ", count)
	}
}
