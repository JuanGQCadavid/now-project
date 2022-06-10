import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/dtos/filterSpotsResponse.dart';

class FilterServiceDTOsMappers {
  List<Spot> fromPlacesToSpotList(Locations locations) {
    List<Spot> spots = [];

    locations.places.forEach((place) {
      print(place.topicInfo.principalTopic);
      print(place.topicInfo.secondaryTopics);

      spots.add(
        Spot.withOutSpotColors(
          principalTag: place.topicInfo.principalTopic.isNotEmpty ||
                  place.topicInfo.secondaryTopics.isNotEmpty
              ? place.topicInfo.principalTopic
              : place.eventInfo.name.toLowerCase().replaceAll(RegExp(r' '), ""),
          secondaryTags: place.topicInfo.secondaryTopics,
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
