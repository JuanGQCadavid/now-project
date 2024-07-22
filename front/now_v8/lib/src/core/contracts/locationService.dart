import 'package:google_maps_flutter/google_maps_flutter.dart';

abstract class ILocationService {

  // This method return the user physical current location
  // in the map
  Future<LatLng> getUserCurrentLocation();
}