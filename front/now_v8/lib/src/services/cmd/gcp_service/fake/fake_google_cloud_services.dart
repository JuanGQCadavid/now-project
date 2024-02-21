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

  final List<PlaceInfo> latLngPlaces = [
    const PlaceInfo(
      name: "Latiaoladraba -#AT#- Circular 4 # 51 - 52",
      lat: 6.251723,
      lon: -75.592771,
      mapProviderId: "HOME1",
    ),
    const PlaceInfo(
      name: "StandUp SF -#AT#- Circular4 # 55 - 22",
      lat: 6.251733,
      lon: -75.592791,
      mapProviderId: "HOME2",
    ),
  ];

  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByLatLon(
    double lat,
    double lon,
  ) async {
    return left(latLngPlaces);
  }

  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByName(
    String placeName,
  ) async {
    return left(places);
  }
}
