import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/user_service/service/service.dart';

class AuthState extends StateNotifier<Either<UserDetails, None>> {
  final IAuthService authService;
  // final IUserService userService;

  AuthState({required this.authService}) : super(right(const None())) {
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

  Future userLogIn(UserDetails details) async {
    await authService.storeUserDetails(details);
    state = left(details);
  }

  Future userLogOut() async {
    await authService.removeUserDetails();
    state = right(const None());
  }
}
