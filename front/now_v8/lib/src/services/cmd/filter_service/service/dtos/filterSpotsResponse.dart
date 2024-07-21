import "package:json_annotation/json_annotation.dart";

part 'filterSpotsResponse.g.dart';

@JsonSerializable(explicitToJson: true)
class FilterProxymityResponse {
  final Locations result;

  FilterProxymityResponse({required this.result});

  factory FilterProxymityResponse.fromJson(Map<String, dynamic> json) =>
      _$FilterProxymityResponseFromJson(json);

  Map<String, dynamic> toJson() => _$FilterProxymityResponseToJson(this);
}

@JsonSerializable(explicitToJson: true)
class FilterProxyResponseWithState {
  final Locations result;
  final SearchSession search_session;

  FilterProxyResponseWithState(
      {required this.result, required this.search_session});

  factory FilterProxyResponseWithState.fromJson(Map<String, dynamic> json) =>
      _$FilterProxyResponseWithStateFromJson(json);

  Map<String, dynamic> toJson() => _$FilterProxyResponseWithStateToJson(this);
}

@JsonSerializable(explicitToJson: true)
class SearchSession {
  final SessionDetails session_details;
  SearchSession({required this.session_details});

  factory SearchSession.fromJson(Map<String, dynamic> json) =>
      _$SearchSessionFromJson(json);

  Map<String, dynamic> toJson() => _$SearchSessionToJson(this);
}

@JsonSerializable(explicitToJson: true)
class SessionDetails {
  final String session_id;
  final String header_name;
  final int ttl;

  SessionDetails({
    this.session_id = "",
    this.header_name = "",
    this.ttl = 0,
  });

  factory SessionDetails.fromJson(Map<String, dynamic> json) =>
      _$SessionDetailsFromJson(json);

  Map<String, dynamic> toJson() => _$SessionDetailsToJson(this);
}

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
  final TopicInfo topicInfo;
  final DateInfo dateInfo;
  final HostInfo hostInfo;

  FilterSpot({
    required this.eventInfo,
    required this.placeInfo,
    required this.topicInfo,
    required this.dateInfo,
    required this.hostInfo, // This could crash :'()
  });

  factory FilterSpot.fromJson(Map<String, dynamic> json) =>
      _$FilterSpotFromJson(json);

  Map<String, dynamic> toJson() => _$FilterSpotToJson(this);
}

@JsonSerializable()
class HostInfo {

  final String id;
  final String name;

  HostInfo({
    this.name = "",
    this.id = ""
  });

  factory HostInfo.fromJson(Map<String, dynamic> json) =>
      _$HostInfoFromJson(json);

  Map<String, dynamic> toJson() => _$HostInfoToJson(this);

}

// THis is an example of how to use it with empty values

// @JsonSerializable()
// class HostInfo {

//   final String id;
//   final String name;

//   HostInfo({
//     required this.name = ""
//   });

//   factory HostInfo.empty() => HostInfo();

//   factory HostInfo.fromJson(Map<String, dynamic> json) =>
//       _$HostInfoFromJson(json);

//   Map<String, dynamic> toJson() => _$HostInfoToJson(this);

// }

@JsonSerializable()
class EventInfo {
  final String name;
  final String id;
  final String emoji;
  final String description;

  EventInfo({
    required this.name,
    required this.id,
    required this.emoji,
    this.description = ""
  });

  factory EventInfo.fromJson(Map<String, dynamic> json) =>
      _$EventInfoFromJson(json);

  Map<String, dynamic> toJson() => _$EventInfoToJson(this);
}

@JsonSerializable()
class DateInfo {
  final String dateTime;
  final String id;
  final String startTime;
  final int durationApproximated;

  DateInfo({
    required this.dateTime,
    required this.id,
    required this.startTime,
    required this.durationApproximated,
  });

  factory DateInfo.fromJson(Map<String, dynamic> json) =>
      _$DateInfoFromJson(json);

  Map<String, dynamic> toJson() => _$DateInfoToJson(this);
}

@JsonSerializable()
class PlaceInfo {
  final String mapProviderId;
  final double lat;
  final double lon;
  final String name;

  PlaceInfo({
    required this.mapProviderId,
    required this.lat,
    required this.lon,
    this.name = ""
  });

  factory PlaceInfo.fromJson(Map<String, dynamic> json) =>
      _$PlaceInfoFromJson(json);

  Map<String, dynamic> toJson() => _$PlaceInfoToJson(this);
}

@JsonSerializable()
class TopicInfo {
  // @JsonKey(defaultValue: "")
  final String principalTopic;
  // @JsonKey(defaultValue: [])
  final List<String> secondaryTopics;

  TopicInfo({
    this.principalTopic = "",
    this.secondaryTopics = const [],
  });

  factory TopicInfo.fromJson(Map<String, dynamic> json) =>
      _$TopicInfoFromJson(json);

  Map<String, dynamic> toJson() => _$TopicInfoToJson(this);
}
