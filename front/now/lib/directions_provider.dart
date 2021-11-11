import 'package:flutter/material.dart';
import 'package:google_maps_webservice/directions.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart' as maps;

class DirectionsProvider extends ChangeNotifier {
  GoogleMapsDirections directionsApi =
      GoogleMapsDirections(apiKey: "AIzaSyDR8XU6eKSgoECU2bfPy7_ZsPncLFiFCgM");

  Set<maps.Polyline> _route = Set();
  String _distanceTime = "";

  Set<maps.Polyline> get currentRoute => _route;
  String get distanceTime => _distanceTime;

  findDirections(maps.LatLng from, maps.LatLng to) async {
    var origin = Location(lat: from.latitude, lng: from.longitude);
    var destination = Location(lat: to.latitude, lng: to.longitude);

    var result = await directionsApi.directionsWithLocation(origin, destination,
        travelMode: TravelMode.walking);

    Set<maps.Polyline> newRoute = Set();

    if (result.isOkay) {
      var route = result.routes[0];
      var legs = route.legs[0];

      List<maps.LatLng> points = [];

      legs.steps.forEach((step) {
        points.add(maps.LatLng(step.startLocation.lat, step.startLocation.lng));
        points.add(maps.LatLng(step.endLocation.lat, step.endLocation.lng));
      });

      var polyLine = maps.Polyline(
        polylineId: const maps.PolylineId("Best route"),
        points: points,
        color: Colors.red,
        width: 4,
      );
      newRoute.add(polyLine);
      _route = newRoute;
      _distanceTime = legs.duration.text;
      notifyListeners();
    }
  }
}
