import "package:json_annotation/json_annotation.dart";

part 'spot.g.dart';

@JsonSerializable()
class Locations {
  final List<Spot> places;

  Locations({required this.places});

  factory Locations.fromJson(Map<String, dynamic> json) =>
      _$LocationsFromJson(json);
}

@JsonSerializable()
class Spot {
  final EventInfo eventInfo;
  final PlaceInfo placeInfo;

  Spot({
    required this.eventInfo,
    required this.placeInfo,
  });

  factory Spot.fromJson(Map<String, dynamic> json) => _$SpotFromJson(json);
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
}
