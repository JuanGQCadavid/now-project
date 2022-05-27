import 'package:google_maps_flutter/google_maps_flutter.dart';

abstract class ILocationService {

  // This method return the user physical current location
  // in the map
  LatLng getUserCurrentLocation();

  // This method will return the location that the user
  // is while browsing in the map
  LatLng getUserMapNavigationLocation();
}