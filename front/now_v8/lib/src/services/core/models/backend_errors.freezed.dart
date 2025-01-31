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
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
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

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc
abstract class _$$InternalErrorImplCopyWith<$Res> {
  factory _$$InternalErrorImplCopyWith(
          _$InternalErrorImpl value, $Res Function(_$InternalErrorImpl) then) =
      __$$InternalErrorImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$InternalErrorImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$InternalErrorImpl>
    implements _$$InternalErrorImplCopyWith<$Res> {
  __$$InternalErrorImplCopyWithImpl(
      _$InternalErrorImpl _value, $Res Function(_$InternalErrorImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$InternalErrorImpl
    with DiagnosticableTreeMixin
    implements InternalError {
  const _$InternalErrorImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.internalError()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'BackendErrors.internalError'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$InternalErrorImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError(this);
    }
    return orElse();
  }
}

abstract class InternalError implements BackendErrors {
  const factory InternalError() = _$InternalErrorImpl;
}

/// @nodoc
abstract class _$$ResourceNotFoundImplCopyWith<$Res> {
  factory _$$ResourceNotFoundImplCopyWith(_$ResourceNotFoundImpl value,
          $Res Function(_$ResourceNotFoundImpl) then) =
      __$$ResourceNotFoundImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ResourceNotFoundImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$ResourceNotFoundImpl>
    implements _$$ResourceNotFoundImplCopyWith<$Res> {
  __$$ResourceNotFoundImplCopyWithImpl(_$ResourceNotFoundImpl _value,
      $Res Function(_$ResourceNotFoundImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$ResourceNotFoundImpl
    with DiagnosticableTreeMixin
    implements ResourceNotFound {
  const _$ResourceNotFoundImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.resourceNotFound()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'BackendErrors.resourceNotFound'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$ResourceNotFoundImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (resourceNotFound != null) {
      return resourceNotFound(this);
    }
    return orElse();
  }
}

abstract class ResourceNotFound implements BackendErrors {
  const factory ResourceNotFound() = _$ResourceNotFoundImpl;
}

/// @nodoc
abstract class _$$ServiceUnavailableImplCopyWith<$Res> {
  factory _$$ServiceUnavailableImplCopyWith(_$ServiceUnavailableImpl value,
          $Res Function(_$ServiceUnavailableImpl) then) =
      __$$ServiceUnavailableImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$ServiceUnavailableImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$ServiceUnavailableImpl>
    implements _$$ServiceUnavailableImplCopyWith<$Res> {
  __$$ServiceUnavailableImplCopyWithImpl(_$ServiceUnavailableImpl _value,
      $Res Function(_$ServiceUnavailableImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$ServiceUnavailableImpl
    with DiagnosticableTreeMixin
    implements ServiceUnavailable {
  const _$ServiceUnavailableImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.serviceUnavailable()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'BackendErrors.serviceUnavailable'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$ServiceUnavailableImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (serviceUnavailable != null) {
      return serviceUnavailable(this);
    }
    return orElse();
  }
}

abstract class ServiceUnavailable implements BackendErrors {
  const factory ServiceUnavailable() = _$ServiceUnavailableImpl;
}

/// @nodoc
abstract class _$$NoInternetConnectionImplCopyWith<$Res> {
  factory _$$NoInternetConnectionImplCopyWith(_$NoInternetConnectionImpl value,
          $Res Function(_$NoInternetConnectionImpl) then) =
      __$$NoInternetConnectionImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$NoInternetConnectionImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$NoInternetConnectionImpl>
    implements _$$NoInternetConnectionImplCopyWith<$Res> {
  __$$NoInternetConnectionImplCopyWithImpl(_$NoInternetConnectionImpl _value,
      $Res Function(_$NoInternetConnectionImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$NoInternetConnectionImpl
    with DiagnosticableTreeMixin
    implements NoInternetConnection {
  const _$NoInternetConnectionImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.noInternetConnection()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'BackendErrors.noInternetConnection'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$NoInternetConnectionImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (noInternetConnection != null) {
      return noInternetConnection(this);
    }
    return orElse();
  }
}

abstract class NoInternetConnection implements BackendErrors {
  const factory NoInternetConnection() = _$NoInternetConnectionImpl;
}

/// @nodoc
abstract class _$$BadResponseFormatImplCopyWith<$Res> {
  factory _$$BadResponseFormatImplCopyWith(_$BadResponseFormatImpl value,
          $Res Function(_$BadResponseFormatImpl) then) =
      __$$BadResponseFormatImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$BadResponseFormatImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$BadResponseFormatImpl>
    implements _$$BadResponseFormatImplCopyWith<$Res> {
  __$$BadResponseFormatImplCopyWithImpl(_$BadResponseFormatImpl _value,
      $Res Function(_$BadResponseFormatImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$BadResponseFormatImpl
    with DiagnosticableTreeMixin
    implements BadResponseFormat {
  const _$BadResponseFormatImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.badResponseFormat()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'BackendErrors.badResponseFormat'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$BadResponseFormatImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (badResponseFormat != null) {
      return badResponseFormat(this);
    }
    return orElse();
  }
}

abstract class BadResponseFormat implements BackendErrors {
  const factory BadResponseFormat() = _$BadResponseFormatImpl;
}

/// @nodoc
abstract class _$$UnknownErrorImplCopyWith<$Res> {
  factory _$$UnknownErrorImplCopyWith(
          _$UnknownErrorImpl value, $Res Function(_$UnknownErrorImpl) then) =
      __$$UnknownErrorImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$UnknownErrorImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$UnknownErrorImpl>
    implements _$$UnknownErrorImplCopyWith<$Res> {
  __$$UnknownErrorImplCopyWithImpl(
      _$UnknownErrorImpl _value, $Res Function(_$UnknownErrorImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$UnknownErrorImpl with DiagnosticableTreeMixin implements UnknownError {
  const _$UnknownErrorImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.unknownError()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'BackendErrors.unknownError'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$UnknownErrorImpl);
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
    required TResult Function(ErrorMessage errorMessage) clientError,
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
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
    TResult Function(ErrorMessage errorMessage)? clientError,
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
    required TResult Function(ClientError value) clientError,
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
    TResult? Function(ClientError value)? clientError,
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (unknownError != null) {
      return unknownError(this);
    }
    return orElse();
  }
}

abstract class UnknownError implements BackendErrors {
  const factory UnknownError() = _$UnknownErrorImpl;
}

/// @nodoc
abstract class _$$ClientErrorImplCopyWith<$Res> {
  factory _$$ClientErrorImplCopyWith(
          _$ClientErrorImpl value, $Res Function(_$ClientErrorImpl) then) =
      __$$ClientErrorImplCopyWithImpl<$Res>;
  @useResult
  $Res call({ErrorMessage errorMessage});
}

/// @nodoc
class __$$ClientErrorImplCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res, _$ClientErrorImpl>
    implements _$$ClientErrorImplCopyWith<$Res> {
  __$$ClientErrorImplCopyWithImpl(
      _$ClientErrorImpl _value, $Res Function(_$ClientErrorImpl) _then)
      : super(_value, _then);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? errorMessage = null,
  }) {
    return _then(_$ClientErrorImpl(
      null == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as ErrorMessage,
    ));
  }
}

/// @nodoc

class _$ClientErrorImpl with DiagnosticableTreeMixin implements ClientError {
  const _$ClientErrorImpl(this.errorMessage);

  @override
  final ErrorMessage errorMessage;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'BackendErrors.clientError(errorMessage: $errorMessage)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'BackendErrors.clientError'))
      ..add(DiagnosticsProperty('errorMessage', errorMessage));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$ClientErrorImpl &&
            (identical(other.errorMessage, errorMessage) ||
                other.errorMessage == errorMessage));
  }

  @override
  int get hashCode => Object.hash(runtimeType, errorMessage);

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$ClientErrorImplCopyWith<_$ClientErrorImpl> get copyWith =>
      __$$ClientErrorImplCopyWithImpl<_$ClientErrorImpl>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() internalError,
    required TResult Function() resourceNotFound,
    required TResult Function() serviceUnavailable,
    required TResult Function() noInternetConnection,
    required TResult Function() badResponseFormat,
    required TResult Function() unknownError,
    required TResult Function(ErrorMessage errorMessage) clientError,
  }) {
    return clientError(errorMessage);
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
    TResult? Function(ErrorMessage errorMessage)? clientError,
  }) {
    return clientError?.call(errorMessage);
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
    TResult Function(ErrorMessage errorMessage)? clientError,
    required TResult orElse(),
  }) {
    if (clientError != null) {
      return clientError(errorMessage);
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
    required TResult Function(ClientError value) clientError,
  }) {
    return clientError(this);
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
    TResult? Function(ClientError value)? clientError,
  }) {
    return clientError?.call(this);
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
    TResult Function(ClientError value)? clientError,
    required TResult orElse(),
  }) {
    if (clientError != null) {
      return clientError(this);
    }
    return orElse();
  }
}

abstract class ClientError implements BackendErrors {
  const factory ClientError(final ErrorMessage errorMessage) =
      _$ClientErrorImpl;

  ErrorMessage get errorMessage;

  /// Create a copy of BackendErrors
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$ClientErrorImplCopyWith<_$ClientErrorImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
