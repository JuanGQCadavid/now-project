import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

class FakeGoogleServices implements IGCPServices {
  final List<PlaceInfo> places = [
    const PlaceInfo(
      name: "WeWork -#AT#- Carrera 1 # 23 - 52",
      lat: 6.245617,
      lon: -75.587757,
      mapProviderId: "A",
    ),
    const PlaceInfo(
      name: "Loft Luxury -#AT#- Carrera 3 # 12 - 22",
      lat: 6.245873,
      lon: -75.591319,
      mapProviderId: "B",
    ),
    const PlaceInfo(
      name: "House of dragon -#AT#- Circular 15 # 7 - 8",
      lat: 6.248561,
      lon: -75.589731,
      mapProviderId: "C",
    )
  ];
  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByLatLon(
    double lat,
    double lon,
  ) async {
    return left(places);
  }

  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByName(
    String placeName,
  ) async {
    // TODO: implement findPlacesByName
    return left(places);
  }
}
