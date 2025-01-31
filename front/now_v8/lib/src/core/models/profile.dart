import 'package:json_annotation/json_annotation.dart';

part 'profile.g.dart';

@JsonSerializable(explicitToJson: true)
class UserProfile {
  // User name
  @JsonKey(name: "userName", required: true)
  final String userName;

  // First + last name
  @JsonKey(name: "first_name", defaultValue: "")
  final String firstName;

  @JsonKey(name: "last_name", defaultValue: "")
  final String lastName;

  @JsonKey(name: "is_first_last_name_public", defaultValue: false)
  final bool isFirstLastNamePublic;

  // Phone number
  @JsonKey(name: "phone_number", defaultValue: "")
  final String phoneNumber;

  @JsonKey(name: "is_phone_number_public", defaultValue: false)
  final bool isPhoneNumberPublic;

  UserProfile({
    required this.userName,
    required this.firstName,
    required this.lastName,
    required this.isFirstLastNamePublic,
    required this.phoneNumber,
    required this.isPhoneNumberPublic,
  });

  factory UserProfile.fromJson(Map<String, dynamic> json) =>
      _$UserProfileFromJson(json);

  Map<String, dynamic> toJson() => _$UserProfileToJson(this);
}
