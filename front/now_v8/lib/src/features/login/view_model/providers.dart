import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/features/login/view_model/state_notifier.dart';

final loginStateNotifierProvider =
    StateNotifierProvider<LoginStateNotifer, LoginState>(
        (ref) => LoginStateNotifer());
