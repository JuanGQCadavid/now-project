import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'backend_errors.freezed.dart';

@freezed
class BackendErrors with _$BackendErrors{
  const factory BackendErrors.internalError() = InternalError;
  const factory BackendErrors.resourceNotFound() = ResourceNotFound;
  const factory BackendErrors.serviceUnavailable() = ServiceUnavailable;
  const factory BackendErrors.noInternetConnection() = NoInternetConnection;
  const factory BackendErrors.badResponseFormat() = BadResponseFormat;
  const factory BackendErrors.unknownError() = UnknownError;

}

