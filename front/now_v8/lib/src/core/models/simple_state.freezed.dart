// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'simple_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$SimpleState<T> {
  T get value => throw _privateConstructorUsedError;
  SimpleOnState get onState => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $SimpleStateCopyWith<T, SimpleState<T>> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SimpleStateCopyWith<T, $Res> {
  factory $SimpleStateCopyWith(
          SimpleState<T> value, $Res Function(SimpleState<T>) then) =
      _$SimpleStateCopyWithImpl<T, $Res, SimpleState<T>>;
  @useResult
  $Res call({T value, SimpleOnState onState});
}

/// @nodoc
class _$SimpleStateCopyWithImpl<T, $Res, $Val extends SimpleState<T>>
    implements $SimpleStateCopyWith<T, $Res> {
  _$SimpleStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? value = freezed,
    Object? onState = null,
  }) {
    return _then(_value.copyWith(
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as T,
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as SimpleOnState,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$SimpleStateImplCopyWith<T, $Res>
    implements $SimpleStateCopyWith<T, $Res> {
  factory _$$SimpleStateImplCopyWith(_$SimpleStateImpl<T> value,
          $Res Function(_$SimpleStateImpl<T>) then) =
      __$$SimpleStateImplCopyWithImpl<T, $Res>;
  @override
  @useResult
  $Res call({T value, SimpleOnState onState});
}

/// @nodoc
class __$$SimpleStateImplCopyWithImpl<T, $Res>
    extends _$SimpleStateCopyWithImpl<T, $Res, _$SimpleStateImpl<T>>
    implements _$$SimpleStateImplCopyWith<T, $Res> {
  __$$SimpleStateImplCopyWithImpl(
      _$SimpleStateImpl<T> _value, $Res Function(_$SimpleStateImpl<T>) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? value = freezed,
    Object? onState = null,
  }) {
    return _then(_$SimpleStateImpl<T>(
      value: freezed == value
          ? _value.value
          : value // ignore: cast_nullable_to_non_nullable
              as T,
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as SimpleOnState,
    ));
  }
}

/// @nodoc

class _$SimpleStateImpl<T>
    with DiagnosticableTreeMixin
    implements _SimpleState<T> {
  const _$SimpleStateImpl({required this.value, required this.onState});

  @override
  final T value;
  @override
  final SimpleOnState onState;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'SimpleState<$T>(value: $value, onState: $onState)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'SimpleState<$T>'))
      ..add(DiagnosticsProperty('value', value))
      ..add(DiagnosticsProperty('onState', onState));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$SimpleStateImpl<T> &&
            const DeepCollectionEquality().equals(other.value, value) &&
            (identical(other.onState, onState) || other.onState == onState));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType, const DeepCollectionEquality().hash(value), onState);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$SimpleStateImplCopyWith<T, _$SimpleStateImpl<T>> get copyWith =>
      __$$SimpleStateImplCopyWithImpl<T, _$SimpleStateImpl<T>>(
          this, _$identity);
}

abstract class _SimpleState<T> implements SimpleState<T> {
  const factory _SimpleState(
      {required final T value,
      required final SimpleOnState onState}) = _$SimpleStateImpl<T>;

  @override
  T get value;
  @override
  SimpleOnState get onState;
  @override
  @JsonKey(ignore: true)
  _$$SimpleStateImplCopyWith<T, _$SimpleStateImpl<T>> get copyWith =>
      throw _privateConstructorUsedError;
}
