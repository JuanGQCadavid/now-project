import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/spot_core_service/service/service.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/notifiers.dart';

class SpotsCreatorCore {
  final IGCPServices gpcService;
  final ISpotCoreService coreService;
  final Either<UserDetails, None<dynamic>> auth;

  const SpotsCreatorCore({
    required this.gpcService,
    required this.coreService,
    required this.auth,
  });

  Future<Either<List<PlaceInfo>, String>> getOptions(String placeName) async {
    return await gpcService.findPlacesByName(placeName);
  }

  Future<Either<List<PlaceInfo>, String>> getAproximatedPlaces(
      double lat, lng) async {
    return await gpcService.findPlacesByLatLon(lat, lng);
  }

  Future<Either<LongSpot, BackendErrors>> createSpot(LongSpot spot) async {
    return auth.fold(
      (l) => coreService.createSpot(
          spot,
          Token(
            // TODO: WE SHOULD MOVE THIS TO A CENTRAL PLACE
            header: "X-Authorization",
            value: l.shortLiveToken,
          )),
      (r) => right(
        BackendErrors.clientError(
          ErrorMessage(
            "LOCAL",
            "User is not logged",
            "LOCAL",
          ),
        ),
      ),
    );
  }
}
