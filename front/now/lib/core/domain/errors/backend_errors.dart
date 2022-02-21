import 'package:freezed_annotation/freezed_annotation.dart';

part 'backend_errors.freezed.dart';

@freezed
class BackendErrors with _$BackendErrors {
  const BackendErrors._();
  const factory BackendErrors.noConnection() = NoConnection;
  const factory BackendErrors.internalError() = InternalError;
}
