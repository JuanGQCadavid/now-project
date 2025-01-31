// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'filteredSpots.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$MapStatus {
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() movingOnMap,
    required TResult Function() movingIdle,
    required TResult Function() movingStarted,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? movingOnMap,
    TResult? Function()? movingIdle,
    TResult? Function()? movingStarted,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? movingOnMap,
    TResult Function()? movingIdle,
    TResult Function()? movingStarted,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(MovingOnMap value) movingOnMap,
    required TResult Function(MovingIdle value) movingIdle,
    required TResult Function(MovingStarted value) movingStarted,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(MovingOnMap value)? movingOnMap,
    TResult? Function(MovingIdle value)? movingIdle,
    TResult? Function(MovingStarted value)? movingStarted,
  }) =>
      throw _privateConstructorUsedError;
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(MovingOnMap value)? movingOnMap,
    TResult Function(MovingIdle value)? movingIdle,
    TResult Function(MovingStarted value)? movingStarted,
    required TResult orElse(),
  }) =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $MapStatusCopyWith<$Res> {
  factory $MapStatusCopyWith(MapStatus value, $Res Function(MapStatus) then) =
      _$MapStatusCopyWithImpl<$Res, MapStatus>;
}

/// @nodoc
class _$MapStatusCopyWithImpl<$Res, $Val extends MapStatus>
    implements $MapStatusCopyWith<$Res> {
  _$MapStatusCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of MapStatus
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc
abstract class _$$MovingOnMapImplCopyWith<$Res> {
  factory _$$MovingOnMapImplCopyWith(
          _$MovingOnMapImpl value, $Res Function(_$MovingOnMapImpl) then) =
      __$$MovingOnMapImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingOnMapImplCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingOnMapImpl>
    implements _$$MovingOnMapImplCopyWith<$Res> {
  __$$MovingOnMapImplCopyWithImpl(
      _$MovingOnMapImpl _value, $Res Function(_$MovingOnMapImpl) _then)
      : super(_value, _then);

  /// Create a copy of MapStatus
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$MovingOnMapImpl with DiagnosticableTreeMixin implements MovingOnMap {
  _$MovingOnMapImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingOnMap()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'MapStatus.movingOnMap'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingOnMapImpl);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() movingOnMap,
    required TResult Function() movingIdle,
    required TResult Function() movingStarted,
  }) {
    return movingOnMap();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? movingOnMap,
    TResult? Function()? movingIdle,
    TResult? Function()? movingStarted,
  }) {
    return movingOnMap?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? movingOnMap,
    TResult Function()? movingIdle,
    TResult Function()? movingStarted,
    required TResult orElse(),
  }) {
    if (movingOnMap != null) {
      return movingOnMap();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(MovingOnMap value) movingOnMap,
    required TResult Function(MovingIdle value) movingIdle,
    required TResult Function(MovingStarted value) movingStarted,
  }) {
    return movingOnMap(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(MovingOnMap value)? movingOnMap,
    TResult? Function(MovingIdle value)? movingIdle,
    TResult? Function(MovingStarted value)? movingStarted,
  }) {
    return movingOnMap?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(MovingOnMap value)? movingOnMap,
    TResult Function(MovingIdle value)? movingIdle,
    TResult Function(MovingStarted value)? movingStarted,
    required TResult orElse(),
  }) {
    if (movingOnMap != null) {
      return movingOnMap(this);
    }
    return orElse();
  }
}

abstract class MovingOnMap implements MapStatus {
  factory MovingOnMap() = _$MovingOnMapImpl;
}

/// @nodoc
abstract class _$$MovingIdleImplCopyWith<$Res> {
  factory _$$MovingIdleImplCopyWith(
          _$MovingIdleImpl value, $Res Function(_$MovingIdleImpl) then) =
      __$$MovingIdleImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingIdleImplCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingIdleImpl>
    implements _$$MovingIdleImplCopyWith<$Res> {
  __$$MovingIdleImplCopyWithImpl(
      _$MovingIdleImpl _value, $Res Function(_$MovingIdleImpl) _then)
      : super(_value, _then);

  /// Create a copy of MapStatus
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$MovingIdleImpl with DiagnosticableTreeMixin implements MovingIdle {
  _$MovingIdleImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingIdle()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'MapStatus.movingIdle'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingIdleImpl);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() movingOnMap,
    required TResult Function() movingIdle,
    required TResult Function() movingStarted,
  }) {
    return movingIdle();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? movingOnMap,
    TResult? Function()? movingIdle,
    TResult? Function()? movingStarted,
  }) {
    return movingIdle?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? movingOnMap,
    TResult Function()? movingIdle,
    TResult Function()? movingStarted,
    required TResult orElse(),
  }) {
    if (movingIdle != null) {
      return movingIdle();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(MovingOnMap value) movingOnMap,
    required TResult Function(MovingIdle value) movingIdle,
    required TResult Function(MovingStarted value) movingStarted,
  }) {
    return movingIdle(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(MovingOnMap value)? movingOnMap,
    TResult? Function(MovingIdle value)? movingIdle,
    TResult? Function(MovingStarted value)? movingStarted,
  }) {
    return movingIdle?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(MovingOnMap value)? movingOnMap,
    TResult Function(MovingIdle value)? movingIdle,
    TResult Function(MovingStarted value)? movingStarted,
    required TResult orElse(),
  }) {
    if (movingIdle != null) {
      return movingIdle(this);
    }
    return orElse();
  }
}

abstract class MovingIdle implements MapStatus {
  factory MovingIdle() = _$MovingIdleImpl;
}

/// @nodoc
abstract class _$$MovingStartedImplCopyWith<$Res> {
  factory _$$MovingStartedImplCopyWith(
          _$MovingStartedImpl value, $Res Function(_$MovingStartedImpl) then) =
      __$$MovingStartedImplCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingStartedImplCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingStartedImpl>
    implements _$$MovingStartedImplCopyWith<$Res> {
  __$$MovingStartedImplCopyWithImpl(
      _$MovingStartedImpl _value, $Res Function(_$MovingStartedImpl) _then)
      : super(_value, _then);

  /// Create a copy of MapStatus
  /// with the given fields replaced by the non-null parameter values.
}

/// @nodoc

class _$MovingStartedImpl
    with DiagnosticableTreeMixin
    implements MovingStarted {
  _$MovingStartedImpl();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingStarted()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties..add(DiagnosticsProperty('type', 'MapStatus.movingStarted'));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingStartedImpl);
  }

  @override
  int get hashCode => runtimeType.hashCode;

  @override
  @optionalTypeArgs
  TResult when<TResult extends Object?>({
    required TResult Function() movingOnMap,
    required TResult Function() movingIdle,
    required TResult Function() movingStarted,
  }) {
    return movingStarted();
  }

  @override
  @optionalTypeArgs
  TResult? whenOrNull<TResult extends Object?>({
    TResult? Function()? movingOnMap,
    TResult? Function()? movingIdle,
    TResult? Function()? movingStarted,
  }) {
    return movingStarted?.call();
  }

  @override
  @optionalTypeArgs
  TResult maybeWhen<TResult extends Object?>({
    TResult Function()? movingOnMap,
    TResult Function()? movingIdle,
    TResult Function()? movingStarted,
    required TResult orElse(),
  }) {
    if (movingStarted != null) {
      return movingStarted();
    }
    return orElse();
  }

  @override
  @optionalTypeArgs
  TResult map<TResult extends Object?>({
    required TResult Function(MovingOnMap value) movingOnMap,
    required TResult Function(MovingIdle value) movingIdle,
    required TResult Function(MovingStarted value) movingStarted,
  }) {
    return movingStarted(this);
  }

  @override
  @optionalTypeArgs
  TResult? mapOrNull<TResult extends Object?>({
    TResult? Function(MovingOnMap value)? movingOnMap,
    TResult? Function(MovingIdle value)? movingIdle,
    TResult? Function(MovingStarted value)? movingStarted,
  }) {
    return movingStarted?.call(this);
  }

  @override
  @optionalTypeArgs
  TResult maybeMap<TResult extends Object?>({
    TResult Function(MovingOnMap value)? movingOnMap,
    TResult Function(MovingIdle value)? movingIdle,
    TResult Function(MovingStarted value)? movingStarted,
    required TResult orElse(),
  }) {
    if (movingStarted != null) {
      return movingStarted(this);
    }
    return orElse();
  }
}

abstract class MovingStarted implements MapStatus {
  factory MovingStarted() = _$MovingStartedImpl;
}

/// @nodoc
mixin _$MapState {
  LatLng get lastPositionKnowed => throw _privateConstructorUsedError;
  double get zoom => throw _privateConstructorUsedError;
  MapStatus get status => throw _privateConstructorUsedError;

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $MapStateCopyWith<MapState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $MapStateCopyWith<$Res> {
  factory $MapStateCopyWith(MapState value, $Res Function(MapState) then) =
      _$MapStateCopyWithImpl<$Res, MapState>;
  @useResult
  $Res call({LatLng lastPositionKnowed, double zoom, MapStatus status});

  $MapStatusCopyWith<$Res> get status;
}

/// @nodoc
class _$MapStateCopyWithImpl<$Res, $Val extends MapState>
    implements $MapStateCopyWith<$Res> {
  _$MapStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? lastPositionKnowed = null,
    Object? zoom = null,
    Object? status = null,
  }) {
    return _then(_value.copyWith(
      lastPositionKnowed: null == lastPositionKnowed
          ? _value.lastPositionKnowed
          : lastPositionKnowed // ignore: cast_nullable_to_non_nullable
              as LatLng,
      zoom: null == zoom
          ? _value.zoom
          : zoom // ignore: cast_nullable_to_non_nullable
              as double,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as MapStatus,
    ) as $Val);
  }

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @pragma('vm:prefer-inline')
  $MapStatusCopyWith<$Res> get status {
    return $MapStatusCopyWith<$Res>(_value.status, (value) {
      return _then(_value.copyWith(status: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$MapStateImplCopyWith<$Res>
    implements $MapStateCopyWith<$Res> {
  factory _$$MapStateImplCopyWith(
          _$MapStateImpl value, $Res Function(_$MapStateImpl) then) =
      __$$MapStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({LatLng lastPositionKnowed, double zoom, MapStatus status});

  @override
  $MapStatusCopyWith<$Res> get status;
}

/// @nodoc
class __$$MapStateImplCopyWithImpl<$Res>
    extends _$MapStateCopyWithImpl<$Res, _$MapStateImpl>
    implements _$$MapStateImplCopyWith<$Res> {
  __$$MapStateImplCopyWithImpl(
      _$MapStateImpl _value, $Res Function(_$MapStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? lastPositionKnowed = null,
    Object? zoom = null,
    Object? status = null,
  }) {
    return _then(_$MapStateImpl(
      lastPositionKnowed: null == lastPositionKnowed
          ? _value.lastPositionKnowed
          : lastPositionKnowed // ignore: cast_nullable_to_non_nullable
              as LatLng,
      zoom: null == zoom
          ? _value.zoom
          : zoom // ignore: cast_nullable_to_non_nullable
              as double,
      status: null == status
          ? _value.status
          : status // ignore: cast_nullable_to_non_nullable
              as MapStatus,
    ));
  }
}

/// @nodoc

class _$MapStateImpl with DiagnosticableTreeMixin implements _MapState {
  const _$MapStateImpl(
      {required this.lastPositionKnowed,
      required this.zoom,
      required this.status});

  @override
  final LatLng lastPositionKnowed;
  @override
  final double zoom;
  @override
  final MapStatus status;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapState(lastPositionKnowed: $lastPositionKnowed, zoom: $zoom, status: $status)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'MapState'))
      ..add(DiagnosticsProperty('lastPositionKnowed', lastPositionKnowed))
      ..add(DiagnosticsProperty('zoom', zoom))
      ..add(DiagnosticsProperty('status', status));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$MapStateImpl &&
            (identical(other.lastPositionKnowed, lastPositionKnowed) ||
                other.lastPositionKnowed == lastPositionKnowed) &&
            (identical(other.zoom, zoom) || other.zoom == zoom) &&
            (identical(other.status, status) || other.status == status));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, lastPositionKnowed, zoom, status);

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$MapStateImplCopyWith<_$MapStateImpl> get copyWith =>
      __$$MapStateImplCopyWithImpl<_$MapStateImpl>(this, _$identity);
}

abstract class _MapState implements MapState {
  const factory _MapState(
      {required final LatLng lastPositionKnowed,
      required final double zoom,
      required final MapStatus status}) = _$MapStateImpl;

  @override
  LatLng get lastPositionKnowed;
  @override
  double get zoom;
  @override
  MapStatus get status;

  /// Create a copy of MapState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$MapStateImplCopyWith<_$MapStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
