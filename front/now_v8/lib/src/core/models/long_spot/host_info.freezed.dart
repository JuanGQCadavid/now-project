// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'host_info.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

HostInfo _$HostInfoFromJson(Map<String, dynamic> json) {
  return _HostInfo.fromJson(json);
}

/// @nodoc
mixin _$HostInfo {
  String get name => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $HostInfoCopyWith<HostInfo> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $HostInfoCopyWith<$Res> {
  factory $HostInfoCopyWith(HostInfo value, $Res Function(HostInfo) then) =
      _$HostInfoCopyWithImpl<$Res>;
  $Res call({String name});
}

/// @nodoc
class _$HostInfoCopyWithImpl<$Res> implements $HostInfoCopyWith<$Res> {
  _$HostInfoCopyWithImpl(this._value, this._then);

  final HostInfo _value;
  // ignore: unused_field
  final $Res Function(HostInfo) _then;

  @override
  $Res call({
    Object? name = freezed,
  }) {
    return _then(_value.copyWith(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
abstract class _$$_HostInfoCopyWith<$Res> implements $HostInfoCopyWith<$Res> {
  factory _$$_HostInfoCopyWith(
          _$_HostInfo value, $Res Function(_$_HostInfo) then) =
      __$$_HostInfoCopyWithImpl<$Res>;
  @override
  $Res call({String name});
}

/// @nodoc
class __$$_HostInfoCopyWithImpl<$Res> extends _$HostInfoCopyWithImpl<$Res>
    implements _$$_HostInfoCopyWith<$Res> {
  __$$_HostInfoCopyWithImpl(
      _$_HostInfo _value, $Res Function(_$_HostInfo) _then)
      : super(_value, (v) => _then(v as _$_HostInfo));

  @override
  _$_HostInfo get _value => super._value as _$_HostInfo;

  @override
  $Res call({
    Object? name = freezed,
  }) {
    return _then(_$_HostInfo(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_HostInfo implements _HostInfo {
  const _$_HostInfo({required this.name});

  factory _$_HostInfo.fromJson(Map<String, dynamic> json) =>
      _$$_HostInfoFromJson(json);

  @override
  final String name;

  @override
  String toString() {
    return 'HostInfo(name: $name)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_HostInfo &&
            const DeepCollectionEquality().equals(other.name, name));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode =>
      Object.hash(runtimeType, const DeepCollectionEquality().hash(name));

  @JsonKey(ignore: true)
  @override
  _$$_HostInfoCopyWith<_$_HostInfo> get copyWith =>
      __$$_HostInfoCopyWithImpl<_$_HostInfo>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_HostInfoToJson(this);
  }
}

abstract class _HostInfo implements HostInfo {
  const factory _HostInfo({required final String name}) = _$_HostInfo;

  factory _HostInfo.fromJson(Map<String, dynamic> json) = _$_HostInfo.fromJson;

  @override
  String get name => throw _privateConstructorUsedError;
  @override
  @JsonKey(ignore: true)
  _$$_HostInfoCopyWith<_$_HostInfo> get copyWith =>
      throw _privateConstructorUsedError;
}
