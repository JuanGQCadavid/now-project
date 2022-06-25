// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'long_spot.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_LongSpot _$$_LongSpotFromJson(Map<String, dynamic> json) => _$_LongSpot(
      eventInfo: EventInfo.fromJson(json['eventInfo'] as Map<String, dynamic>),
      hostInfo: HostInfo.fromJson(json['hostInfo'] as Map<String, dynamic>),
      placeInfo: PlaceInfo.fromJson(json['placeInfo'] as Map<String, dynamic>),
      topicInfo: TopicsInfo.fromJson(json['topicInfo'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$_LongSpotToJson(_$_LongSpot instance) =>
    <String, dynamic>{
      'eventInfo': instance.eventInfo,
      'hostInfo': instance.hostInfo,
      'placeInfo': instance.placeInfo,
      'topicInfo': instance.topicInfo,
    };
