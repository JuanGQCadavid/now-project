import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/services/filter_service/service/dtos/filterSpotsResponse.dart';

class FilterServiceDTOsMappers {
  List<Spot> fromPlacesToSpotList(Locations locations) {
    List<Spot> spots = [];

    locations.places.forEach((place) {
      spots.add(
        Spot.withOutSpotColors(
          principalTag: place.eventInfo.name,
          secondaryTags: [],
          latLng: LatLng(
            place.placeInfo.lat,
            place.placeInfo.lon,
          ),
          spotId: place.eventInfo.id,
        ),
      );
    });

    return spots;
  }
}
