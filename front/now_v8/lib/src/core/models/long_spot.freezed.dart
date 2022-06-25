// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'long_spot.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

LongSpot _$LongSpotFromJson(Map<String, dynamic> json) {
  return _LongSpot.fromJson(json);
}

/// @nodoc
mixin _$LongSpot {
  EventInfo get eventInfo => throw _privateConstructorUsedError;
  HostInfo get hostInfo => throw _privateConstructorUsedError;
  PlaceInfo get placeInfo => throw _privateConstructorUsedError;
  TopicsInfo get topicInfo => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $LongSpotCopyWith<LongSpot> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $LongSpotCopyWith<$Res> {
  factory $LongSpotCopyWith(LongSpot value, $Res Function(LongSpot) then) =
      _$LongSpotCopyWithImpl<$Res>;
  $Res call(
      {EventInfo eventInfo,
      HostInfo hostInfo,
      PlaceInfo placeInfo,
      TopicsInfo topicInfo});

  $EventInfoCopyWith<$Res> get eventInfo;
  $HostInfoCopyWith<$Res> get hostInfo;
  $PlaceInfoCopyWith<$Res> get placeInfo;
  $TopicsInfoCopyWith<$Res> get topicInfo;
}

/// @nodoc
class _$LongSpotCopyWithImpl<$Res> implements $LongSpotCopyWith<$Res> {
  _$LongSpotCopyWithImpl(this._value, this._then);

  final LongSpot _value;
  // ignore: unused_field
  final $Res Function(LongSpot) _then;

  @override
  $Res call({
    Object? eventInfo = freezed,
    Object? hostInfo = freezed,
    Object? placeInfo = freezed,
    Object? topicInfo = freezed,
  }) {
    return _then(_value.copyWith(
      eventInfo: eventInfo == freezed
          ? _value.eventInfo
          : eventInfo // ignore: cast_nullable_to_non_nullable
              as EventInfo,
      hostInfo: hostInfo == freezed
          ? _value.hostInfo
          : hostInfo // ignore: cast_nullable_to_non_nullable
              as HostInfo,
      placeInfo: placeInfo == freezed
          ? _value.placeInfo
          : placeInfo // ignore: cast_nullable_to_non_nullable
              as PlaceInfo,
      topicInfo: topicInfo == freezed
          ? _value.topicInfo
          : topicInfo // ignore: cast_nullable_to_non_nullable
              as TopicsInfo,
    ));
  }

  @override
  $EventInfoCopyWith<$Res> get eventInfo {
    return $EventInfoCopyWith<$Res>(_value.eventInfo, (value) {
      return _then(_value.copyWith(eventInfo: value));
    });
  }

  @override
  $HostInfoCopyWith<$Res> get hostInfo {
    return $HostInfoCopyWith<$Res>(_value.hostInfo, (value) {
      return _then(_value.copyWith(hostInfo: value));
    });
  }

  @override
  $PlaceInfoCopyWith<$Res> get placeInfo {
    return $PlaceInfoCopyWith<$Res>(_value.placeInfo, (value) {
      return _then(_value.copyWith(placeInfo: value));
    });
  }

  @override
  $TopicsInfoCopyWith<$Res> get topicInfo {
    return $TopicsInfoCopyWith<$Res>(_value.topicInfo, (value) {
      return _then(_value.copyWith(topicInfo: value));
    });
  }
}

/// @nodoc
abstract class _$$_LongSpotCopyWith<$Res> implements $LongSpotCopyWith<$Res> {
  factory _$$_LongSpotCopyWith(
          _$_LongSpot value, $Res Function(_$_LongSpot) then) =
      __$$_LongSpotCopyWithImpl<$Res>;
  @override
  $Res call(
      {EventInfo eventInfo,
      HostInfo hostInfo,
      PlaceInfo placeInfo,
      TopicsInfo topicInfo});

  @override
  $EventInfoCopyWith<$Res> get eventInfo;
  @override
  $HostInfoCopyWith<$Res> get hostInfo;
  @override
  $PlaceInfoCopyWith<$Res> get placeInfo;
  @override
  $TopicsInfoCopyWith<$Res> get topicInfo;
}

/// @nodoc
class __$$_LongSpotCopyWithImpl<$Res> extends _$LongSpotCopyWithImpl<$Res>
    implements _$$_LongSpotCopyWith<$Res> {
  __$$_LongSpotCopyWithImpl(
      _$_LongSpot _value, $Res Function(_$_LongSpot) _then)
      : super(_value, (v) => _then(v as _$_LongSpot));

  @override
  _$_LongSpot get _value => super._value as _$_LongSpot;

  @override
  $Res call({
    Object? eventInfo = freezed,
    Object? hostInfo = freezed,
    Object? placeInfo = freezed,
    Object? topicInfo = freezed,
  }) {
    return _then(_$_LongSpot(
      eventInfo: eventInfo == freezed
          ? _value.eventInfo
          : eventInfo // ignore: cast_nullable_to_non_nullable
              as EventInfo,
      hostInfo: hostInfo == freezed
          ? _value.hostInfo
          : hostInfo // ignore: cast_nullable_to_non_nullable
              as HostInfo,
      placeInfo: placeInfo == freezed
          ? _value.placeInfo
          : placeInfo // ignore: cast_nullable_to_non_nullable
              as PlaceInfo,
      topicInfo: topicInfo == freezed
          ? _value.topicInfo
          : topicInfo // ignore: cast_nullable_to_non_nullable
              as TopicsInfo,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_LongSpot implements _LongSpot {
  const _$_LongSpot(
      {required this.eventInfo,
      required this.hostInfo,
      required this.placeInfo,
      required this.topicInfo});

  factory _$_LongSpot.fromJson(Map<String, dynamic> json) =>
      _$$_LongSpotFromJson(json);

  @override
  final EventInfo eventInfo;
  @override
  final HostInfo hostInfo;
  @override
  final PlaceInfo placeInfo;
  @override
  final TopicsInfo topicInfo;

  @override
  String toString() {
    return 'LongSpot(eventInfo: $eventInfo, hostInfo: $hostInfo, placeInfo: $placeInfo, topicInfo: $topicInfo)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_LongSpot &&
            const DeepCollectionEquality().equals(other.eventInfo, eventInfo) &&
            const DeepCollectionEquality().equals(other.hostInfo, hostInfo) &&
            const DeepCollectionEquality().equals(other.placeInfo, placeInfo) &&
            const DeepCollectionEquality().equals(other.topicInfo, topicInfo));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(eventInfo),
      const DeepCollectionEquality().hash(hostInfo),
      const DeepCollectionEquality().hash(placeInfo),
      const DeepCollectionEquality().hash(topicInfo));

  @JsonKey(ignore: true)
  @override
  _$$_LongSpotCopyWith<_$_LongSpot> get copyWith =>
      __$$_LongSpotCopyWithImpl<_$_LongSpot>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_LongSpotToJson(this);
  }
}

abstract class _LongSpot implements LongSpot {
  const factory _LongSpot(
      {required final EventInfo eventInfo,
      required final HostInfo hostInfo,
      required final PlaceInfo placeInfo,
      required final TopicsInfo topicInfo}) = _$_LongSpot;

  factory _LongSpot.fromJson(Map<String, dynamic> json) = _$_LongSpot.fromJson;

  @override
  EventInfo get eventInfo => throw _privateConstructorUsedError;
  @override
  HostInfo get hostInfo => throw _privateConstructorUsedError;
  @override
  PlaceInfo get placeInfo => throw _privateConstructorUsedError;
  @override
  TopicsInfo get topicInfo => throw _privateConstructorUsedError;
  @override
  @JsonKey(ignore: true)
  _$$_LongSpotCopyWith<_$_LongSpot> get copyWith =>
      throw _privateConstructorUsedError;
}
