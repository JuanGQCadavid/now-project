

import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:now_v8/src/core/models/long_spot/event_info.dart';
import 'package:now_v8/src/core/models/long_spot/host_info.dart';
import 'package:now_v8/src/core/models/long_spot/place_info.dart';
import 'package:now_v8/src/core/models/long_spot/topics_info.dart';

part 'long_spot.freezed.dart';
part 'long_spot.g.dart';

@freezed
class LongSpot with _$LongSpot{
  const factory LongSpot({
    required EventInfo eventInfo,
    required HostInfo hostInfo,
    required PlaceInfo placeInfo,
    required TopicsInfo topicInfo
  }) = _LongSpot;
  factory LongSpot.fromJson(Map<String, Object?> json)
      => _$LongSpotFromJson(json);
}