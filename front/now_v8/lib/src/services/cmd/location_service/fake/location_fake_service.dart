import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/location_service.dart';

class LocationFakeService implements ILocationService {
  @override
  Future<LatLng> getUserCurrentLocation() async {
    // await Future.delayed(Duration(seconds: 3)); // Simulate delay
    return const LatLng(6.261487428735279, -75.60657674320917);
  }
}
