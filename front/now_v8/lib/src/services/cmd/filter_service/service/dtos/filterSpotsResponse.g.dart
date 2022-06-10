// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'filterSpotsResponse.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Locations _$LocationsFromJson(Map<String, dynamic> json) => Locations(
      places: (json['places'] as List<dynamic>)
          .map((e) => FilterSpot.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$LocationsToJson(Locations instance) => <String, dynamic>{
      'places': instance.places.map((e) => e.toJson()).toList(),
    };

FilterSpot _$FilterSpotFromJson(Map<String, dynamic> json) => FilterSpot(
      eventInfo: EventInfo.fromJson(json['eventInfo'] as Map<String, dynamic>),
      placeInfo: PlaceInfo.fromJson(json['placeInfo'] as Map<String, dynamic>),
      topicInfo: TopicInfo.fromJson(json['topicInfo'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$FilterSpotToJson(FilterSpot instance) =>
    <String, dynamic>{
      'eventInfo': instance.eventInfo,
      'placeInfo': instance.placeInfo,
      'topicInfo': instance.topicInfo,
    };

EventInfo _$EventInfoFromJson(Map<String, dynamic> json) => EventInfo(
      name: json['name'] as String,
      id: json['id'] as String,
      eventType: json['eventType'] as String,
      emoji: json['emoji'] as String,
    );

Map<String, dynamic> _$EventInfoToJson(EventInfo instance) => <String, dynamic>{
      'name': instance.name,
      'id': instance.id,
      'eventType': instance.eventType,
      'emoji': instance.emoji,
    };

PlaceInfo _$PlaceInfoFromJson(Map<String, dynamic> json) => PlaceInfo(
      mapProviderId: json['mapProviderId'] as String,
      lat: (json['lat'] as num).toDouble(),
      lon: (json['lon'] as num).toDouble(),
    );

Map<String, dynamic> _$PlaceInfoToJson(PlaceInfo instance) => <String, dynamic>{
      'mapProviderId': instance.mapProviderId,
      'lat': instance.lat,
      'lon': instance.lon,
    };

TopicInfo _$TopicInfoFromJson(Map<String, dynamic> json) => TopicInfo(
      principalTopic: json['principalTopic'] as String? ?? "",
      secondaryTopics: (json['secondaryTopics'] as List<dynamic>?)
              ?.map((e) => e as String)
              .toList() ??
          const [],
    );

Map<String, dynamic> _$TopicInfoToJson(TopicInfo instance) => <String, dynamic>{
      'principalTopic': instance.principalTopic,
      'secondaryTopics': instance.secondaryTopics,
    };
