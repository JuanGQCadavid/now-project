import 'package:google_maps_flutter_platform_interface/src/types/location.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';

class LocationFakeService implements ILocationService {
  // TODO -> Implement this service!
  @override
  Future<LatLng> getUserCurrentLocation() async {
    return const LatLng(6.261487428735279, -75.60657674320917);
  }

  @override
  LatLng getUserMapNavigationLocation() {
    // TODO: implement getUserMapNavigationLocation
    throw UnimplementedError();
  }
}
