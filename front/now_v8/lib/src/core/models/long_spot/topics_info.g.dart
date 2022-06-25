// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'topics_info.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_TopicsInfo _$$_TopicsInfoFromJson(Map<String, dynamic> json) =>
    _$_TopicsInfo(
      principalTag: json['principalTag'] as String,
      secondaryTags: (json['secondaryTags'] as List<dynamic>)
          .map((e) => e as String)
          .toList(),
    );

Map<String, dynamic> _$$_TopicsInfoToJson(_$_TopicsInfo instance) =>
    <String, dynamic>{
      'principalTag': instance.principalTag,
      'secondaryTags': instance.secondaryTags,
    };
