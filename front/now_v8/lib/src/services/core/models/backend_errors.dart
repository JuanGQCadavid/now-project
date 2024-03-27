import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';
// import 'package:json_annotation/json_annotation.dart';

part 'backend_errors.g.dart';
part 'backend_errors.freezed.dart';

@freezed
class BackendErrors with _$BackendErrors {
  const factory BackendErrors.internalError() = InternalError;
  const factory BackendErrors.resourceNotFound() = ResourceNotFound;
  const factory BackendErrors.serviceUnavailable() = ServiceUnavailable;
  const factory BackendErrors.noInternetConnection() = NoInternetConnection;
  const factory BackendErrors.badResponseFormat() = BadResponseFormat;
  const factory BackendErrors.unknownError() = UnknownError;
  const factory BackendErrors.clientError(ErrorMessage errorMessage) =
      ClientError;
}

@JsonSerializable(explicitToJson: true)
class ErrorMessage {
  @JsonKey(name: "id") //@JsonKey(name: "id", required: true)
  String id;

  @JsonKey(name: "message")
  String message;

  @JsonKey(name: "internalError")
  String internalError;

  ErrorMessage(this.id, this.message, this.internalError);

  @override
  String toString() {
    return 'ErrrorMesaage:\n\t ID: $id\n\t Message: $message\n\t InternalError: $internalError';
  }

  factory ErrorMessage.fromJson(Map<String, dynamic> json) =>
      _$ErrorMessageFromJson(json);

  Map<String, dynamic> toJson() => _$ErrorMessageToJson(this);
}
