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
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

/// @nodoc
mixin _$SpotCreatorState {
  OnState get onState => throw _privateConstructorUsedError;
  int get totalSteps => throw _privateConstructorUsedError;
  int get actualStep => throw _privateConstructorUsedError;
  LongSpot get spot => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
  $SpotCreatorStateCopyWith<SpotCreatorState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SpotCreatorStateCopyWith<$Res> {
  factory $SpotCreatorStateCopyWith(
          SpotCreatorState value, $Res Function(SpotCreatorState) then) =
      _$SpotCreatorStateCopyWithImpl<$Res, SpotCreatorState>;
  @useResult
  $Res call({OnState onState, int totalSteps, int actualStep, LongSpot spot});

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

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? onState = null,
    Object? totalSteps = null,
    Object? actualStep = null,
    Object? spot = null,
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
    ) as $Val);
  }

  @override
  @pragma('vm:prefer-inline')
  $LongSpotCopyWith<$Res> get spot {
    return $LongSpotCopyWith<$Res>(_value.spot, (value) {
      return _then(_value.copyWith(spot: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$_SpotCreatorStateCopyWith<$Res>
    implements $SpotCreatorStateCopyWith<$Res> {
  factory _$$_SpotCreatorStateCopyWith(
          _$_SpotCreatorState value, $Res Function(_$_SpotCreatorState) then) =
      __$$_SpotCreatorStateCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({OnState onState, int totalSteps, int actualStep, LongSpot spot});

  @override
  $LongSpotCopyWith<$Res> get spot;
}

/// @nodoc
class __$$_SpotCreatorStateCopyWithImpl<$Res>
    extends _$SpotCreatorStateCopyWithImpl<$Res, _$_SpotCreatorState>
    implements _$$_SpotCreatorStateCopyWith<$Res> {
  __$$_SpotCreatorStateCopyWithImpl(
      _$_SpotCreatorState _value, $Res Function(_$_SpotCreatorState) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? onState = null,
    Object? totalSteps = null,
    Object? actualStep = null,
    Object? spot = null,
  }) {
    return _then(_$_SpotCreatorState(
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
    ));
  }
}

/// @nodoc

class _$_SpotCreatorState
    with DiagnosticableTreeMixin
    implements _SpotCreatorState {
  const _$_SpotCreatorState(
      {required this.onState,
      required this.totalSteps,
      required this.actualStep,
      required this.spot});

  @override
  final OnState onState;
  @override
  final int totalSteps;
  @override
  final int actualStep;
  @override
  final LongSpot spot;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'SpotCreatorState(onState: $onState, totalSteps: $totalSteps, actualStep: $actualStep, spot: $spot)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'SpotCreatorState'))
      ..add(DiagnosticsProperty('onState', onState))
      ..add(DiagnosticsProperty('totalSteps', totalSteps))
      ..add(DiagnosticsProperty('actualStep', actualStep))
      ..add(DiagnosticsProperty('spot', spot));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_SpotCreatorState &&
            (identical(other.onState, onState) || other.onState == onState) &&
            (identical(other.totalSteps, totalSteps) ||
                other.totalSteps == totalSteps) &&
            (identical(other.actualStep, actualStep) ||
                other.actualStep == actualStep) &&
            (identical(other.spot, spot) || other.spot == spot));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, onState, totalSteps, actualStep, spot);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_SpotCreatorStateCopyWith<_$_SpotCreatorState> get copyWith =>
      __$$_SpotCreatorStateCopyWithImpl<_$_SpotCreatorState>(this, _$identity);
}

abstract class _SpotCreatorState implements SpotCreatorState {
  const factory _SpotCreatorState(
      {required final OnState onState,
      required final int totalSteps,
      required final int actualStep,
      required final LongSpot spot}) = _$_SpotCreatorState;

  @override
  OnState get onState;
  @override
  int get totalSteps;
  @override
  int get actualStep;
  @override
  LongSpot get spot;
  @override
  @JsonKey(ignore: true)
  _$$_SpotCreatorStateCopyWith<_$_SpotCreatorState> get copyWith =>
      throw _privateConstructorUsedError;
}
