// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'long_spot.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$LongSpotImpl _$$LongSpotImplFromJson(Map<String, dynamic> json) =>
    _$LongSpotImpl(
      dateInfo: DateInfo.fromJson(json['dateInfo'] as Map<String, dynamic>),
      eventInfo: EventInfo.fromJson(json['eventInfo'] as Map<String, dynamic>),
      hostInfo: HostInfo.fromJson(json['hostInfo'] as Map<String, dynamic>),
      placeInfo: PlaceInfo.fromJson(json['placeInfo'] as Map<String, dynamic>),
      topicInfo: TopicsInfo.fromJson(json['topicInfo'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$LongSpotImplToJson(_$LongSpotImpl instance) =>
    <String, dynamic>{
      'dateInfo': instance.dateInfo,
      'eventInfo': instance.eventInfo,
      'hostInfo': instance.hostInfo,
      'placeInfo': instance.placeInfo,
      'topicInfo': instance.topicInfo,
    };

_$TopicsInfoImpl _$$TopicsInfoImplFromJson(Map<String, dynamic> json) =>
    _$TopicsInfoImpl(
      principalTopic: json['principalTopic'] as String,
      secondaryTopics: (json['secondaryTopics'] as List<dynamic>)
          .map((e) => e as String)
          .toList(),
    );

Map<String, dynamic> _$$TopicsInfoImplToJson(_$TopicsInfoImpl instance) =>
    <String, dynamic>{
      'principalTopic': instance.principalTopic,
      'secondaryTopics': instance.secondaryTopics,
    };

_$PlaceInfoImpl _$$PlaceInfoImplFromJson(Map<String, dynamic> json) =>
    _$PlaceInfoImpl(
      name: json['name'] as String,
      lat: (json['lat'] as num).toDouble(),
      lon: (json['lon'] as num).toDouble(),
      mapProviderId: json['mapProviderId'] as String,
    );

Map<String, dynamic> _$$PlaceInfoImplToJson(_$PlaceInfoImpl instance) =>
    <String, dynamic>{
      'name': instance.name,
      'lat': instance.lat,
      'lon': instance.lon,
      'mapProviderId': instance.mapProviderId,
    };

_$HostInfoImpl _$$HostInfoImplFromJson(Map<String, dynamic> json) =>
    _$HostInfoImpl(
      name: json['name'] as String,
    );

Map<String, dynamic> _$$HostInfoImplToJson(_$HostInfoImpl instance) =>
    <String, dynamic>{
      'name': instance.name,
    };

_$EventInfoImpl _$$EventInfoImplFromJson(Map<String, dynamic> json) =>
    _$EventInfoImpl(
      name: json['name'] as String,
      id: json['id'] as String,
      description: json['description'] as String,
      maximunCapacty: (json['maximunCapacty'] as num).toInt(),
      emoji: json['emoji'] as String,
    );

Map<String, dynamic> _$$EventInfoImplToJson(_$EventInfoImpl instance) =>
    <String, dynamic>{
      'name': instance.name,
      'id': instance.id,
      'description': instance.description,
      'maximunCapacty': instance.maximunCapacty,
      'emoji': instance.emoji,
    };

_$DateInfoImpl _$$DateInfoImplFromJson(Map<String, dynamic> json) =>
    _$DateInfoImpl(
      dateTime: json['dateTime'] as String,
      id: json['id'] as String,
      startTime: json['startTime'] as String,
      durationApproximatedInSeconds:
          (json['durationApproximatedInSeconds'] as num).toInt(),
    );

Map<String, dynamic> _$$DateInfoImplToJson(_$DateInfoImpl instance) =>
    <String, dynamic>{
      'dateTime': instance.dateTime,
      'id': instance.id,
      'startTime': instance.startTime,
      'durationApproximatedInSeconds': instance.durationApproximatedInSeconds,
    };
