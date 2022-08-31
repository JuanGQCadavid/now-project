package spotservicelambda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/filter/internal/core/ports"
)

type SpotServiceLambda struct {
	SpotServiceURL string
	GetSpotsURI    string
}

func NewSpotServiceLambda() *SpotServiceLambda {

	spotURL, isPresent := os.LookupEnv("spotServiceURL")

	if !isPresent {
		log.Fatal("spotServiceURL is not present in the env")
	}

	return &SpotServiceLambda{
		SpotServiceURL: spotURL,
		GetSpotsURI:    "getSpots",
	}
}

func (srv *SpotServiceLambda) GetSpotsCardsInfo(spots []string, format ports.OutputFormat) ([]domain.Spot, error) {

	body, err := json.Marshal(map[string]interface{}{
		"spotIds": spots,
	})

	log.Printf("Request Body: \n\t%s\n", body)

	if err != nil {
		log.Fatalln("An error while marshalling the body: ", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/%s?format=%s", srv.SpotServiceURL, srv.GetSpotsURI, string(format)), "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatalln("An error while making the request: ", err)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("An error while Reading the resquest: ", err)
	}

	spotResponse := SpotResponse{}

	unmarshalError := json.Unmarshal(responseBody, &spotResponse)

	if unmarshalError != nil {
		log.Fatalln("An error while Unmarshal the resquest: ", err)
	}

	log.Printf("Response: \n\t%+v\n", spotResponse)

	return spotResponse.Spots, nil
}
