

import 'package:freezed_annotation/freezed_annotation.dart';

part 'topics_info.freezed.dart';
part 'topics_info.g.dart';

@freezed
class TopicsInfo with _$TopicsInfo {
  const factory TopicsInfo({
    required String principalTag,
    required List<String> secondaryTags,
  }) = _TopicsInfo;

  factory TopicsInfo.fromJson(Map<String, Object?> json)
      => _$TopicsInfoFromJson(json);
}