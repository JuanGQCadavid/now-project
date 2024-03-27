// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'backend_errors.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ErrorMessage _$ErrorMessageFromJson(Map<String, dynamic> json) => ErrorMessage(
      json['id'] as String,
      json['message'] as String,
      json['internalError'] as String,
    );

Map<String, dynamic> _$ErrorMessageToJson(ErrorMessage instance) =>
    <String, dynamic>{
      'id': instance.id,
      'message': instance.message,
      'internalError': instance.internalError,
    };
