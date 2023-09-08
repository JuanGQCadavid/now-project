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
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

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
}

/// @nodoc
abstract class _$$MovingOnMapCopyWith<$Res> {
  factory _$$MovingOnMapCopyWith(
          _$MovingOnMap value, $Res Function(_$MovingOnMap) then) =
      __$$MovingOnMapCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingOnMapCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingOnMap>
    implements _$$MovingOnMapCopyWith<$Res> {
  __$$MovingOnMapCopyWithImpl(
      _$MovingOnMap _value, $Res Function(_$MovingOnMap) _then)
      : super(_value, _then);
}

/// @nodoc

class _$MovingOnMap with DiagnosticableTreeMixin implements MovingOnMap {
  _$MovingOnMap();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingOnMap()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'MapStatus.movingOnMap'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingOnMap);
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
  factory MovingOnMap() = _$MovingOnMap;
}

/// @nodoc
abstract class _$$MovingIdleCopyWith<$Res> {
  factory _$$MovingIdleCopyWith(
          _$MovingIdle value, $Res Function(_$MovingIdle) then) =
      __$$MovingIdleCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingIdleCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingIdle>
    implements _$$MovingIdleCopyWith<$Res> {
  __$$MovingIdleCopyWithImpl(
      _$MovingIdle _value, $Res Function(_$MovingIdle) _then)
      : super(_value, _then);
}

/// @nodoc

class _$MovingIdle with DiagnosticableTreeMixin implements MovingIdle {
  _$MovingIdle();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingIdle()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'MapStatus.movingIdle'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingIdle);
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
  factory MovingIdle() = _$MovingIdle;
}

/// @nodoc
abstract class _$$MovingStartedCopyWith<$Res> {
  factory _$$MovingStartedCopyWith(
          _$MovingStarted value, $Res Function(_$MovingStarted) then) =
      __$$MovingStartedCopyWithImpl<$Res>;
}

/// @nodoc
class __$$MovingStartedCopyWithImpl<$Res>
    extends _$MapStatusCopyWithImpl<$Res, _$MovingStarted>
    implements _$$MovingStartedCopyWith<$Res> {
  __$$MovingStartedCopyWithImpl(
      _$MovingStarted _value, $Res Function(_$MovingStarted) _then)
      : super(_value, _then);
}

/// @nodoc

class _$MovingStarted with DiagnosticableTreeMixin implements MovingStarted {
  _$MovingStarted();

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'MapStatus.movingStarted()';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty('type', 'MapStatus.movingStarted'));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType && other is _$MovingStarted);
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
  factory MovingStarted() = _$MovingStarted;
}

/// @nodoc
mixin _$MapState {
  LatLng get lastPositionKnowed => throw _privateConstructorUsedError;
  double get zoom => throw _privateConstructorUsedError;
  MapStatus get status => throw _privateConstructorUsedError;

  @JsonKey(ignore: true)
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

  @override
  @pragma('vm:prefer-inline')
  $MapStatusCopyWith<$Res> get status {
    return $MapStatusCopyWith<$Res>(_value.status, (value) {
      return _then(_value.copyWith(status: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$_MapStateCopyWith<$Res> implements $MapStateCopyWith<$Res> {
  factory _$$_MapStateCopyWith(
          _$_MapState value, $Res Function(_$_MapState) then) =
      __$$_MapStateCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({LatLng lastPositionKnowed, double zoom, MapStatus status});

  @override
  $MapStatusCopyWith<$Res> get status;
}

/// @nodoc
class __$$_MapStateCopyWithImpl<$Res>
    extends _$MapStateCopyWithImpl<$Res, _$_MapState>
    implements _$$_MapStateCopyWith<$Res> {
  __$$_MapStateCopyWithImpl(
      _$_MapState _value, $Res Function(_$_MapState) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? lastPositionKnowed = null,
    Object? zoom = null,
    Object? status = null,
  }) {
    return _then(_$_MapState(
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

class _$_MapState with DiagnosticableTreeMixin implements _MapState {
  const _$_MapState(
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
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_MapState &&
            (identical(other.lastPositionKnowed, lastPositionKnowed) ||
                other.lastPositionKnowed == lastPositionKnowed) &&
            (identical(other.zoom, zoom) || other.zoom == zoom) &&
            (identical(other.status, status) || other.status == status));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, lastPositionKnowed, zoom, status);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_MapStateCopyWith<_$_MapState> get copyWith =>
      __$$_MapStateCopyWithImpl<_$_MapState>(this, _$identity);
}

abstract class _MapState implements MapState {
  const factory _MapState(
      {required final LatLng lastPositionKnowed,
      required final double zoom,
      required final MapStatus status}) = _$_MapState;

  @override
  LatLng get lastPositionKnowed;
  @override
  double get zoom;
  @override
  MapStatus get status;
  @override
  @JsonKey(ignore: true)
  _$$_MapStateCopyWith<_$_MapState> get copyWith =>
      throw _privateConstructorUsedError;
}
