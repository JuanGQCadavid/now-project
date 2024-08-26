import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/models/spotColors.dart';

class ColorsService implements IColorService {
  int index = 0;
  final List<SpotsColors> _availableColors = [
    SpotsColors(
      color: Colors.red,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueRed),
    ),
    SpotsColors(
      color: Colors.blue,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueAzure),
    ),
    SpotsColors(
      color: Colors.cyan,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueCyan),
    ),
    SpotsColors(
      color: Colors.green,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueGreen),
    ),
    SpotsColors(
      color: Colors.pink,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueMagenta),
    ),
    SpotsColors(
      color: Colors.purple,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueViolet),
    ),
    SpotsColors(
      color: Colors.orange.shade800,
      hue: BitmapDescriptor.defaultMarkerWithHue(1),
    ),
    SpotsColors(
      color: Colors.blue.shade700,
      hue: BitmapDescriptor.defaultMarkerWithHue(200),
    ),
    SpotsColors(
      color: Colors.pink.shade700,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueRose),
    ),
    SpotsColors(
      color: Colors.grey.shade400,
      hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueOrange),
    )
  ];

  ColorsService();

  @override
  SpotsColors getColor() {
    return _availableColors[2];
  }

  SpotsColors getRandomColor() {
    if (index >= _availableColors.length) {
      index = 0;
    }

    SpotsColors colorToReturn = _availableColors[index];
    index = index + 1;

    return colorToReturn;
  }

  @override
  SpotsColors getScheduleColor() {
    return _availableColors[_availableColors.length - 1];
  }
}
