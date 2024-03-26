import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/features/login/view_model/state_notifier.dart';
import 'package:now_v8/src/services/core/providers.dart';

final loginStateNotifierProvider =
    StateNotifierProvider.autoDispose<LoginStateNotifer, LoginState>((ref) {
  var auth = ref.read(authStateProvider.notifier);
  var userService = ref.read(userServiceProvider);
  return LoginStateNotifer(
    auth: auth,
    userService: userService,
  );
});
