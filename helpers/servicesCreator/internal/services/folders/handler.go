package folders

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/domain"
	"gopkg.in/yaml.v2"
)

type FolderStructure struct {
	Root []Element `yaml:"root"`
}

type ElemetType string

const (
	FolderType ElemetType = "folder"
	FileType   ElemetType = "file"
)

type Element struct {
	Type     ElemetType `yaml:"type"`
	Name     string     `yaml:"name"`
	Inner    []*Element `yaml:"inner"`
	CopyFrom string     `yaml:"copyFrom"`
}

var (
	//go:embed resources/structure.yml
	configFile string
	//go:embed resources/build.sh
	buildsh string
	//go:embed resources/deploy.sh
	deploysh string
	//go:embed resources/.gitignore
	gitignore string
)

func CreateFolderStructure(serviceInfo *domain.ServiceInformation) {
	var folderStructure *FolderStructure = &FolderStructure{}

	err := yaml.Unmarshal([]byte(configFile), folderStructure)

	if err != nil {
		log.Fatalln("There were an erro while marshialing the structure CreateFolderStructure", "error: ", err.Error())
	}

	log.Println(fmt.Sprintf("%+v", folderStructure))

	// Initial root
	root := fmt.Sprintf("%s/%s", serviceInfo.ServicePath, serviceInfo.ServiceName.ToLoweCase())
	err = os.MkdirAll(root, 0777)

	if err != nil {
		log.Fatal("Error while creating folder", root, "err:", err.Error())
	}

	for _, element := range folderStructure.Root {
		doFolderCreation(element, root)
	}

}

func doFolderCreation(element Element, root string) {
	log.Println(element.Type, element.Name, "At:", root)

	newPath := fmt.Sprintf("%s/%s", root, element.Name)
	log.Println("New path:", newPath)

	if element.Type == FolderType {
		doFolder(newPath)
	} else if element.Type == FileType {
		doFile(element, newPath)
	} else {
		log.Println("Element type not found", element.Type)
	}

	if element.Inner != nil {
		for _, innerElement := range element.Inner {
			doFolderCreation(*innerElement, newPath)
		}
	}
}

func doFolder(newPath string) {
	err := os.MkdirAll(newPath, 0777)
	if err != nil {
		log.Fatal("Error while creating folder", newPath, "err:", err.Error())
	}
}

func doFile(element Element, newPath string) {
	file, err := os.Create(newPath)
	if err != nil {
		log.Fatal("Error while creating file", newPath, "err:", err.Error())
	}

	switch element.Name {
	case "build.sh":
		file.WriteString(buildsh)
		break
	case "deploy.sh":
		file.WriteString(deploysh)
		break
	case ".gitignore":
		file.WriteString(gitignore)
		break
	}
}
