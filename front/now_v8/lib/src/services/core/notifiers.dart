import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/models/user.dart';

class OnAuthState extends StateNotifier<Either<UserDetails, None>> {
  final IAuthService authService;
  OnAuthState({required this.authService}) : super(right(const None()));

  Future userLogIn(UserDetails details) async {
    await authService.storeUserDetails(details);
    state = left(details);
  }

  Future userLogOut() async {
    await authService.removeUserDetails();
    state = right(const None());
  }
}
