import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';

import 'package:now_v8/src/core/models/long_spot/host_info.dart';
import 'package:now_v8/src/core/models/long_spot/event_info.dart' as long_event;
import 'package:now_v8/src/core/models/long_spot/place_info.dart' as long_place;
import 'package:now_v8/src/core/models/long_spot/topics_info.dart'
    as long_topic;

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
      Locations locations =
          FilterProxymityResponse.fromJson(requestResponse).result;
      return mappers.fromPlacesToSpotList(locations);
    });
  }

  Future<StateResponse<List<LongSpot>, String>> getByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 10,
    String token = "",
  }) async {
    Either<BackendErrors, dynamic> backendResponse;

    if (token.isEmpty) {
      // If token is empty then send with the create option
      print("token is empty then send with the create option");
      backendResponse = await nowServicesCaller.request(
        Method.GET,
        proximityResource,
        queryParameters: {
          "cpLat": cpLat,
          "cpLon": cpLng,
          "createSession": "true"
        },
      );
    } else {
      // If token is not empty send it with the token header
      print("token is not empty send it with the token header");
      backendResponse = await nowServicesCaller
          .request(Method.GET, proximityResource, queryParameters: {
        "cpLat": cpLat,
        "cpLon": cpLng,
      }, headers: {
        "X-Now-Search-Session-Id": token
      });
    }

    return backendResponse.fold<StateResponse<List<LongSpot>, String>>((error) {
      print("Oh shit, an error!");
      print(error.toString());

      return StateResponse(response: [], token: "");
    }, (bodyResponse) {
      FilterProxyResponseWithState response =
          FilterProxyResponseWithState.fromJson(bodyResponse);

      List<LongSpot> spots = [];

      response.result.places.forEach(
        (FilterSpot spot) {
          spots.add(
            LongSpot(
              hostInfo: const HostInfo(name: "MISSING_FROM_BACKEND"),
              eventInfo: long_event.EventInfo(
                  description: "MISSING_FROM_BACKEND",
                  emoji: spot.eventInfo.emoji,
                  eventType: spot.eventInfo.eventType,
                  id: spot.eventInfo.id,
                  maximunCapacty: 0,
                  name: spot.eventInfo.name),
              placeInfo: long_place.PlaceInfo(
                  lat: spot.placeInfo.lat,
                  lon: spot.placeInfo.lon,
                  mapProviderId: spot.placeInfo.mapProviderId,
                  name: "MISSING_FROM_BACKEND"),
              topicInfo: long_topic.TopicsInfo(
                principalTag: spot.topicInfo.principalTopic,
                secondaryTags: spot.topicInfo.secondaryTopics
              ),
            ),
          );
        },
      );

      return StateResponse(
          response: spots,
          token: response.search_session.session_details.session_id);
    });
  }
}
