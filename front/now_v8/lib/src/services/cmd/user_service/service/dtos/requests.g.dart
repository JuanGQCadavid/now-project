// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'requests.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

MethodVerificator _$MethodVerificatorFromJson(Map<String, dynamic> json) =>
    MethodVerificator(
      json['language'] as String? ?? 'en',
      sms: json['sms'] as bool? ?? false,
      whatsapp: json['whatsapp'] as bool? ?? false,
    );

Map<String, dynamic> _$MethodVerificatorToJson(MethodVerificator instance) =>
    <String, dynamic>{
      'language': instance.language,
      'sms': instance.sms,
      'whatsapp': instance.whatsapp,
    };

InitLogin _$InitLoginFromJson(Map<String, dynamic> json) {
  $checkKeys(
    json,
    requiredKeys: const ['phoneNumber', 'methodVerificator'],
  );
  return InitLogin(
    json['phoneNumber'] as String,
    MethodVerificator.fromJson(
        json['methodVerificator'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$InitLoginToJson(InitLogin instance) => <String, dynamic>{
      'phoneNumber': instance.phoneNumber,
      'methodVerificator': instance.methodVerificator.toJson(),
    };
