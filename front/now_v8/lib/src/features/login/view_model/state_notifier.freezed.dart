// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'state_notifier.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$UserError {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() userDoesNotExist,
    required TResult Function(String error) internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? userDoesNotExist,
    TResult? Function(String error)? internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? userDoesNotExist,
    TResult Function(String error)? internalError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(UserDoesNotExist value) userDoesNotExist,
    required TResult Function(InternalError value) internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(UserDoesNotExist value)? userDoesNotExist,
    TResult? Function(InternalError value)? internalError,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(UserDoesNotExist value)? userDoesNotExist,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $UserErrorCopyWith<$Res> {
  factory $UserErrorCopyWith(UserError value, $Res Function(UserError) then) =
      _$UserErrorCopyWithImpl<$Res, UserError>;
}

/// @nodoc
class _$UserErrorCopyWithImpl<$Res, $Val extends UserError>
    implements $UserErrorCopyWith<$Res> {
  _$UserErrorCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;
}

/// @nodoc
abstract class _$$UserDoesNotExistCopyWith<$Res> {
  factory _$$UserDoesNotExistCopyWith(
          _$UserDoesNotExist value, $Res Function(_$UserDoesNotExist) then) =
      __$$UserDoesNotExistCopyWithImpl<$Res>;
}

/// @nodoc
class __$$UserDoesNotExistCopyWithImpl<$Res>
    extends _$UserErrorCopyWithImpl<$Res, _$UserDoesNotExist>
    implements _$$UserDoesNotExistCopyWith<$Res> {
  __$$UserDoesNotExistCopyWithImpl(
      _$UserDoesNotExist _value, $Res Function(_$UserDoesNotExist) _then)
      : super(_value, _then);
}

/// @nodoc

class _$UserDoesNotExist
    with DiagnosticableTreeMixin
    implements UserDoesNotExist {
  _$UserDoesNotExist();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'UserError.userDoesNotExist()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'UserError.userDoesNotExist'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$UserDoesNotExist);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() userDoesNotExist,
    required TResult Function(String error) internalError,
  }) {
    return userDoesNotExist();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? userDoesNotExist,
    TResult? Function(String error)? internalError,
  }) {
    return userDoesNotExist?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? userDoesNotExist,
    TResult Function(String error)? internalError,
    required TResult orElse(),
  }) {
    if (userDoesNotExist != null) {
      return userDoesNotExist();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(UserDoesNotExist value) userDoesNotExist,
    required TResult Function(InternalError value) internalError,
  }) {
    return userDoesNotExist(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(UserDoesNotExist value)? userDoesNotExist,
    TResult? Function(InternalError value)? internalError,
  }) {
    return userDoesNotExist?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(UserDoesNotExist value)? userDoesNotExist,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) {
    if (userDoesNotExist != null) {
      return userDoesNotExist(this);
    }
    return orElse();
  }
}

abstract class UserDoesNotExist implements UserError {
  factory UserDoesNotExist() = _$UserDoesNotExist;
}

/// @nodoc
abstract class _$$InternalErrorCopyWith<$Res> {
  factory _$$InternalErrorCopyWith(
          _$InternalError value, $Res Function(_$InternalError) then) =
      __$$InternalErrorCopyWithImpl<$Res>;
  @useResult
  $Res call({String error});
}

/// @nodoc
class __$$InternalErrorCopyWithImpl<$Res>
    extends _$UserErrorCopyWithImpl<$Res, _$InternalError>
    implements _$$InternalErrorCopyWith<$Res> {
  __$$InternalErrorCopyWithImpl(
      _$InternalError _value, $Res Function(_$InternalError) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? error = null,
  }) {
    return _then(_$InternalError(
      null == error
          ? _value.error
          : error // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc

class _$InternalError with DiagnosticableTreeMixin implements InternalError {
  _$InternalError(this.error);

  @override
  final String error;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'UserError.internalError(error: $error)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'UserError.internalError'))
      ..add(DiagnosticsProperty('error', error));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$InternalError &&
            (identical(other.error, error) || other.error == error));
  }

  @override
  int get hashCode => Object.hash(runtimeType, error);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$InternalErrorCopyWith<_$InternalError> get copyWith =>
      __$$InternalErrorCopyWithImpl<_$InternalError>(this, _$identity);

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() userDoesNotExist,
    required TResult Function(String error) internalError,
  }) {
    return internalError(error);
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? userDoesNotExist,
    TResult? Function(String error)? internalError,
  }) {
    return internalError?.call(error);
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? userDoesNotExist,
    TResult Function(String error)? internalError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError(error);
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(UserDoesNotExist value) userDoesNotExist,
    required TResult Function(InternalError value) internalError,
  }) {
    return internalError(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(UserDoesNotExist value)? userDoesNotExist,
    TResult? Function(InternalError value)? internalError,
  }) {
    return internalError?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(UserDoesNotExist value)? userDoesNotExist,
    TResult Function(InternalError value)? internalError,
    required TResult orElse(),
  }) {
    if (internalError != null) {
      return internalError(this);
    }
    return orElse();
  }
}

abstract class InternalError implements UserError {
  factory InternalError(final String error) = _$InternalError;

  String get error;
  @JsonKey(ignore: true)
  _$$InternalErrorCopyWith<_$InternalError> get copyWith =>
      throw _privateConstructorUsedError;
}
