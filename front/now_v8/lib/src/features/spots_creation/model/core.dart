import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

class SpotsCreatorCore {
  final IGCPServices gpcService;

  const SpotsCreatorCore({required this.gpcService});

  Future<Either<List<PlaceInfo>, String>> getOptions(String placeName) async {
    return await gpcService.findPlacesByName(placeName);
  }

  Future<Either<List<PlaceInfo>, String>> getAproximatedPlaces(
      double lat, lng) async {
    return await gpcService.findPlacesByLatLon(lat, lng);
  }
}
