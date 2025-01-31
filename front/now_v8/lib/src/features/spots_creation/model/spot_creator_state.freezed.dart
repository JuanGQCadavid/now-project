// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'spot_creator_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$SpotCreatorState {
  OnState get onState => throw _privateConstructorUsedError;
  int get totalSteps => throw _privateConstructorUsedError;
  int get actualStep => throw _privateConstructorUsedError;
  LongSpot get spot => throw _privateConstructorUsedError;
  String get onError => throw _privateConstructorUsedError;

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $SpotCreatorStateCopyWith<SpotCreatorState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SpotCreatorStateCopyWith<$Res> {
  factory $SpotCreatorStateCopyWith(
          SpotCreatorState value, $Res Function(SpotCreatorState) then) =
      _$SpotCreatorStateCopyWithImpl<$Res, SpotCreatorState>;
  @useResult
  $Res call(
      {OnState onState,
      int totalSteps,
      int actualStep,
      LongSpot spot,
      String onError});

  $LongSpotCopyWith<$Res> get spot;
}

/// @nodoc
class _$SpotCreatorStateCopyWithImpl<$Res, $Val extends SpotCreatorState>
    implements $SpotCreatorStateCopyWith<$Res> {
  _$SpotCreatorStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? onState = null,
    Object? totalSteps = null,
    Object? actualStep = null,
    Object? spot = null,
    Object? onError = null,
  }) {
    return _then(_value.copyWith(
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as OnState,
      totalSteps: null == totalSteps
          ? _value.totalSteps
          : totalSteps // ignore: cast_nullable_to_non_nullable
              as int,
      actualStep: null == actualStep
          ? _value.actualStep
          : actualStep // ignore: cast_nullable_to_non_nullable
              as int,
      spot: null == spot
          ? _value.spot
          : spot // ignore: cast_nullable_to_non_nullable
              as LongSpot,
      onError: null == onError
          ? _value.onError
          : onError // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @pragma('vm:prefer-inline')
  $LongSpotCopyWith<$Res> get spot {
    return $LongSpotCopyWith<$Res>(_value.spot, (value) {
      return _then(_value.copyWith(spot: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$SpotCreatorStateImplCopyWith<$Res>
    implements $SpotCreatorStateCopyWith<$Res> {
  factory _$$SpotCreatorStateImplCopyWith(_$SpotCreatorStateImpl value,
          $Res Function(_$SpotCreatorStateImpl) then) =
      __$$SpotCreatorStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {OnState onState,
      int totalSteps,
      int actualStep,
      LongSpot spot,
      String onError});

  @override
  $LongSpotCopyWith<$Res> get spot;
}

/// @nodoc
class __$$SpotCreatorStateImplCopyWithImpl<$Res>
    extends _$SpotCreatorStateCopyWithImpl<$Res, _$SpotCreatorStateImpl>
    implements _$$SpotCreatorStateImplCopyWith<$Res> {
  __$$SpotCreatorStateImplCopyWithImpl(_$SpotCreatorStateImpl _value,
      $Res Function(_$SpotCreatorStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? onState = null,
    Object? totalSteps = null,
    Object? actualStep = null,
    Object? spot = null,
    Object? onError = null,
  }) {
    return _then(_$SpotCreatorStateImpl(
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as OnState,
      totalSteps: null == totalSteps
          ? _value.totalSteps
          : totalSteps // ignore: cast_nullable_to_non_nullable
              as int,
      actualStep: null == actualStep
          ? _value.actualStep
          : actualStep // ignore: cast_nullable_to_non_nullable
              as int,
      spot: null == spot
          ? _value.spot
          : spot // ignore: cast_nullable_to_non_nullable
              as LongSpot,
      onError: null == onError
          ? _value.onError
          : onError // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc

class _$SpotCreatorStateImpl
    with DiagnosticableTreeMixin
    implements _SpotCreatorState {
  const _$SpotCreatorStateImpl(
      {required this.onState,
      required this.totalSteps,
      required this.actualStep,
      required this.spot,
      required this.onError});

  @override
  final OnState onState;
  @override
  final int totalSteps;
  @override
  final int actualStep;
  @override
  final LongSpot spot;
  @override
  final String onError;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'SpotCreatorState(onState: $onState, totalSteps: $totalSteps, actualStep: $actualStep, spot: $spot, onError: $onError)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'SpotCreatorState'))
      ..add(DiagnosticsProperty('onState', onState))
      ..add(DiagnosticsProperty('totalSteps', totalSteps))
      ..add(DiagnosticsProperty('actualStep', actualStep))
      ..add(DiagnosticsProperty('spot', spot))
      ..add(DiagnosticsProperty('onError', onError));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$SpotCreatorStateImpl &&
            (identical(other.onState, onState) || other.onState == onState) &&
            (identical(other.totalSteps, totalSteps) ||
                other.totalSteps == totalSteps) &&
            (identical(other.actualStep, actualStep) ||
                other.actualStep == actualStep) &&
            (identical(other.spot, spot) || other.spot == spot) &&
            (identical(other.onError, onError) || other.onError == onError));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, onState, totalSteps, actualStep, spot, onError);

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$SpotCreatorStateImplCopyWith<_$SpotCreatorStateImpl> get copyWith =>
      __$$SpotCreatorStateImplCopyWithImpl<_$SpotCreatorStateImpl>(
          this, _$identity);
}

abstract class _SpotCreatorState implements SpotCreatorState {
  const factory _SpotCreatorState(
      {required final OnState onState,
      required final int totalSteps,
      required final int actualStep,
      required final LongSpot spot,
      required final String onError}) = _$SpotCreatorStateImpl;

  @override
  OnState get onState;
  @override
  int get totalSteps;
  @override
  int get actualStep;
  @override
  LongSpot get spot;
  @override
  String get onError;

  /// Create a copy of SpotCreatorState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$SpotCreatorStateImplCopyWith<_$SpotCreatorStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
