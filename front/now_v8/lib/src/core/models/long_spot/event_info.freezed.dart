// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target

part of 'event_info.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

EventInfo _$EventInfoFromJson(Map<String, dynamic> json) {
  return _EventInfo.fromJson(json);
}

/// @nodoc
mixin _$EventInfo {
  String get name => throw _privateConstructorUsedError;
  String get id => throw _privateConstructorUsedError;
  String get description => throw _privateConstructorUsedError;
  int get maximunCapacty => throw _privateConstructorUsedError;
  String get eventType => throw _privateConstructorUsedError;
  String get emoji => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $EventInfoCopyWith<EventInfo> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $EventInfoCopyWith<$Res> {
  factory $EventInfoCopyWith(EventInfo value, $Res Function(EventInfo) then) =
      _$EventInfoCopyWithImpl<$Res>;
  $Res call(
      {String name,
      String id,
      String description,
      int maximunCapacty,
      String eventType,
      String emoji});
}

/// @nodoc
class _$EventInfoCopyWithImpl<$Res> implements $EventInfoCopyWith<$Res> {
  _$EventInfoCopyWithImpl(this._value, this._then);

  final EventInfo _value;
  // ignore: unused_field
  final $Res Function(EventInfo) _then;

  @override
  $Res call({
    Object? name = freezed,
    Object? id = freezed,
    Object? description = freezed,
    Object? maximunCapacty = freezed,
    Object? eventType = freezed,
    Object? emoji = freezed,
  }) {
    return _then(_value.copyWith(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      id: id == freezed
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      description: description == freezed
          ? _value.description
          : description // ignore: cast_nullable_to_non_nullable
              as String,
      maximunCapacty: maximunCapacty == freezed
          ? _value.maximunCapacty
          : maximunCapacty // ignore: cast_nullable_to_non_nullable
              as int,
      eventType: eventType == freezed
          ? _value.eventType
          : eventType // ignore: cast_nullable_to_non_nullable
              as String,
      emoji: emoji == freezed
          ? _value.emoji
          : emoji // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
abstract class _$$_EventInfoCopyWith<$Res> implements $EventInfoCopyWith<$Res> {
  factory _$$_EventInfoCopyWith(
          _$_EventInfo value, $Res Function(_$_EventInfo) then) =
      __$$_EventInfoCopyWithImpl<$Res>;
  @override
  $Res call(
      {String name,
      String id,
      String description,
      int maximunCapacty,
      String eventType,
      String emoji});
}

/// @nodoc
class __$$_EventInfoCopyWithImpl<$Res> extends _$EventInfoCopyWithImpl<$Res>
    implements _$$_EventInfoCopyWith<$Res> {
  __$$_EventInfoCopyWithImpl(
      _$_EventInfo _value, $Res Function(_$_EventInfo) _then)
      : super(_value, (v) => _then(v as _$_EventInfo));

  @override
  _$_EventInfo get _value => super._value as _$_EventInfo;

  @override
  $Res call({
    Object? name = freezed,
    Object? id = freezed,
    Object? description = freezed,
    Object? maximunCapacty = freezed,
    Object? eventType = freezed,
    Object? emoji = freezed,
  }) {
    return _then(_$_EventInfo(
      name: name == freezed
          ? _value.name
          : name // ignore: cast_nullable_to_non_nullable
              as String,
      id: id == freezed
          ? _value.id
          : id // ignore: cast_nullable_to_non_nullable
              as String,
      description: description == freezed
          ? _value.description
          : description // ignore: cast_nullable_to_non_nullable
              as String,
      maximunCapacty: maximunCapacty == freezed
          ? _value.maximunCapacty
          : maximunCapacty // ignore: cast_nullable_to_non_nullable
              as int,
      eventType: eventType == freezed
          ? _value.eventType
          : eventType // ignore: cast_nullable_to_non_nullable
              as String,
      emoji: emoji == freezed
          ? _value.emoji
          : emoji // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_EventInfo with DiagnosticableTreeMixin implements _EventInfo {
  const _$_EventInfo(
      {required this.name,
      required this.id,
      required this.description,
      required this.maximunCapacty,
      required this.eventType,
      required this.emoji});

  factory _$_EventInfo.fromJson(Map<String, dynamic> json) =>
      _$$_EventInfoFromJson(json);

  @override
  final String name;
  @override
  final String id;
  @override
  final String description;
  @override
  final int maximunCapacty;
  @override
  final String eventType;
  @override
  final String emoji;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'EventInfo(name: $name, id: $id, description: $description, maximunCapacty: $maximunCapacty, eventType: $eventType, emoji: $emoji)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'EventInfo'))
      ..add(DiagnosticsProperty('name', name))
      ..add(DiagnosticsProperty('id', id))
      ..add(DiagnosticsProperty('description', description))
      ..add(DiagnosticsProperty('maximunCapacty', maximunCapacty))
      ..add(DiagnosticsProperty('eventType', eventType))
      ..add(DiagnosticsProperty('emoji', emoji));
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_EventInfo &&
            const DeepCollectionEquality().equals(other.name, name) &&
            const DeepCollectionEquality().equals(other.id, id) &&
            const DeepCollectionEquality()
                .equals(other.description, description) &&
            const DeepCollectionEquality()
                .equals(other.maximunCapacty, maximunCapacty) &&
            const DeepCollectionEquality().equals(other.eventType, eventType) &&
            const DeepCollectionEquality().equals(other.emoji, emoji));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(
      runtimeType,
      const DeepCollectionEquality().hash(name),
      const DeepCollectionEquality().hash(id),
      const DeepCollectionEquality().hash(description),
      const DeepCollectionEquality().hash(maximunCapacty),
      const DeepCollectionEquality().hash(eventType),
      const DeepCollectionEquality().hash(emoji));

  @JsonKey(ignore: true)
  @override
  _$$_EventInfoCopyWith<_$_EventInfo> get copyWith =>
      __$$_EventInfoCopyWithImpl<_$_EventInfo>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_EventInfoToJson(this);
  }
}

abstract class _EventInfo implements EventInfo {
  const factory _EventInfo(
      {required final String name,
      required final String id,
      required final String description,
      required final int maximunCapacty,
      required final String eventType,
      required final String emoji}) = _$_EventInfo;

  factory _EventInfo.fromJson(Map<String, dynamic> json) =
      _$_EventInfo.fromJson;

  @override
  String get name => throw _privateConstructorUsedError;
  @override
  String get id => throw _privateConstructorUsedError;
  @override
  String get description => throw _privateConstructorUsedError;
  @override
  int get maximunCapacty => throw _privateConstructorUsedError;
  @override
  String get eventType => throw _privateConstructorUsedError;
  @override
  String get emoji => throw _privateConstructorUsedError;
  @override
  @JsonKey(ignore: true)
  _$$_EventInfoCopyWith<_$_EventInfo> get copyWith =>
      throw _privateConstructorUsedError;
}
