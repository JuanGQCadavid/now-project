import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/user_service/service/service.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';

class AuthState extends StateNotifier<Either<UserDetails, None>> {
  final IAuthService authService;
  final IUserService userService;
  final String authHeader = "X-Auth";

  AuthState({
    required this.authService,
    required this.userService,
  }) : super(right(const None())) {
    initState();
  }

  initState() {
    authService.getUserDetails().then((value) {
      state = value.fold(
        (l) => left(l),
        (r) => right(
          const None(),
        ),
      );
    });
  }

  Future<Either<Token, BackendErrors>> getToken() async {
    var resp = await authService.getUserDetails();
    return resp.fold(
      (l) => left(
        Token(
          header: authHeader,
          value: l.shortLiveToken,
        ),
      ),
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

  Future userLogIn(UserDetails details) async {
    await authService.storeUserDetails(details);
    state = left(details);
  }

  Future userLogOut() async {
    await authService.removeUserDetails();
    state = right(const None());
  }

  /// Exposing User Interface
  Future<Either<None, UserError>> initLoging(
    String userPhoneNumber,
  ) =>
      userService.initLoging(userPhoneNumber);

  Future<Either<None, UserError>> initSingUp(
    String userPhoneNumber,
    String userName,
  ) =>
      userService.initSingUp(userPhoneNumber, userName);

  Future<Either<UserDetails, UserError>> validate(
    String username,
    String userPhoneNumber,
    List<String> userCode,
  ) async {
    var response = await userService.validate(userPhoneNumber, userCode);

    await response.fold((userDetails) async {
      print("ID: " +
          userDetails.userId +
          "Tokens" +
          "\n" +
          userDetails.refreshToken +
          "\n" +
          userDetails.shortLiveToken +
          "\n" +
          userDetails.shortLiveTokenTTL +
          "\n");

      await userLogIn(
        UserDetails(
          userId: userDetails.userId,
          userName: username,
          refreshToken: userDetails.refreshToken,
          shortLiveToken: userDetails.shortLiveToken,
          shortLiveTokenTTL: userDetails.shortLiveTokenTTL,
        ),
      );
    }, (r) => null);

    return response;
  }
}
