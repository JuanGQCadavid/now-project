// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'backend_errors.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$BackendErrors {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $BackendErrorsCopyWith<$Res> {
  factory $BackendErrorsCopyWith(
          BackendErrors value, $Res Function(BackendErrors) then) =
      _$BackendErrorsCopyWithImpl<$Res, BackendErrors>;
}

/// @nodoc
class _$BackendErrorsCopyWithImpl<$Res, $Val extends BackendErrors>
    implements $BackendErrorsCopyWith<$Res> {
  _$BackendErrorsCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;
}

/// @nodoc
abstract class _$$InternalErrorCopyWith<$Res> {
  factory _$$InternalErrorCopyWith(
          _$InternalError value, $Res Function(_$InternalError) then) =
      __$$InternalErrorCopyWithImpl<$Res>;
}

/// @nodoc
class __$$InternalErrorCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$InternalError>
    implements _$$InternalErrorCopyWith<$Res> {
  __$$InternalErrorCopyWithImpl(
      _$InternalError _value, $Res Function(_$InternalError) _then)
      : super(_value, _then);
}

/// @nodoc

class _$InternalError with DiagnosticableTreeMixin implements InternalError {
  const _$InternalError();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.internalError()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'BackendErrors.internalError'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$InternalError);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return internalError();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return internalError?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return internalError(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return internalError?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError(this);
    }
    return orElse();
  }
}

abstract class InternalError implements BackendErrors {
  const factory InternalError() = _$InternalError;
}

/// @nodoc
abstract class _$$ResourceNotFoundCopyWith<$Res> {
  factory _$$ResourceNotFoundCopyWith(
          _$ResourceNotFound value, $Res Function(_$ResourceNotFound) then) =
      __$$ResourceNotFoundCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ResourceNotFoundCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$ResourceNotFound>
    implements _$$ResourceNotFoundCopyWith<$Res> {
  __$$ResourceNotFoundCopyWithImpl(
      _$ResourceNotFound _value, $Res Function(_$ResourceNotFound) _then)
      : super(_value, _then);
}

/// @nodoc

class _$ResourceNotFound
    with DiagnosticableTreeMixin
    implements ResourceNotFound {
  const _$ResourceNotFound();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.resourceNotFound()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
        .add(DiagnosticsProperty('type', 'BackendErrors.resourceNotFound'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$ResourceNotFound);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return resourceNotFound();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return resourceNotFound?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (resourceNotFound != null) {
      return resourceNotFound();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return resourceNotFound(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return resourceNotFound?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (resourceNotFound != null) {
      return resourceNotFound(this);
    }
    return orElse();
  }
}

abstract class ResourceNotFound implements BackendErrors {
  const factory ResourceNotFound() = _$ResourceNotFound;
}

/// @nodoc
abstract class _$$ServiceUnavailableCopyWith<$Res> {
  factory _$$ServiceUnavailableCopyWith(_$ServiceUnavailable value,
          $Res Function(_$ServiceUnavailable) then) =
      __$$ServiceUnavailableCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ServiceUnavailableCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$ServiceUnavailable>
    implements _$$ServiceUnavailableCopyWith<$Res> {
  __$$ServiceUnavailableCopyWithImpl(
      _$ServiceUnavailable _value, $Res Function(_$ServiceUnavailable) _then)
      : super(_value, _then);
}

/// @nodoc

class _$ServiceUnavailable
    with DiagnosticableTreeMixin
    implements ServiceUnavailable {
  const _$ServiceUnavailable();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.serviceUnavailable()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
        .add(DiagnosticsProperty('type', 'BackendErrors.serviceUnavailable'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$ServiceUnavailable);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return serviceUnavailable();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return serviceUnavailable?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (serviceUnavailable != null) {
      return serviceUnavailable();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return serviceUnavailable(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return serviceUnavailable?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (serviceUnavailable != null) {
      return serviceUnavailable(this);
    }
    return orElse();
  }
}

abstract class ServiceUnavailable implements BackendErrors {
  const factory ServiceUnavailable() = _$ServiceUnavailable;
}

/// @nodoc
abstract class _$$NoInternetConnectionCopyWith<$Res> {
  factory _$$NoInternetConnectionCopyWith(_$NoInternetConnection value,
          $Res Function(_$NoInternetConnection) then) =
      __$$NoInternetConnectionCopyWithImpl<$Res>;
}

/// @nodoc
class __$$NoInternetConnectionCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$NoInternetConnection>
    implements _$$NoInternetConnectionCopyWith<$Res> {
  __$$NoInternetConnectionCopyWithImpl(_$NoInternetConnection _value,
      $Res Function(_$NoInternetConnection) _then)
      : super(_value, _then);
}

/// @nodoc

class _$NoInternetConnection
    with DiagnosticableTreeMixin
    implements NoInternetConnection {
  const _$NoInternetConnection();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.noInternetConnection()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
        .add(DiagnosticsProperty('type', 'BackendErrors.noInternetConnection'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$NoInternetConnection);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return noInternetConnection();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return noInternetConnection?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (noInternetConnection != null) {
      return noInternetConnection();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return noInternetConnection(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return noInternetConnection?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (noInternetConnection != null) {
      return noInternetConnection(this);
    }
    return orElse();
  }
}

abstract class NoInternetConnection implements BackendErrors {
  const factory NoInternetConnection() = _$NoInternetConnection;
}

/// @nodoc
abstract class _$$BadResponseFormatCopyWith<$Res> {
  factory _$$BadResponseFormatCopyWith(
          _$BadResponseFormat value, $Res Function(_$BadResponseFormat) then) =
      __$$BadResponseFormatCopyWithImpl<$Res>;
}

/// @nodoc
class __$$BadResponseFormatCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$BadResponseFormat>
    implements _$$BadResponseFormatCopyWith<$Res> {
  __$$BadResponseFormatCopyWithImpl(
      _$BadResponseFormat _value, $Res Function(_$BadResponseFormat) _then)
      : super(_value, _then);
}

/// @nodoc

class _$BadResponseFormat
    with DiagnosticableTreeMixin
    implements BadResponseFormat {
  const _$BadResponseFormat();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.badResponseFormat()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
        .add(DiagnosticsProperty('type', 'BackendErrors.badResponseFormat'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$BadResponseFormat);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return badResponseFormat();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return badResponseFormat?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (badResponseFormat != null) {
      return badResponseFormat();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return badResponseFormat(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return badResponseFormat?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (badResponseFormat != null) {
      return badResponseFormat(this);
    }
    return orElse();
  }
}

abstract class BadResponseFormat implements BackendErrors {
  const factory BadResponseFormat() = _$BadResponseFormat;
}

/// @nodoc
abstract class _$$UnknownErrorCopyWith<$Res> {
  factory _$$UnknownErrorCopyWith(
          _$UnknownError value, $Res Function(_$UnknownError) then) =
      __$$UnknownErrorCopyWithImpl<$Res>;
}

/// @nodoc
class __$$UnknownErrorCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$UnknownError>
    implements _$$UnknownErrorCopyWith<$Res> {
  __$$UnknownErrorCopyWithImpl(
      _$UnknownError _value, $Res Function(_$UnknownError) _then)
      : super(_value, _then);
}

/// @nodoc

class _$UnknownError with DiagnosticableTreeMixin implements UnknownError {
  const _$UnknownError();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.unknownError()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'BackendErrors.unknownError'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$UnknownError);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
  }) {
    return unknownError();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? internalError,
    TResult? Function()? resourceNotFound,
    TResult? Function()? serviceUnavailable,
    TResult? Function()? noInternetConnection,
    TResult? Function()? badResponseFormat,
    TResult? Function()? unknownError,
  }) {
    return unknownError?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? internalError,
    TResult Function()? resourceNotFound,
    TResult Function()? serviceUnavailable,
    TResult Function()? noInternetConnection,
    TResult Function()? badResponseFormat,
    TResult Function()? unknownError,
    required TResult orElse(),
  }) {
    if (unknownError != null) {
      return unknownError();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(InternalError value) internalError,
    required TResult Function(ResourceNotFound value) resourceNotFound,
    required TResult Function(ServiceUnavailable value) serviceUnavailable,
    required TResult Function(NoInternetConnection value) noInternetConnection,
    required TResult Function(BadResponseFormat value) badResponseFormat,
    required TResult Function(UnknownError value) unknownError,
  }) {
    return unknownError(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(InternalError value)? internalError,
    TResult? Function(ResourceNotFound value)? resourceNotFound,
    TResult? Function(ServiceUnavailable value)? serviceUnavailable,
    TResult? Function(NoInternetConnection value)? noInternetConnection,
    TResult? Function(BadResponseFormat value)? badResponseFormat,
    TResult? Function(UnknownError value)? unknownError,
  }) {
    return unknownError?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(InternalError value)? internalError,
    TResult Function(ResourceNotFound value)? resourceNotFound,
    TResult Function(ServiceUnavailable value)? serviceUnavailable,
    TResult Function(NoInternetConnection value)? noInternetConnection,
    TResult Function(BadResponseFormat value)? badResponseFormat,
    TResult Function(UnknownError value)? unknownError,
    required TResult orElse(),
  }) {
    if (unknownError != null) {
      return unknownError(this);
    }
    return orElse();
  }
}

abstract class UnknownError implements BackendErrors {
  const factory UnknownError() = _$UnknownError;
}
