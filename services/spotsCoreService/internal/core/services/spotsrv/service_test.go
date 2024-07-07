package spotsrv

import (
	"testing"

	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/ports"
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/neo4jRepository"
	"github.com/stretchr/testify/mock"
)

type spotsRepositoryMock struct {
	*neo4jRepository.Neo4jSpotRepo
	mock.Mock
}

func (m *spotsRepositoryMock) Get(id string, format ports.OutputFormat) (domain.Spot, error) {

	m.Called(id, format)
	return domain.Spot{EventInfo: domain.Event{Name: "Mah dude"}}, nil
}

func TestService_Get(t *testing.T) {
	testObj := new(spotsRepositoryMock)
	// testObj.On("Get", "1", ports.FULL_FORMAT).Return(domain.Spot{EventInfo: domain.Event{Name: "Hi"}}, nil)
	//
	a, _ := testObj.Get("1", ports.FULL_FORMAT)

	t.Log(a.EventInfo.Name)

}

func TestService_GetSpotsByDatesIds(t *testing.T) {

}
func TestService_CreateSpot(t *testing.T) {

}
func TestService_UpdateSpotEvent(t *testing.T) {

}
func TestService_UpdateSpotTopic(t *testing.T) {

}
func TestService_UpdateSpotPlace(t *testing.T) {

}
