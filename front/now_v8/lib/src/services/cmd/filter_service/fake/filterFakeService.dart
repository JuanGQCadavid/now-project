import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/models/spot.dart';

// What about a provider with configuration abput the server api gateway ?
class FilterFakeService implements IFilterService {
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

  FilterFakeService();

  Future<List<Spot>> getByProximity(
      {required double cpLat, required double cpLng, double radious = 10})  async {
    return spots;
  }
}
