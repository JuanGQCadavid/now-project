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

InitSingUp _$InitSingUpFromJson(Map<String, dynamic> json) {
  $checkKeys(
    json,
    requiredKeys: const ['phoneNumber', 'userName', 'methodVerificator'],
  );
  return InitSingUp(
    json['phoneNumber'] as String,
    json['userName'] as String,
    MethodVerificator.fromJson(
        json['methodVerificator'] as Map<String, dynamic>),
  );
}

Map<String, dynamic> _$InitSingUpToJson(InitSingUp instance) =>
    <String, dynamic>{
      'phoneNumber': instance.phoneNumber,
      'userName': instance.userName,
      'methodVerificator': instance.methodVerificator.toJson(),
    };

ValidateOTP _$ValidateOTPFromJson(Map<String, dynamic> json) {
  $checkKeys(
    json,
    requiredKeys: const ['phoneNumber', 'code'],
  );
  return ValidateOTP(
    json['phoneNumber'] as String,
    (json['code'] as List<dynamic>).map((e) => e as int).toList(),
  );
}

Map<String, dynamic> _$ValidateOTPToJson(ValidateOTP instance) =>
    <String, dynamic>{
      'phoneNumber': instance.userPhoneNumber,
      'code': instance.userCode,
    };
