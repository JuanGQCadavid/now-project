// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'topics_info.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

TopicsInfo _$TopicsInfoFromJson(Map<String, dynamic> json) {
  return _TopicsInfo.fromJson(json);
}

/// @nodoc
mixin _$TopicsInfo {
  String get principalTag => throw _privateConstructorUsedError;
  List<String> get secondaryTags => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $TopicsInfoCopyWith<TopicsInfo> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $TopicsInfoCopyWith<$Res> {
  factory $TopicsInfoCopyWith(
          TopicsInfo value, $Res Function(TopicsInfo) then) =
      _$TopicsInfoCopyWithImpl<$Res>;
  $Res call({String principalTag, List<String> secondaryTags});
}

/// @nodoc
class _$TopicsInfoCopyWithImpl<$Res> implements $TopicsInfoCopyWith<$Res> {
  _$TopicsInfoCopyWithImpl(this._value, this._then);

  final TopicsInfo _value;
  // ignore: unused_field
  final $Res Function(TopicsInfo) _then;

  @override
  $Res call({
    Object? principalTag = freezed,
    Object? secondaryTags = freezed,
  }) {
    return _then(_value.copyWith(
      principalTag: principalTag == freezed
          ? _value.principalTag
          : principalTag // ignore: cast_nullable_to_non_nullable
              as String,
      secondaryTags: secondaryTags == freezed
          ? _value.secondaryTags
          : secondaryTags // ignore: cast_nullable_to_non_nullable
              as List<String>,
    ));
  }
}

/// @nodoc
abstract class _$$_TopicsInfoCopyWith<$Res>
    implements $TopicsInfoCopyWith<$Res> {
  factory _$$_TopicsInfoCopyWith(
          _$_TopicsInfo value, $Res Function(_$_TopicsInfo) then) =
      __$$_TopicsInfoCopyWithImpl<$Res>;
  @override
  $Res call({String principalTag, List<String> secondaryTags});
}

/// @nodoc
class __$$_TopicsInfoCopyWithImpl<$Res> extends _$TopicsInfoCopyWithImpl<$Res>
    implements _$$_TopicsInfoCopyWith<$Res> {
  __$$_TopicsInfoCopyWithImpl(
      _$_TopicsInfo _value, $Res Function(_$_TopicsInfo) _then)
      : super(_value, (v) => _then(v as _$_TopicsInfo));

  @override
  _$_TopicsInfo get _value => super._value as _$_TopicsInfo;

  @override
  $Res call({
    Object? principalTag = freezed,
    Object? secondaryTags = freezed,
  }) {
    return _then(_$_TopicsInfo(
      principalTag: principalTag == freezed
          ? _value.principalTag
          : principalTag // ignore: cast_nullable_to_non_nullable
              as String,
      secondaryTags: secondaryTags == freezed
          ? _value._secondaryTags
          : secondaryTags // ignore: cast_nullable_to_non_nullable
              as List<String>,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_TopicsInfo implements _TopicsInfo {
  const _$_TopicsInfo(
      {required this.principalTag, required final List<String> secondaryTags})
      : _secondaryTags = secondaryTags;

  factory _$_TopicsInfo.fromJson(Map<String, dynamic> json) =>
      _$$_TopicsInfoFromJson(json);

  @override
  final String principalTag;
  final List<String> _secondaryTags;
  @override
  List<String> get secondaryTags {
    // ignore: implicit_dynamic_type
    return EqualUnmodifiableListView(_secondaryTags);
  }

  @override
  String toString() {
    return 'TopicsInfo(principalTag: $principalTag, secondaryTags: $secondaryTags)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_TopicsInfo &&
            const DeepCollectionEquality()
                .equals(other.principalTag, principalTag) &&
            const DeepCollectionEquality()
                .equals(other._secondaryTags, _secondaryTags));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(principalTag),
      const DeepCollectionEquality().hash(_secondaryTags));

  @JsonKey(ignore: true)
  @override
  _$$_TopicsInfoCopyWith<_$_TopicsInfo> get copyWith =>
      __$$_TopicsInfoCopyWithImpl<_$_TopicsInfo>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_TopicsInfoToJson(this);
  }
}

abstract class _TopicsInfo implements TopicsInfo {
  const factory _TopicsInfo(
      {required final String principalTag,
      required final List<String> secondaryTags}) = _$_TopicsInfo;

  factory _TopicsInfo.fromJson(Map<String, dynamic> json) =
      _$_TopicsInfo.fromJson;

  @override
  String get principalTag => throw _privateConstructorUsedError;
  @override
  List<String> get secondaryTags => throw _privateConstructorUsedError;
  @override
  @JsonKey(ignore: true)
  _$$_TopicsInfoCopyWith<_$_TopicsInfo> get copyWith =>
      throw _privateConstructorUsedError;
}
