// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'profile.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

UserProfile _$UserProfileFromJson(Map<String, dynamic> json) {
  $checkKeys(
    json,
    requiredKeys: const ['userName'],
  );
  return UserProfile(
    userName: json['userName'] as String,
    firstName: json['first_name'] as String? ?? '',
    lastName: json['last_name'] as String? ?? '',
    isFirstLastNamePublic: json['is_first_last_name_public'] as bool? ?? false,
    phoneNumber: json['phone_number'] as String? ?? '',
    isPhoneNumberPublic: json['is_phone_number_public'] as bool? ?? false,
  );
}

Map<String, dynamic> _$UserProfileToJson(UserProfile instance) =>
    <String, dynamic>{
      'userName': instance.userName,
      'first_name': instance.firstName,
      'last_name': instance.lastName,
      'is_first_last_name_public': instance.isFirstLastNamePublic,
      'phone_number': instance.phoneNumber,
      'is_phone_number_public': instance.isPhoneNumberPublic,
    };
