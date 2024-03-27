import 'package:dartz/dartz.dart';
import 'package:flutter/services.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/token.dart';
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
    print("************************************");
    var response = await caller.request(
      Method.POST,
      createSpotResource,
      data: spot.toJson(),
      headers: token.toJson(),
    );

    print("Token ${token.header} - ${token.value}");
    return response.fold((l) {
      print("Error ${l}");
      return right(l);
    }, (r) {
      print("We sucesss!");
      var spot = LongSpot.fromJson(r);
      print("Id: ${spot.eventInfo.id}");
      return left(spot);
    });
  }

  @override
  Future<Either<LongSpot, BackendErrors>> fetchSpot(String spotID) {
    throw MissingPluginException();
  }
}
