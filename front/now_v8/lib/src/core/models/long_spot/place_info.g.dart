// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'place_info.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_PlaceInfo _$$_PlaceInfoFromJson(Map<String, dynamic> json) => _$_PlaceInfo(
      name: json['name'] as String,
      lat: (json['lat'] as num).toDouble(),
      lon: (json['lon'] as num).toDouble(),
      mapProviderId: json['mapProviderId'] as String,
    );

Map<String, dynamic> _$$_PlaceInfoToJson(_$_PlaceInfo instance) =>
    <String, dynamic>{
      'name': instance.name,
      'lat': instance.lat,
      'lon': instance.lon,
      'mapProviderId': instance.mapProviderId,
    };
