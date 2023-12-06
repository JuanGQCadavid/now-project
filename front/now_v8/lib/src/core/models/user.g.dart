// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'user.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

UserDetails _$UserDetailsFromJson(Map<String, dynamic> json) {
  $checkKeys(
    json,
    requiredKeys: const [
      'userId',
      'refreshToken',
      'shortLiveToken',
      'shortLiveTokenTTL'
    ],
  );
  return UserDetails(
    userId: json['userId'] as String,
    userName: json['userName'] as String? ?? '',
    refreshToken: json['refreshToken'] as String,
    shortLiveToken: json['shortLiveToken'] as String,
    shortLiveTokenTTL: json['shortLiveTokenTTL'] as String,
  );
}

Map<String, dynamic> _$UserDetailsToJson(UserDetails instance) =>
    <String, dynamic>{
      'userId': instance.userId,
      'userName': instance.userName,
      'refreshToken': instance.refreshToken,
      'shortLiveToken': instance.shortLiveToken,
      'shortLiveTokenTTL': instance.shortLiveTokenTTL,
    };
