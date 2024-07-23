import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

abstract class IGCPServices {
  Future<Either<List<PlaceInfo>, String>> findPlacesByName(
    String placeName,
  );

  Future<Either<List<PlaceInfo>, String>> findPlacesByLatLon(
    double lat,
    double lon,
  );
}