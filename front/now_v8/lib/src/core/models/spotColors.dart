import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

// TODO -> What could we do with freezed ?
class SpotsColors {
  final Color color;
  final BitmapDescriptor hue;

  const SpotsColors({
    required this.color,
    required this.hue,
  });

  const SpotsColors.empty({
    this.color = Colors.black,
    this.hue = BitmapDescriptor.defaultMarker,
  });
}
