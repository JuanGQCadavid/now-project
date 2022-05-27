import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:flutter/material.dart';


// I will need to investigate more about freezed in order to perfom copyWithData or something like that 
// What about builder ?
class Spot {
  final String principalTag;
  final List<String> secondaryTags;
  final LatLng latLng;
  final String spotId;
  SpotsColors spotsColor;

  Spot({
    required this.principalTag,
    required this.secondaryTags,
    required this.latLng,
    required this.spotId,
    required this.spotsColor
  });

  Spot.withOutSpotColors({
    required this.principalTag,
    required this.secondaryTags,
    required this.latLng,
    required this.spotId,
    this.spotsColor = const SpotsColors.empty()
  });

}
