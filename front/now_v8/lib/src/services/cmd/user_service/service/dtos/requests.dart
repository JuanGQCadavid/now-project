import 'package:json_annotation/json_annotation.dart';

part 'requests.g.dart';

@JsonSerializable(explicitToJson: true)
class MethodVerificator {
  @JsonKey(name: "language", defaultValue: "en")
  String language;

  @JsonKey(name: "sms", defaultValue: false)
  bool sms;

  @JsonKey(name: "whatsapp", defaultValue: false)
  bool whatsapp;

  MethodVerificator(this.language, {this.sms = false, this.whatsapp = false});

  factory MethodVerificator.fromJson(Map<String, dynamic> json) =>
      _$MethodVerificatorFromJson(json);

  Map<String, dynamic> toJson() => _$MethodVerificatorToJson(this);
}

@JsonSerializable(explicitToJson: true)
class InitLogin {
  @JsonKey(name: "phoneNumber", required: true)
  String phoneNumber;

  @JsonKey(name: "methodVerificator", required: true)
  MethodVerificator methodVerificator;

  InitLogin(this.phoneNumber, this.methodVerificator);

  factory InitLogin.fromJson(Map<String, dynamic> json) =>
      _$InitLoginFromJson(json);

  Map<String, dynamic> toJson() => _$InitLoginToJson(this);
}

@JsonSerializable(explicitToJson: true)
class InitSingUp {
  @JsonKey(name: "phoneNumber", required: true)
  String phoneNumber;

  @JsonKey(name: "userName", required: true)
  String userName;

  @JsonKey(name: "methodVerificator", required: true)
  MethodVerificator methodVerificator;

  InitSingUp(this.phoneNumber, this.userName, this.methodVerificator);

  factory InitSingUp.fromJson(Map<String, dynamic> json) =>
      _$InitSingUpFromJson(json);

  Map<String, dynamic> toJson() => _$InitSingUpToJson(this);
}

@JsonSerializable(explicitToJson: true)
class ValidateOTP {
  @JsonKey(name: "phoneNumber", required: true)
  String userPhoneNumber;

  @JsonKey(name: "code", required: true)
  List<int> userCode;

  ValidateOTP(this.userPhoneNumber, this.userCode);

  factory ValidateOTP.fromJson(Map<String, dynamic> json) =>
      _$ValidateOTPFromJson(json);

  Map<String, dynamic> toJson() => _$ValidateOTPToJson(this);
}
