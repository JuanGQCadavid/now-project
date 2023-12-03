import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/services/cmd/user_service/fake/service.dart';
import 'package:now_v8/src/services/cmd/user_service/service/service.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

class LoginStateNotifer extends StateNotifier<LoginState> {
  final IUserService userService = UserService(
    apiConfig: ApiConfig.toProd(),
  );

  LoginStateNotifer()
      : super(
          const LoginState(
            stateConfig: OnStateConfig(
              showPhoneNumber: true,
              showCodeInput: false,
              showUserName: false,
            ),
            errorMessage: "",
            phoneNumber: "",
            userName: "",
            onState: OnState.onInit,
          ),
        );

  void onNext(
      String userPhoneNumber, String userName, List<String> userCode) async {
    print(
        "\n userPhoneNumber: ${userPhoneNumber} \n userName: ${userName}\n userCode: ${userCode}\n State: ${state.onState}");

    if (state.onState == OnState.onInit) {
      onUserAttemptToLogIn(userPhoneNumber);
    } else if (state.onState == OnState.onLogin) {
      validateLogin(userPhoneNumber, userCode);
    } else if (state.onState == OnState.onSingUp) {
      preSignUser(userPhoneNumber, userCode);
    } else if (state.onState == OnState.onSingUpPhoneValidation) {
      validateSignUser(userPhoneNumber, userCode, userName);
    }
  }

  void validateSignUser(
    String userPhoneNumber,
    List<String> userCode,
    String userName,
  ) {
    state = state.copyWith(
      onState: OnState.onDone,
      stateConfig: stateConfigMaps[OnState.onDone]!,
    );
  }

  void preSignUser(
    String userPhoneNumber,
    List<String> userCode,
  ) async {
    state = state.copyWith(
      onState: OnState.onSingUpPhoneValidation,
      stateConfig: stateConfigMaps[OnState.onSingUpPhoneValidation]!,
    );
  }

  void validateLogin(
    String userPhoneNumber,
    List<String> userCode,
  ) async {
    state = state.copyWith(
      onState: OnState.onDone,
      stateConfig: stateConfigMaps[OnState.onDone]!,
    );
  }

  // HERE
  void onUserAttemptToLogIn(String userPhoneNumber) async {
    var serviceResponse = await userService.login(userPhoneNumber);

    serviceResponse.fold(
      (l) => {
        state = state.copyWith(
          onState: OnState.onLogin,
          phoneNumber: userPhoneNumber,
          stateConfig: stateConfigMaps[OnState.onLogin]!,
        )
      },
      (r) => {
        r.whenOrNull(
          userDoesNotExist: () {
            print("userDoesNotExist");
            state = state.copyWith(
              onState: OnState.onSingUp,
              phoneNumber: userPhoneNumber,
              stateConfig: stateConfigMaps[OnState.onSingUp]!,
            );
          },
          internalError: (String err) {
            print("internal err - " + err);
            state = state.copyWith(
              errorMessage: err,
              phoneNumber: userPhoneNumber,
            );
          },
        )
      },
    );
  }
}
