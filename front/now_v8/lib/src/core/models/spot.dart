import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';

// I will need to investigate more about freezed in order to perfom copyWithData or something like that
// What about builder ?
class Spot {
  final String principalTag;
  final List<String> secondaryTags;
  final LatLng latLng;
  final String spotId;
  final DateTime date;

  SpotsColors spotsColor;

  Spot({
    required this.principalTag,
    required this.secondaryTags,
    required this.latLng,
    required this.spotId,
    required this.spotsColor,
    required this.date,
  });

  Spot.withOutSpotColors(
      {required this.principalTag,
      required this.secondaryTags,
      required this.latLng,
      required this.spotId,
      required this.date,
      this.spotsColor = const SpotsColors.empty()});

  // HERE BABE
  factory Spot.fromLongSpot(LongSpot longSpot) {   

    var timeFormatted = longSpot.dateInfo.startTime.split(" ");

    DateTime date; 

    if (timeFormatted.length == 1) {
        date = DateTime.parse("${longSpot.dateInfo.dateTime}T${longSpot.dateInfo.startTime}");
      } else{
        date = DateTime.parse("${timeFormatted[0]} ${timeFormatted[1]}");
      }

    return Spot.withOutSpotColors(
      principalTag: longSpot.topicInfo.principalTopic,
      secondaryTags: longSpot.topicInfo.secondaryTopics,
      latLng: LatLng(longSpot.placeInfo.lat, longSpot.placeInfo.lon),
      spotId: longSpot.eventInfo.id,
      date: date,
    );
  }
}
