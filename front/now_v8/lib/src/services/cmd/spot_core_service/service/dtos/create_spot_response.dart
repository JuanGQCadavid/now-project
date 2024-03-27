// import 'package:freezed_annotation/freezed_annotation.dart';

// part 'create_spot_response.freezed.dart';
// part 'create_spot_response.g.dart';

// @freezed
// class CreateSpotResponse with _$CreateSpotResponse {
//   const factory LongSpot(
//       {required EventInfo eventInfo,
//       required HostInfo hostInfo,
//       required PlaceInfo placeInfo,
//       required TopicsInfo topicInfo}) = _CreateSpotResponse;
//   factory CreateSpotResponse.fromJson(Map<String, Object?> json) =>
//       _$CreateSpotResponseFromJson(json);
// }

// @freezed
// class TopicsInfo with _$TopicsInfo {
//   const factory TopicsInfo({
//     required String principalTag,
//     required List<String> secondaryTags,
//   }) = _TopicsInfo;

//   factory TopicsInfo.fromJson(Map<String, Object?> json) =>
//       _$TopicsInfoFromJson(json);
// }

// @freezed
// class PlaceInfo with _$PlaceInfo {
//   const factory PlaceInfo(
//       {required String name,
//       required double lat,
//       required double lon,
//       required String mapProviderId}) = _PlaceInfo;

//   factory PlaceInfo.fromJson(Map<String, Object?> json) =>
//       _$PlaceInfoFromJson(json);
// }

// @freezed
// class HostInfo with _$HostInfo {
//   const factory HostInfo({required String name}) = _HostInfo;

//   factory HostInfo.fromJson(Map<String, Object?> json) =>
//       _$HostInfoFromJson(json);
// }

// @freezed
// class EventInfo with _$EventInfo {
//   const factory EventInfo({
//     required String name,
//     required String id,
//     required String description,
//     required int maximunCapacty,
//     required String emoji,
//   }) = _EventInfo;

//   factory EventInfo.fromJson(Map<String, Object?> json) =>
//       _$EventInfoFromJson(json);
// }
