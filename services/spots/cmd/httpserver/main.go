package httpserver

import (
	"github.com/JuanGQCadavid/now-project/services/spots/internal/core/services/spotsrv"
	"github.com/JuanGQCadavid/now-project/services/spots/internal/repositories/menRepository"
	"github.com/JuanGQCadavid/now-project/services/spots/pkg/uuidgen"
	"github.com/gin-gonic/gin"
)

func main() {
	repoSpot := menRepository.New()
	repoLocation := menRepository.NewLocationRepository()
	uuid := uuidgen.New()

	service := spotsrv.New(repoSpot, repoLocation, uuid)

	router := gin.Default()

	router.GET("/")

}
