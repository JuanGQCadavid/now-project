import 'package:dartz/dartz.dart';
import 'package:flutter/services.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/services/cmd/spot_core_service/service/dtos/create_spot_response.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';
import 'package:now_v8/src/services/core/now_services_caller.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

class SpotCoreService implements ISpotCoreService {
  final ApiConfig apiConfig;
  final NowServicesCaller caller;

  final String createSpotResource = "/";
  final String fetchSpotResource = "/";

  const SpotCoreService({
    required this.apiConfig,
    required this.caller,
  });

  @override
  Future<Either<LongSpot, BackendErrors>> createSpot(
    LongSpot spot,
    Token token,
  ) async {
    var response = await caller.request(
      Method.POST,
      createSpotResource,
      data: spot.toJson(),
      headers: token.toJson(),
    );

    return response.fold((l) {
      print("Error ${l}");
      return right(l);
    }, (r) {
      print("We sucesss!");
      var spotID = CreateSpotResponse.fromJson(r);
      print("Id: ${spotID.eventInfo.id}");

      var finalSpot = spot.copyWith(
        eventInfo: spot.eventInfo.copyWith(id: spotID.eventInfo.id),
      );

      return left(finalSpot);
    });
  }

  @override
  Future<Either<LongSpot, BackendErrors>> fetchSpot(String spotID) {
    throw MissingPluginException();
  }
}
