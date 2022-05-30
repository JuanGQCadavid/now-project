import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/services/filter_service/service/dtos/filterSpotsResponse.dart';
import 'package:now_v8/src/services/filter_service/service/mappers/mappers.dart';
import 'package:now_v8/src/services/services_api_configuration.dart';

// TODO -> What could we do in case of a error ?


class FilterService implements IFilterService {
  late FilterServiceDTOsMappers mappers;
  final proximityResource = "/proximity";
  final ApiConfig apiConfig;
  late Dio dio;

  FilterService({required this.apiConfig}) {
    dio = Dio(
      BaseOptions(
        baseUrl: apiConfig.getFilterEndpoint(),
      ),
    );
    mappers = FilterServiceDTOsMappers();
  }

  @override
  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10,
  }) async {
    try {
      Response response = await dio.get(
        proximityResource,
        queryParameters: {
          "cpLat": cpLat,
          "cpLon": cpLng,
        },
      );

      if (response.statusCode! < 200 || response.statusCode! > 300) {
        print("there were an error");
        return List.empty();
      }

      Locations locations = Locations.fromJson(response.data);

      return mappers.fromPlacesToSpotList(locations);
    } catch (e) {
      print(e);
    }

    return List.empty();
  }
}
