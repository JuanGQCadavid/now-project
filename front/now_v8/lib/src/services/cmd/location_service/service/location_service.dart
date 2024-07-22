import 'package:google_maps_flutter_platform_interface/src/types/location.dart';
import 'package:location/location.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';

class LocationService implements ILocationService {
  late Location location;

  LocationService() {
    location = Location();
  }
  @override
  Future<LatLng> getUserCurrentLocation() async {
    try {
      // bool _serviceEnabled;
      // PermissionStatus _permissionGranted;
      LocationData _locationData;

      // _serviceEnabled = await location.serviceEnabled();
      // if (!_serviceEnabled) {
      //   print("Was not enable");
      //   _serviceEnabled = await location.requestService();
      //   if (!_serviceEnabled) {
      //     print("Was not enable twice");
      //     return LatLng(0, 0);
      //   }
      // }

      // _permissionGranted = await location.hasPermission();
      // print(_permissionGranted);
      // if (_permissionGranted == PermissionStatus.denied) {
      //   print("Permission Denied");
      //   _permissionGranted = await location.requestPermission();
      //   if (_permissionGranted != PermissionStatus.granted) {
      //     print("Permission Denied x2");
      //     return LatLng(0, 0);
      //   }
      // }

      _locationData = await location.getLocation();
      return LatLng(_locationData.latitude!, _locationData.longitude!);
    } catch (e) {
      print("ERROR!!! - error!");
      print(e.toString());
      return const LatLng(6.251723, -75.592771);
    }
  }
}
