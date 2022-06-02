import "package:json_annotation/json_annotation.dart";

part 'filterSpotsResponse.g.dart';

@JsonSerializable(explicitToJson: true)
class Locations {
  final List<FilterSpot> places;

  Locations({required this.places});

  factory Locations.fromJson(Map<String, dynamic> json) =>
      _$LocationsFromJson(json);

  Map<String, dynamic> toJson() => _$LocationsToJson(this);

}

@JsonSerializable()
class FilterSpot {
  final EventInfo eventInfo;
  final PlaceInfo placeInfo;

  FilterSpot({
    required this.eventInfo,
    required this.placeInfo,
  });

  factory FilterSpot.fromJson(Map<String, dynamic> json) => _$FilterSpotFromJson(json);

  Map<String, dynamic> toJson() => _$FilterSpotToJson(this);
}


@JsonSerializable()
class EventInfo {
  final String name;
  final String id;
  final String eventType;
  final String emoji;

  EventInfo({
    required this.name,
    required this.id,
    required this.eventType,
    required this.emoji,
  });

  factory EventInfo.fromJson(Map<String, dynamic> json) =>
      _$EventInfoFromJson(json);

  Map<String, dynamic> toJson() => _$EventInfoToJson(this);
}

@JsonSerializable()
class PlaceInfo {
  final String mapProviderId;
  final double lat;
  final double lon;

  PlaceInfo({
    required this.mapProviderId,
    required this.lat,
    required this.lon,
  });

  factory PlaceInfo.fromJson(Map<String, dynamic> json) =>
      _$PlaceInfoFromJson(json);

  Map<String, dynamic> toJson() => _$PlaceInfoToJson(this);
}
