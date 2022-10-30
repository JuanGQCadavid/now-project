package folders

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/JuanGQCadavid/now-project/helpers/servicesCreator/internal/domain"
	"gopkg.in/yaml.v2"
)

type FolderStructure struct {
	Root   []Element    `yaml:"root"`
	Config FolderConfig `yaml:"config"`
}
type FolderConfig struct {
	GoModRootPath   string   `yaml:"goModRootPath"`
	InitialPackages []string `yaml:"initialPackages"`
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

	log.Println("Creating folders")
	for _, element := range folderStructure.Root {
		doFolderCreation(element, root)
	}
	log.Println("Folders creation done.")

	doConfig(serviceInfo, folderStructure.Config, root)
}

func doConfig(serviceInfo *domain.ServiceInformation, config FolderConfig, root string) {

	goInitCMD := exec.Command("go", "mod", "init", fmt.Sprintf("%s/%s", config.GoModRootPath, root))
	goInitCMD.Dir = root
	out, err := goInitCMD.CombinedOutput()

	log.Println(string(out))
	if err != nil {
		log.Fatalln("And erro happen while running go int, err: ", err.Error())
	}

	for _, packag := range config.InitialPackages {
		goInstall := exec.Command("go", "get", packag)
		goInstall.Dir = root
		out, err := goInstall.CombinedOutput()

		log.Println(string(out))
		if err != nil {
			log.Fatalln("And erro happen while running go install, error: ", err.Error())
		}
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
