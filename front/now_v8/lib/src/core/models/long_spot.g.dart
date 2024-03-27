// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'long_spot.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_LongSpot _$$_LongSpotFromJson(Map<String, dynamic> json) => _$_LongSpot(
      dateInfo: DateInfo.fromJson(json['dateInfo'] as Map<String, dynamic>),
      eventInfo: EventInfo.fromJson(json['eventInfo'] as Map<String, dynamic>),
      hostInfo: HostInfo.fromJson(json['hostInfo'] as Map<String, dynamic>),
      placeInfo: PlaceInfo.fromJson(json['placeInfo'] as Map<String, dynamic>),
      topicInfo: TopicsInfo.fromJson(json['topicInfo'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$_LongSpotToJson(_$_LongSpot instance) =>
    <String, dynamic>{
      'dateInfo': instance.dateInfo,
      'eventInfo': instance.eventInfo,
      'hostInfo': instance.hostInfo,
      'placeInfo': instance.placeInfo,
      'topicInfo': instance.topicInfo,
    };

_$_TopicsInfo _$$_TopicsInfoFromJson(Map<String, dynamic> json) =>
    _$_TopicsInfo(
      principalTopic: json['principalTopic'] as String,
      secondaryTopics: (json['secondaryTopics'] as List<dynamic>)
          .map((e) => e as String)
          .toList(),
    );

Map<String, dynamic> _$$_TopicsInfoToJson(_$_TopicsInfo instance) =>
    <String, dynamic>{
      'principalTopic': instance.principalTopic,
      'secondaryTopics': instance.secondaryTopics,
    };

_$_PlaceInfo _$$_PlaceInfoFromJson(Map<String, dynamic> json) => _$_PlaceInfo(
      name: json['name'] as String,
      lat: (json['lat'] as num).toDouble(),
      lon: (json['lon'] as num).toDouble(),
      mapProviderId: json['mapProviderId'] as String,
    );

Map<String, dynamic> _$$_PlaceInfoToJson(_$_PlaceInfo instance) =>
    <String, dynamic>{
      'name': instance.name,
      'lat': instance.lat,
      'lon': instance.lon,
      'mapProviderId': instance.mapProviderId,
    };

_$_HostInfo _$$_HostInfoFromJson(Map<String, dynamic> json) => _$_HostInfo(
      name: json['name'] as String,
    );

Map<String, dynamic> _$$_HostInfoToJson(_$_HostInfo instance) =>
    <String, dynamic>{
      'name': instance.name,
    };

_$_EventInfo _$$_EventInfoFromJson(Map<String, dynamic> json) => _$_EventInfo(
      name: json['name'] as String,
      id: json['id'] as String,
      description: json['description'] as String,
      maximunCapacty: json['maximunCapacty'] as int,
      emoji: json['emoji'] as String,
    );

Map<String, dynamic> _$$_EventInfoToJson(_$_EventInfo instance) =>
    <String, dynamic>{
      'name': instance.name,
      'id': instance.id,
      'description': instance.description,
      'maximunCapacty': instance.maximunCapacty,
      'emoji': instance.emoji,
    };

_$_DateInfo _$$_DateInfoFromJson(Map<String, dynamic> json) => _$_DateInfo(
      dateTime: json['dateTime'] as String,
      id: json['id'] as String,
      startTime: json['startTime'] as String,
      durationApproximatedInSeconds:
          json['durationApproximatedInSeconds'] as int,
    );

Map<String, dynamic> _$$_DateInfoToJson(_$_DateInfo instance) =>
    <String, dynamic>{
      'dateTime': instance.dateTime,
      'id': instance.id,
      'startTime': instance.startTime,
      'durationApproximatedInSeconds': instance.durationApproximatedInSeconds,
    };
