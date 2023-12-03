import 'dart:developer';

import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/user_service/service/dtos/requests.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';
import 'package:now_v8/src/services/core/now_services_caller.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

const String InitLoginResource = "/init/login";
const String InitSingUpResource = "/init/singup";

class UserService implements IUserService {
  late NowServicesCaller nowServicesCaller;
  final ApiConfig apiConfig;

  UserService({required this.apiConfig}) {
    nowServicesCaller = NowServicesCaller(baseUrl: apiConfig.getUserEndpoint());
  }

  @override
  Future<Either<None, UserError>> initSingUp(
    String userPhoneNumber,
    String userName,
  ) async {
    InitSingUp request = InitSingUp(
      userPhoneNumber,
      userName,
      MethodVerificator("en", sms: true),
    );

    return await perfomPOST(InitSingUpResource, request.toJson());
  }

  @override
  Future<Either<None, UserError>> initLoging(String userPhoneNumber) async {
    InitLogin request = InitLogin(
      userPhoneNumber,
      MethodVerificator("en", sms: true),
    );

    return await perfomPOST(InitLoginResource, request.toJson());
  }

  Future<Either<None, UserError>> perfomPOST(
    String resource,
    Map<String, dynamic> body,
  ) async {
    print(body);

    Either<BackendErrors, dynamic> response = await nowServicesCaller.request(
      Method.POST,
      resource,
      data: body,
    );

    return response.fold((l) {
      return getUserError(l);
    }, (r) {
      return left(const None());
    });
  }

  Either<None, UserError> getUserError(BackendErrors l) {
    return l.when(
      clientError: (errorMessage) {
        log("We found a client error: ");
        log(errorMessage.toString());
        return right(mapMessageErrorToUserError(errorMessage));
      },
      internalError: () {
        return right(UserError.internalError("Internal error"));
      },
      resourceNotFound: () {
        return right(UserError.internalError("Internal error"));
      },
      serviceUnavailable: () {
        return right(UserError.internalError("Internal error"));
      },
      noInternetConnection: () {
        return right(UserError.internalError("Internal error"));
      },
      badResponseFormat: () {
        return right(UserError.internalError("Internal error"));
      },
      unknownError: () {
        return right(UserError.internalError("Internal error"));
      },
    );
  }
}
