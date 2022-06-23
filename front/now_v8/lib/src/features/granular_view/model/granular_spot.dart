import 'package:now_v8/src/core/models/spot.dart';

class GranularSpot {
  final GranularSpotWindow spotWindow;
  final Spot spot;

  GranularSpot({
    required this.spot,
    required this.spotWindow,
  });
}

class GranularSpotWindow {
  final String nextOne;
  final String actualOne;
  final String previousOne;

  const GranularSpotWindow({
    required this.actualOne,
    required this.nextOne,
    required this.previousOne,
  });
}
