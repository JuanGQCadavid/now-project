// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'event_info.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_EventInfo _$$_EventInfoFromJson(Map<String, dynamic> json) => _$_EventInfo(
      name: json['name'] as String,
      id: json['id'] as String,
      description: json['description'] as String,
      maximunCapacty: json['maximunCapacty'] as int,
      eventType: json['eventType'] as String,
      emoji: json['emoji'] as String,
    );

Map<String, dynamic> _$$_EventInfoToJson(_$_EventInfo instance) =>
    <String, dynamic>{
      'name': instance.name,
      'id': instance.id,
      'description': instance.description,
      'maximunCapacty': instance.maximunCapacty,
      'eventType': instance.eventType,
      'emoji': instance.emoji,
    };
