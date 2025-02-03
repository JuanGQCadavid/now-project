import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/notifiers.dart';

class SpotsCreatorCore {
  final IGCPServices gpcService;
  final ISpotCoreService coreService;
  final AuthState authState;

  const SpotsCreatorCore({
    required this.gpcService,
    required this.coreService,
    required this.authState,
  });

  Future<Either<List<PlaceInfo>, String>> getOptions(String placeName) async {
    return await gpcService.findPlacesByName(placeName);
  }

  Future<Either<List<PlaceInfo>, String>> getAproximatedPlaces(
      double lat, lng) async {
    return await gpcService.findPlacesByLatLon(lat, lng);
  }

  Future<Either<LongSpot, BackendErrors>> createSpot(LongSpot spot) async {
    var response = await authState.getToken();
    return response.fold(
      (token) => coreService.createSpot(spot, token),
      (r) => right(
        BackendErrors.clientError(
          ErrorMessage(
            "LOCAL",
            "No user saved",
            "local",
          ),
        ),
      ),
    );
  }
}
