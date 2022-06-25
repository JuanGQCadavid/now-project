

import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

part 'place_info.freezed.dart';
part 'place_info.g.dart';

@freezed
class PlaceInfo with _$PlaceInfo{
  const factory PlaceInfo({
    required String name,
    required double lat,
    required double lon,
    required String mapProviderId
  }) = _PlaceInfo;

  factory PlaceInfo.fromJson(Map<String, Object?> json)
      => _$PlaceInfoFromJson(json);
}