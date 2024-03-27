import 'package:freezed_annotation/freezed_annotation.dart';

part 'create_spot_response.g.dart';

@JsonSerializable()
class CreateSpotResponse {
  final EventInfo eventInfo;

  CreateSpotResponse({
    required this.eventInfo,
  });

  factory CreateSpotResponse.fromJson(Map<String, dynamic> json) =>
      _$CreateSpotResponseFromJson(json);

  Map<String, dynamic> toJson() => _$CreateSpotResponseToJson(this);
}

@JsonSerializable()
class EventInfo {
  final String id;

  EventInfo({
    required this.id,
  });

  factory EventInfo.fromJson(Map<String, dynamic> json) =>
      _$EventInfoFromJson(json);

  Map<String, dynamic> toJson() => _$EventInfoToJson(this);
}
