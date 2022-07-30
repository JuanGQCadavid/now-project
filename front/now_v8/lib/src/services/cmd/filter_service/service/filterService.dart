
import 'package:dartz/dartz.dart';
import 'package:dio/dio.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/state_response.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/dtos/filterSpotsResponse.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/mappers/mappers.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';
import 'package:now_v8/src/services/core/now_services_caller.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

class FilterService implements IFilterService {
  late FilterServiceDTOsMappers mappers;
  late NowServicesCaller nowServicesCaller;
  final ApiConfig apiConfig;

  // Constants
  final proximityResource = "/proximity";

  FilterService({required this.apiConfig}) {
    nowServicesCaller = NowServicesCaller(
      baseUrl: apiConfig.getFilterEndpoint(),
    );

    mappers = FilterServiceDTOsMappers();
  }

  @override
  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10,
  }) async {
    Either<BackendErrors, dynamic> response = await nowServicesCaller
        .request(Method.GET, proximityResource, queryParameters: {
      "cpLat": cpLat,
      "cpLon": cpLng,
    });

    return response.fold<List<Spot>>((l) {
      // On error! What are we going to do bro ?
      return List.empty();
    }, (requestResponse) {
      Locations locations = FilterProxymityResponse.fromJson(requestResponse).result;
      return mappers.fromPlacesToSpotList(locations);
    });
    
  }

  Future<StateResponse<List<LongSpot>, String>> getByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 10,
    String token = "",
  }) async {
    return StateResponse(response: [], token: "token");
  }
}
