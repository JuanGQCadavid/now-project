import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/models/spotColors.dart';

class ColorsService implements IColorService{
  int index = 0;
  List<SpotsColors> _availableColors =[];

  ColorsService() {
    populateColors();
  }

  SpotsColors getColor() {
    if (index >= _availableColors.length) {
      index = 0;
    }

    SpotsColors colorToReturn = _availableColors[index];
    index = index + 1;

    return colorToReturn;
  }

  void populateColors() {

    _availableColors.add(
      SpotsColors(
        color: Colors.red,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueRed),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.blue,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueAzure),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.cyan,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueCyan),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.green,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueGreen),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.pink,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueMagenta),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.purple,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueViolet),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.orange.shade800,
        hue: BitmapDescriptor.defaultMarkerWithHue(1),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.blue.shade700,
        hue: BitmapDescriptor.defaultMarkerWithHue(200),
      ),
    );

    _availableColors.add(
      SpotsColors(
        color: Colors.pink.shade700,
        hue: BitmapDescriptor.defaultMarkerWithHue(BitmapDescriptor.hueRose),
      ),
    );

  }
}
