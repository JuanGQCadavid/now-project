// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'create_spot_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

CreateSpotResponse _$CreateSpotResponseFromJson(Map<String, dynamic> json) =>
    CreateSpotResponse(
      eventInfo: EventInfo.fromJson(json['eventInfo'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$CreateSpotResponseToJson(CreateSpotResponse instance) =>
    <String, dynamic>{
      'eventInfo': instance.eventInfo,
    };

EventInfo _$EventInfoFromJson(Map<String, dynamic> json) => EventInfo(
      id: json['id'] as String,
    );

Map<String, dynamic> _$EventInfoToJson(EventInfo instance) => <String, dynamic>{
      'id': instance.id,
    };
