import 'package:json_annotation/json_annotation.dart';

part 'user.g.dart';

@JsonSerializable(explicitToJson: true)
class UserDetails {
  // User data
  @JsonKey(name: "userId", required: true)
  final String userId;

  @JsonKey(name: "userName", defaultValue: "")
  final String userName;

  // Tokens
  @JsonKey(name: "refreshToken", required: true)
  final String refreshToken;

  @JsonKey(name: "shortLiveToken", required: true)
  final String shortLiveToken;

  @JsonKey(name: "shortLiveTokenTTL", required: true)
  final String shortLiveTokenTTL;

  UserDetails({
    required this.userId,
    required this.userName,
    required this.refreshToken,
    required this.shortLiveToken,
    required this.shortLiveTokenTTL,
  });

  factory UserDetails.fromJson(Map<String, dynamic> json) =>
      _$UserDetailsFromJson(json);

  Map<String, dynamic> toJson() => _$UserDetailsToJson(this);
}
