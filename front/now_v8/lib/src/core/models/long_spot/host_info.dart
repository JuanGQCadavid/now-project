


import 'package:freezed_annotation/freezed_annotation.dart';

part 'host_info.freezed.dart';
part 'host_info.g.dart';

@freezed
class HostInfo with _$HostInfo {
  const factory HostInfo({
    required String name
  }) = _HostInfo;

  factory HostInfo.fromJson(Map<String, Object?> json)
      => _$HostInfoFromJson(json);
}