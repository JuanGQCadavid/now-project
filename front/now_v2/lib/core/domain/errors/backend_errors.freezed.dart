// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'backend_errors.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more informations: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
class _$BackendErrorsTearOff {
  const _$BackendErrorsTearOff();

  NoConnection noConnection() {
    return const NoConnection();
  }

  InternalError internalError() {
    return const InternalError();
  }
}

/// @nodoc
const $BackendErrors = _$BackendErrorsTearOff();

/// @nodoc
mixin _$BackendErrors {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() noConnection,
    required TResult Function() internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(NoConnection value) noConnection,
    required TResult Function(InternalError value) internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $BackendErrorsCopyWith<$Res> {
  factory $BackendErrorsCopyWith(
          BackendErrors value, $Res Function(BackendErrors) then) =
      _$BackendErrorsCopyWithImpl<$Res>;
}

/// @nodoc
class _$BackendErrorsCopyWithImpl<$Res>
    implements $BackendErrorsCopyWith<$Res> {
  _$BackendErrorsCopyWithImpl(this._value, this._then);

  final BackendErrors _value;
  // ignore: unused_field
  final $Res Function(BackendErrors) _then;
}

/// @nodoc
abstract class $NoConnectionCopyWith<$Res> {
  factory $NoConnectionCopyWith(
          NoConnection value, $Res Function(NoConnection) then) =
      _$NoConnectionCopyWithImpl<$Res>;
}

/// @nodoc
class _$NoConnectionCopyWithImpl<$Res> extends _$BackendErrorsCopyWithImpl<$Res>
    implements $NoConnectionCopyWith<$Res> {
  _$NoConnectionCopyWithImpl(
      NoConnection _value, $Res Function(NoConnection) _then)
      : super(_value, (v) => _then(v as NoConnection));

  @override
  NoConnection get _value => super._value as NoConnection;
}

/// @nodoc

class _$NoConnection extends NoConnection {
  const _$NoConnection() : super._();

  @override
  String toString() {
    return 'BackendErrors.noConnection()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is NoConnection);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() noConnection,
    required TResult Function() internalError,
  }) {
    return noConnection();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
  }) {
    return noConnection?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
    required TResult orElse(),
  }) {
    if (noConnection != null) {
      return noConnection();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(NoConnection value) noConnection,
    required TResult Function(InternalError value) internalError,
  }) {
    return noConnection(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
  }) {
    return noConnection?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) {
    if (noConnection != null) {
      return noConnection(this);
    }
    return orElse();
  }
}

abstract class NoConnection extends BackendErrors {
  const factory NoConnection() = _$NoConnection;
  const NoConnection._() : super._();
}

/// @nodoc
abstract class $InternalErrorCopyWith<$Res> {
  factory $InternalErrorCopyWith(
          InternalError value, $Res Function(InternalError) then) =
      _$InternalErrorCopyWithImpl<$Res>;
}

/// @nodoc
class _$InternalErrorCopyWithImpl<$Res>
    extends _$BackendErrorsCopyWithImpl<$Res>
    implements $InternalErrorCopyWith<$Res> {
  _$InternalErrorCopyWithImpl(
      InternalError _value, $Res Function(InternalError) _then)
      : super(_value, (v) => _then(v as InternalError));

  @override
  InternalError get _value => super._value as InternalError;
}

/// @nodoc

class _$InternalError extends InternalError {
  const _$InternalError() : super._();

  @override
  String toString() {
    return 'BackendErrors.internalError()';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is InternalError);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() noConnection,
    required TResult Function() internalError,
  }) {
    return internalError();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
  }) {
    return internalError?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? noConnection,
    TResult Function()? internalError,
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
    required TResult Function(NoConnection value) noConnection,
    required TResult Function(InternalError value) internalError,
  }) {
    return internalError(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
  }) {
    return internalError?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(NoConnection value)? noConnection,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError(this);
    }
    return orElse();
  }
}

abstract class InternalError extends BackendErrors {
  const factory InternalError() = _$InternalError;
  const InternalError._() : super._();
}
