
import 'package:google_maps_flutter/google_maps_flutter.dart';


import '../../../../core/contracts/filterService.dart';
import '../../../../core/models/long_spot.dart';
import '../../../../core/models/long_spot/event_info.dart';
import '../../../../core/models/long_spot/host_info.dart';
import '../../../../core/models/long_spot/place_info.dart';
import '../../../../core/models/long_spot/topics_info.dart';
import '../../../../core/models/spot.dart';
import '../../../../core/models/state_response.dart';


// What about a provider with configuration abput the server api gateway ?
class FilterFakeService implements IFilterService {
  FilterFakeService();

  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10,
  }) async {
    List<Spot> spots = [
      Spot.withOutSpotColors(
        principalTag: "CoffeLovers",
        secondaryTags: List.empty(),
        latLng: LatLng(6.253723, -75.592771),
        spotId: "spotId_1",
      ),
      Spot.withOutSpotColors(
        principalTag: "ReadingClub",
        secondaryTags: List.empty(),
        latLng: LatLng(6.255733, -75.592771),
        spotId: "spotId_2",
      ),
      Spot.withOutSpotColors(
          principalTag: "StreePainting",
          secondaryTags: List.empty(),
          latLng: LatLng(6.257743, -75.592771),
          spotId: "spotId_3"),
      Spot.withOutSpotColors(
        principalTag: "Dance",
        secondaryTags: List.empty(),
        latLng: LatLng(6.259753, -75.592771),
        spotId: "spotId_4",
      )
    ];

    return spots;
  }

  Future<StateResponse<List<LongSpot>, String>> getByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 10,
    String token = "",
  }) async {
    List<LongSpot> longSpots = [
      const LongSpot(
        eventInfo: EventInfo(
          description: "Spot 1 Description",
          emoji: ":p",
          eventType: "online",
          id: "SPOT_ID_1",
          maximunCapacty: 1,
          name: "Spot 1 Name"
        ),
        hostInfo: HostInfo(
          name: "Juan Gonzalo",
        ),
        placeInfo: PlaceInfo(
          name: "Spot 1 Place name",
          lat: 6.253723,
          lon: -75.592771,
          mapProviderId: "MAP_PROV_ID_1",
        ),
        topicInfo: TopicsInfo(
          principalTag: "Spot1principalTag",
          secondaryTags: ["Spot1Secondarytags"]
        ),
      ),

      const LongSpot(
        eventInfo: EventInfo(
          description: "Spot 2 Description",
          emoji: ":o",
          eventType: "online",
          id: "SPOT_ID_2",
          maximunCapacty: 1,
          name: "Spot 2 Name"
        ),
        hostInfo: HostInfo(
          name: "Adriana Lucia",
        ),
        placeInfo: PlaceInfo(
          name: "Spot 2 Place name",
          lat: 6.255733,
          lon: -75.592771,
          mapProviderId: "MAP_PROV_ID_2",
        ),
        topicInfo: TopicsInfo(
          principalTag: "Spot2principalTag",
          secondaryTags: ["Spot2Secondarytags"]
        ),
      ),

      const LongSpot(
        eventInfo: EventInfo(
          description: "Spot 3 Description",
          emoji: ":d",
          eventType: "online",
          id: "SPOT_ID_3",
          maximunCapacty: 1,
          name: "Spot 3 Name"
        ),
        hostInfo: HostInfo(
          name: "Sandra Patricia",
        ),
        placeInfo: PlaceInfo(
          name: "Spot 3 Place name",
          lat: 6.257743,
          lon: -75.592771,
          mapProviderId: "MAP_PROV_ID_3",
        ),
        topicInfo: TopicsInfo(
          principalTag: "Spot3principalTag",
          secondaryTags: ["Spot3Secondarytags"]
        ),
      ),

      const LongSpot(
        eventInfo: EventInfo(
          description: "Spot 4 Description",
          emoji: ":v",
          eventType: "online",
          id: "SPOT_ID_4",
          maximunCapacty: 1,
          name: "Spot 4 Name"
        ),
        hostInfo: HostInfo(
          name: "Valeria Serna",
        ),
        placeInfo: PlaceInfo(
          name: "Spot 4 Place name",
          lat: 6.259753,
          lon: -75.592771,
          mapProviderId: "MAP_PROV_ID_4",
        ),
        topicInfo: TopicsInfo(
          principalTag: "Spot4principalTag",
          secondaryTags: ["Spot4Secondarytags"]
        ),
      ),
    ];

    return StateResponse<List<LongSpot>, String>(
        response: longSpots, token: "FakeToken:(");
  }
}
