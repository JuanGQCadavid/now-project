import 'package:google_maps_flutter_platform_interface/src/types/location.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';

class LocationFakeService implements ILocationService {
  // TODO -> Implement this service!
  @override
  Future<LatLng> getUserCurrentLocation() async {
    return const LatLng(6.251723, -75.592771);
  }

  @override
  LatLng getUserMapNavigationLocation() {
    // TODO: implement getUserMapNavigationLocation
    throw UnimplementedError();
  }
}
