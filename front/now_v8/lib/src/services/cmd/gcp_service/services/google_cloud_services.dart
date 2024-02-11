import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:dio/dio.dart';

class GoogleCloudServices implements IGCPServices {
  late Dio _dio;
  late Options placeOptions;
  late Options geoOptions;

  GoogleCloudServices() {
    _dio = Dio();
    placeOptions = Options(headers: {
      "X-Goog-Api-Key": "",
      "X-Goog-FieldMask":
          "places.location,places.shortFormattedAddress,places.id,places.name"
    });
  }

  // https://developers.google.com/maps/documentation/places/web-service/text-search
  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByLatLon(
    double lat,
    double lon,
  ) {
    throw UnimplementedError();
  }

  // https://developers.google.com/maps/documentation/geocoding/requests-reverse-geocoding
  @override
  Future<Either<List<PlaceInfo>, String>> findPlacesByName(
    String placeName,
  ) async {
    var response = await _dio.post(
      "https://places.googleapis.com/v1/places:searchText",
      data: {
        "textQuery": placeName,
      },
      options: placeOptions,
    );

    print(response);

    return left([
      const PlaceInfo(
        name: "name",
        lat: 0,
        lon: 0,
        mapProviderId: "mapProviderId",
      )
    ]);
  }
}
