import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

class GoogleCloudServices implements IGCPServices {
  // https://developers.google.com/maps/documentation/places/web-service/text-search
  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByLatLon(
    double lat,
    double lon,
  ) {
    // TODO: implement findPlacesByLatLon
    throw UnimplementedError();
  }

  // https://developers.google.com/maps/documentation/geocoding/requests-reverse-geocoding
  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByName(
    String placeName,
  ) {
    // TODO: implement findPlacesByName
    throw UnimplementedError();
  }
}
