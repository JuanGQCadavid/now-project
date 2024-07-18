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
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
)

type SpotServiceLambda struct {
	SpotServiceURL string
	GetSpotsURI    string
}

const (
	SPOT_URL = "spotServiceURL"
)

func NewSpotServiceLambda() (*SpotServiceLambda, error) {

	spotURL, isPresent := os.LookupEnv(SPOT_URL)

	if !isPresent {
		log.Println("[ERROR] spotServiceURL is not present in the env")
		return nil, ports.ErrMissingEnvParams
	}

	return &SpotServiceLambda{
		SpotServiceURL: spotURL,
		GetSpotsURI:    "bulk/fetch",
	}, nil
}

func (srv *SpotServiceLambda) GetSpotsCardsInfo(datesIds []string, format ports.OutputFormat) ([]domain.Spot, error) {
	log.Printf("GetSpotsCardsInfo | \ndatesIds:%+v ,\nformat:%s\n", datesIds, string(format))

	body, err := json.Marshal(map[string]interface{}{
		"datesIds": datesIds,
	})

	if err != nil {
		log.Println("[ERROR] An error while marshalling the body: ", err.Error())
		return nil, ports.ErrBodyRequestUnmarshal
	}

	//resp, err := http.Post(fmt.Sprintf("%s/%s?format=%s", srv.SpotServiceURL, srv.GetSpotsURI, string(format)), "application/json", bytes.NewBuffer(body))

	// FROM HERE
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s?format=%s", srv.SpotServiceURL, srv.GetSpotsURI, string(format)), bytes.NewBuffer(body))

	if err != nil {
		log.Println("[ERROR] An error while creating the req: ", err.Error())
		return nil, ports.ErrSendingRequest
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth", "AVOID")

	resp, err := client.Do(req)

	logs.Info.Println("Call to Spot core done, respond status.", resp.StatusCode)
	/// UNTIL HERE

	if err != nil {
		log.Println("[ERROR] An error while making the request: ", err.Error())
		return nil, ports.ErrSendingRequest
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("[ERROR] An error while Reading the resquest: ", err.Error())
		return nil, ports.ErrBodyResponseReadFail
	}

	spotResponse := SpotResponse{}

	unmarshalError := json.Unmarshal(responseBody, &spotResponse)

	if unmarshalError != nil {
		log.Println("[ERROR] An error while Unmarshal the resquest: ", unmarshalError)
		return nil, ports.ErrBodyResponseUnmarshal
	}

	// log.Printf("Response: \n\t%+v\n", spotResponse)
	return spotResponse.Spots, nil
}
