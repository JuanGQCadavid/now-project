// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'last_search_area.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$LastSearchArea {
  MapState get mapState => throw _privateConstructorUsedError;
  bool get jump => throw _privateConstructorUsedError;

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $LastSearchAreaCopyWith<LastSearchArea> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $LastSearchAreaCopyWith<$Res> {
  factory $LastSearchAreaCopyWith(
          LastSearchArea value, $Res Function(LastSearchArea) then) =
      _$LastSearchAreaCopyWithImpl<$Res, LastSearchArea>;
  @useResult
  $Res call({MapState mapState, bool jump});

  $MapStateCopyWith<$Res> get mapState;
}

/// @nodoc
class _$LastSearchAreaCopyWithImpl<$Res, $Val extends LastSearchArea>
    implements $LastSearchAreaCopyWith<$Res> {
  _$LastSearchAreaCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? mapState = null,
    Object? jump = null,
  }) {
    return _then(_value.copyWith(
      mapState: null == mapState
          ? _value.mapState
          : mapState // ignore: cast_nullable_to_non_nullable
              as MapState,
      jump: null == jump
          ? _value.jump
          : jump // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @override
  @pragma('vm:prefer-inline')
  $MapStateCopyWith<$Res> get mapState {
    return $MapStateCopyWith<$Res>(_value.mapState, (value) {
      return _then(_value.copyWith(mapState: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$MapStateImplCopyWith<$Res>
    implements $LastSearchAreaCopyWith<$Res> {
  factory _$$MapStateImplCopyWith(
          _$MapStateImpl value, $Res Function(_$MapStateImpl) then) =
      __$$MapStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({MapState mapState, bool jump});

  @override
  $MapStateCopyWith<$Res> get mapState;
}

/// @nodoc
class __$$MapStateImplCopyWithImpl<$Res>
    extends _$LastSearchAreaCopyWithImpl<$Res, _$MapStateImpl>
    implements _$$MapStateImplCopyWith<$Res> {
  __$$MapStateImplCopyWithImpl(
      _$MapStateImpl _value, $Res Function(_$MapStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? mapState = null,
    Object? jump = null,
  }) {
    return _then(_$MapStateImpl(
      mapState: null == mapState
          ? _value.mapState
          : mapState // ignore: cast_nullable_to_non_nullable
              as MapState,
      jump: null == jump
          ? _value.jump
          : jump // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc

class _$MapStateImpl implements _MapState {
  const _$MapStateImpl({required this.mapState, required this.jump});

  @override
  final MapState mapState;
  @override
  final bool jump;

  @override
  String toString() {
    return 'LastSearchArea(mapState: $mapState, jump: $jump)';
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$MapStateImpl &&
            (identical(other.mapState, mapState) ||
                other.mapState == mapState) &&
            (identical(other.jump, jump) || other.jump == jump));
  }

  @override
  int get hashCode => Object.hash(runtimeType, mapState, jump);

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$MapStateImplCopyWith<_$MapStateImpl> get copyWith =>
      __$$MapStateImplCopyWithImpl<_$MapStateImpl>(this, _$identity);
}

abstract class _MapState implements LastSearchArea {
  const factory _MapState(
      {required final MapState mapState,
      required final bool jump}) = _$MapStateImpl;

  @override
  MapState get mapState;
  @override
  bool get jump;

  /// Create a copy of LastSearchArea
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$MapStateImplCopyWith<_$MapStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
