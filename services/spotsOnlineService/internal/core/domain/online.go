// {
//     "dateInfo": {
//         "durationApproximated": "HH:MM",
//         "maximunCapacty": 58
//     },
//     "placeInfo": {
//         "name": "Place Laureles",
//         "lat": 6.245887,
//         "lon": -75.589868,
//         "mapProviderId": "Place_Laureles"
//     }
// }

package domain

type OnlineSpot struct {
	SpotInfo  Spot       `json:"spotInfo"`
	DatesInfo []SpotDate `json:"dateInfo"`
	PlaceInfo Place      `json:"placeInfo"`
}
