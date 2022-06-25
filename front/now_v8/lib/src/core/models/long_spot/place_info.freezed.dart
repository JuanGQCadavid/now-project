// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'place_info.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

PlaceInfo _$PlaceInfoFromJson(Map<String, dynamic> json) {
  return _PlaceInfo.fromJson(json);
}

/// @nodoc
mixin _$PlaceInfo {
  String get name => throw _privateConstructorUsedError;
  double get lat => throw _privateConstructorUsedError;
  double get lon => throw _privateConstructorUsedError;
  String get mapProviderId => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $PlaceInfoCopyWith<PlaceInfo> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $PlaceInfoCopyWith<$Res> {
  factory $PlaceInfoCopyWith(PlaceInfo value, $Res Function(PlaceInfo) then) =
      _$PlaceInfoCopyWithImpl<$Res>;
  $Res call({String name, double lat, double lon, String mapProviderId});
}

/// @nodoc
class _$PlaceInfoCopyWithImpl<$Res> implements $PlaceInfoCopyWith<$Res> {
  _$PlaceInfoCopyWithImpl(this._value, this._then);

  final PlaceInfo _value;
  // ignore: unused_field
  final $Res Function(PlaceInfo) _then;

  @override
  $Res call({
    Object? name = freezed,
    Object? lat = freezed,
    Object? lon = freezed,
    Object? mapProviderId = freezed,
  }) {
    return _then(_value.copyWith(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      lat: lat == freezed
          ? _value.lat
          : lat // ignore: cast_nullable_to_non_nullable
              as double,
      lon: lon == freezed
          ? _value.lon
          : lon // ignore: cast_nullable_to_non_nullable
              as double,
      mapProviderId: mapProviderId == freezed
          ? _value.mapProviderId
          : mapProviderId // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
abstract class _$$_PlaceInfoCopyWith<$Res> implements $PlaceInfoCopyWith<$Res> {
  factory _$$_PlaceInfoCopyWith(
          _$_PlaceInfo value, $Res Function(_$_PlaceInfo) then) =
      __$$_PlaceInfoCopyWithImpl<$Res>;
  @override
  $Res call({String name, double lat, double lon, String mapProviderId});
}

/// @nodoc
class __$$_PlaceInfoCopyWithImpl<$Res> extends _$PlaceInfoCopyWithImpl<$Res>
    implements _$$_PlaceInfoCopyWith<$Res> {
  __$$_PlaceInfoCopyWithImpl(
      _$_PlaceInfo _value, $Res Function(_$_PlaceInfo) _then)
      : super(_value, (v) => _then(v as _$_PlaceInfo));

  @override
  _$_PlaceInfo get _value => super._value as _$_PlaceInfo;

  @override
  $Res call({
    Object? name = freezed,
    Object? lat = freezed,
    Object? lon = freezed,
    Object? mapProviderId = freezed,
  }) {
    return _then(_$_PlaceInfo(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      lat: lat == freezed
          ? _value.lat
          : lat // ignore: cast_nullable_to_non_nullable
              as double,
      lon: lon == freezed
          ? _value.lon
          : lon // ignore: cast_nullable_to_non_nullable
              as double,
      mapProviderId: mapProviderId == freezed
          ? _value.mapProviderId
          : mapProviderId // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_PlaceInfo implements _PlaceInfo {
  const _$_PlaceInfo(
      {required this.name,
      required this.lat,
      required this.lon,
      required this.mapProviderId});

  factory _$_PlaceInfo.fromJson(Map<String, dynamic> json) =>
      _$$_PlaceInfoFromJson(json);

  @override
  final String name;
  @override
  final double lat;
  @override
  final double lon;
  @override
  final String mapProviderId;

  @override
  String toString() {
    return 'PlaceInfo(name: $name, lat: $lat, lon: $lon, mapProviderId: $mapProviderId)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_PlaceInfo &&
            const DeepCollectionEquality().equals(other.name, name) &&
            const DeepCollectionEquality().equals(other.lat, lat) &&
            const DeepCollectionEquality().equals(other.lon, lon) &&
            const DeepCollectionEquality()
                .equals(other.mapProviderId, mapProviderId));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(name),
      const DeepCollectionEquality().hash(lat),
      const DeepCollectionEquality().hash(lon),
      const DeepCollectionEquality().hash(mapProviderId));

  @JsonKey(ignore: true)
  @override
  _$$_PlaceInfoCopyWith<_$_PlaceInfo> get copyWith =>
      __$$_PlaceInfoCopyWithImpl<_$_PlaceInfo>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_PlaceInfoToJson(this);
  }
}

abstract class _PlaceInfo implements PlaceInfo {
  const factory _PlaceInfo(
      {required final String name,
      required final double lat,
      required final double lon,
      required final String mapProviderId}) = _$_PlaceInfo;

  factory _PlaceInfo.fromJson(Map<String, dynamic> json) =
      _$_PlaceInfo.fromJson;

  @override
  String get name => throw _privateConstructorUsedError;
  @override
  double get lat => throw _privateConstructorUsedError;
  @override
  double get lon => throw _privateConstructorUsedError;
  @override
  String get mapProviderId => throw _privateConstructorUsedError;
  @override
  @JsonKey(ignore: true)
  _$$_PlaceInfoCopyWith<_$_PlaceInfo> get copyWith =>
      throw _privateConstructorUsedError;
}
