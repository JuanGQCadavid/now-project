
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'event_info.freezed.dart';
part 'event_info.g.dart';

@freezed
class EventInfo with _$EventInfo {
  const factory EventInfo({
    required String name,
    required String id,
    required String description,
    required int maximunCapacty,
    required String eventType,
    required String emoji,
  }) = _EventInfo;

  factory EventInfo.fromJson(Map<String, Object?> json)
      => _$EventInfoFromJson(json);
}