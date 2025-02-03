import 'dart:developer';

import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/profile_service.dart';
import 'package:now_v8/src/core/models/profile.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';
import 'package:now_v8/src/services/core/now_services_caller.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

class UserProfileService implements IUserProfileService {
  final ApiConfig apiConfig;
  final NowServicesCaller caller;

  final String userIdPattern = ":userId";
  final String getProfilePath = "/:userId";

  const UserProfileService({
    required this.apiConfig,
    required this.caller,
  });

  @override
  Future<Either<UserProfile, BackendErrors>> getUserProfile(
      String userId, Token token) async {
    var callPath = getProfilePath.replaceAll(userIdPattern, userId);
    var response = await caller.request(
      Method.GET,
      callPath,
      headers: token.toJson(),
    );
    return response.fold(
      (l) {
        log("Error on calling profile service on path - $callPath, \n error: $l");
        return right(l);
      },
      (r) {
        try {
          return left(UserProfile.fromJson(r));
        } catch (error) {
          log("We face an error while casting the backend response, request path - $callPath, \n error: ${error.toString()}");
          return right(BackendErrors.internalError());
        }
      },
    );
  }
}
