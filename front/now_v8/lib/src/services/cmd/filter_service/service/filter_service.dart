import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/models/long_spot.dart' as longSpot;
import 'package:google_maps_flutter/google_maps_flutter.dart';
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
  final searchHeader = "X-Now-Search-Session-Id";

  FilterService({required this.apiConfig}) {
    print(apiConfig.getFilterEndpoint());

    nowServicesCaller = NowServicesCaller(
      baseUrl: apiConfig.getFilterEndpoint(),
    );

    mappers = FilterServiceDTOsMappers();
  }

  @override
  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 0.03,
  }) async {
    Either<BackendErrors, dynamic> response = await nowServicesCaller
        .request(Method.GET, proximityResource, queryParameters: {
      "cpLat": cpLat,
      "cpLon": cpLng,
      "radious": radious,
    });

    return response.fold<List<Spot>>((l) {
      // On error! What are we going to do bro ?
      print("Where are on error ");
      print(l);
      return List.empty();
    }, (requestResponse) {
      print("We get a response");
      Locations locations =
          FilterProxymityResponse.fromJson(requestResponse).result;
      return mappers.fromPlacesToSpotList(locations);
    });
  }

  @override
  Future<StateResponse<List<Spot>, String>> getSpotsByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 10,
    String token = "",
  }) async {
    StateResponse backendResponse;

    backendResponse = await filterProximityState(
      cpLat: cpLat,
      cpLng: cpLng,
      token: token,
      format: "small",
      radious: radious,
      castFunction: fromFilterSpotToSpot,
    );

    List<Spot> response = [];

    for (var dynSpot in cast<List<dynamic>>(backendResponse.response)) {
      response.add(cast<Spot>(dynSpot));
    }
    String tokenResponse = cast<String>(backendResponse.token);

    return StateResponse<List<Spot>, String>(
        response: response, token: tokenResponse);
  }

  @override
  Future<StateResponse<List<longSpot.LongSpot>, String>>
      getLongSpotByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 0.03,
    String token = "",
  }) async {
    StateResponse backendResponse;

    backendResponse = await filterProximityState(
      cpLat: cpLat,
      cpLng: cpLng,
      token: token,
      format: "full",
      radious: radious,
      castFunction: fromFilterSpotToLongSpot,
    );

    List<longSpot.LongSpot> response = [];

    print("-------------------------------------------------------");
    for (var dynSpot in cast<List<dynamic>>(backendResponse.response)) {
      response.add(cast<longSpot.LongSpot>(dynSpot));
      print(cast<longSpot.LongSpot>(dynSpot));
    }

    print(response.length);

    
    print("-------------------------------------------------------");

    String tokenResponse = cast<String>(backendResponse.token);

    return StateResponse<List<longSpot.LongSpot>, String>(
        response: response, 
        token: tokenResponse,
    );

  }

  Future<StateResponse> filterProximityState({
    required double cpLat,
    required double cpLng,
    required String token,
    required Function(List<FilterSpot> places) castFunction,
    double radious = 0.5,
    String format = "small",
  }) async {
    Either<BackendErrors, dynamic> backendResponse;
    bool createSession = false;
    Map<String, dynamic>? headers;

    if (token.isEmpty) {
      print("token is empty then send with the create option");
      createSession = true;
    } else {
      print("token is not empty send it with the token header");
      createSession = false;
      headers = {searchHeader: token};
    }

    backendResponse =
        await nowServicesCaller.request(Method.GET, proximityResource,
            queryParameters: {
              "cpLat": cpLat,
              "cpLon": cpLng,
              "format": format,
              "createSession": createSession ? "true" : "false",
              "radious": radious
            },
            headers: headers);

    return backendResponse.fold((error) {
      print(error.toString());
      return StateResponse(response: [], token: "");
    }, (bodyResponse) {
      FilterProxyResponseWithState response =
          FilterProxyResponseWithState.fromJson(bodyResponse);
      return StateResponse(
          response: castFunction(response.result.places),
          token: response.search_session.session_details.session_id);
    });
  }

  List<longSpot.LongSpot> fromFilterSpotToLongSpot(List<FilterSpot> places) {
    List<longSpot.LongSpot> spots = [];

    for (var spot in places) {
      String hostName = spot.hostInfo?.name ?? "";
      spots.add(
        longSpot.LongSpot(
          hostInfo: longSpot.HostInfo(name: hostName),
          eventInfo: longSpot.EventInfo(
              description: spot.eventInfo.description,
              emoji: spot.eventInfo.emoji,
              id: spot.eventInfo.id,
              maximunCapacty: 0,
              name: spot.eventInfo.name),
          placeInfo: longSpot.PlaceInfo(
              lat: spot.placeInfo.lat,
              lon: spot.placeInfo.lon,
              mapProviderId: spot.placeInfo.mapProviderId,
              name: spot.placeInfo.name),
          topicInfo: longSpot.TopicsInfo(
              principalTopic: spot.topicInfo.principalTopic,
              secondaryTopics: spot.topicInfo.secondaryTopics),
          // MISSING
          dateInfo: longSpot.DateInfo(
            dateTime: spot.dateInfo.dateTime,
            id: spot.dateInfo.id,
            startTime: spot.dateInfo.startTime,
            durationApproximatedInSeconds: spot.dateInfo.durationApproximated,
          ),
        ),
      );
    }
    return spots;
  }

  List<Spot> fromFilterSpotToSpot(List<FilterSpot> places) {
    List<Spot> spots = [];

    for (var spot in places) {
      var timeFormatted = spot.dateInfo.startTime.split(" ");

      DateTime date;

      if (timeFormatted.length == 1) {
        date = DateTime.parse("${spot.dateInfo.dateTime}T${spot.dateInfo.startTime}");
      } else{
        date = DateTime.parse("${timeFormatted[0]} ${timeFormatted[1]}");
      }

      spots.add(
        Spot.withOutSpotColors(
            principalTag: spot.topicInfo.principalTopic.isNotEmpty ||
                    spot.topicInfo.secondaryTopics.isNotEmpty
                ? spot.topicInfo.principalTopic
                : spot.eventInfo.name
                    .toLowerCase()
                    .replaceAll(RegExp(r' '), ""),
            secondaryTags: spot.topicInfo.secondaryTopics,
            latLng: LatLng(
              spot.placeInfo.lat,
              spot.placeInfo.lon,
            ),
            spotId: spot.dateInfo.id,
            date: date),
      );
    }
    return spots;
  }
}



  // @override
  // Future<StateResponse<List<longSpot.LongSpot>, String>> getLongSpotByProximityWithState({
  //   required double cpLat,
  //   required double cpLng,
  //   double radious = 0.5,
  //   String token = "",
  // }) async {
  //   Either<BackendErrors, dynamic> backendResponse;

  //   backendResponse = await filterProximityState(
  //     cpLat: cpLat, 
  //     cpLng: cpLng,
  //     token: token,
  //     format: "full",
  //     radious: radious,
  //   );

  //   return backendResponse.fold<StateResponse<List<longSpot.LongSpot>, String>>((error) {
  //     print(error.toString());
  //     return StateResponse(response: [], token: "");
  //   }, (bodyResponse) {
  //     FilterProxyResponseWithState response =
  //         FilterProxyResponseWithState.fromJson(bodyResponse);
  //     return StateResponse(
  //         response: fromFilterSpotToLongSpot(response.result.places),
  //         token: response.search_session.session_details.session_id);
  //   });
  // }


  // @override
  // Future<StateResponse<List<Spot>, String>> getSpotsByProximityWithState({
  //   required double cpLat,
  //   required double cpLng,
  //   double radious = 10,
  //   String token = "",
  // }) async {
  //   Either<BackendErrors, dynamic> backendResponse;

  //   backendResponse = await filterProximityState(
  //     cpLat: cpLat, 
  //     cpLng: cpLng,
  //     token: token,
  //     format: "small",
  //     radious: radious,
  //   );

  //   return backendResponse.fold<StateResponse<List<Spot>, String>>((error) {
  //     print(error.toString());
  //     return StateResponse(response: [], token: "");
  //   }, (bodyResponse) {
  //     FilterProxyResponseWithState response =
  //         FilterProxyResponseWithState.fromJson(bodyResponse);
  //     return StateResponse(
  //         response: fromFilterSpotToSpot(response.result.places),
  //         token: response.search_session.session_details.session_id);
  //   });
  // }

  // Future<Either<BackendErrors, dynamic>> filterProximityState2({
  //   required double cpLat,
  //   required double cpLng,
  //   required String token,
  //   double radious = 0.5,
  //   String format = "small",
  // }){
  //   if (token.isEmpty) { 
  //     print("token is empty then send with the create option");
  //     return nowServicesCaller.request(
  //       Method.GET,
  //       proximityResource,
  //       queryParameters: {
  //         "cpLat": cpLat,
  //         "cpLon": cpLng,
  //         "format": format,
  //         "createSession": "true",
  //         "radious": radious
  //       },
  //     );

  //   }else {
  //     print("token is not empty send it with the token header");
  //     return nowServicesCaller.request(
  //       Method.GET,
  //       proximityResource,
  //       queryParameters: {
  //         "cpLat": cpLat,
  //         "cpLon": cpLng,
  //         "format": format,
  //         "createSession": "false",
  //         "radious": radious
  //       },
  //       headers: {
  //         "X-Now-Search-Session-Id": token
  //       },
  //     );
  //   }
    
  // }
