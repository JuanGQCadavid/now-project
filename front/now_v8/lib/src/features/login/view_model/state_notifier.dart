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
      validate(userPhoneNumber, userCode);
    } else if (state.onState == OnState.onSingUp) {
      initSignUp(userPhoneNumber, userName);
    } else if (state.onState == OnState.onSingUpPhoneValidation) {
      validate(userPhoneNumber, userCode);
    }
  }

  void initSignUp(
    String userPhoneNumber,
    String userName,
  ) async {
    var serviceResponse =
        await userService.initSingUp(userPhoneNumber, userName);

    serviceResponse.fold(
      (l) => {
        state = state.copyWith(
          onState: OnState.onSingUpPhoneValidation,
          userName: userName,
          errorMessage: "",
          stateConfig: stateConfigMaps[OnState.onSingUpPhoneValidation]!,
        )
      },
      (r) => {
        r.whenOrNull(
          phoneNumberAlreadyTaken: () {
            print("phoneNumberAlreadyTaken");
            state = state.copyWith(
              onState: OnState.onLogin,
              userName: userName,
              stateConfig: stateConfigMaps[OnState.onLogin]!,
            );
          },
          otpAlive: () {
            print("otpAlive");
            state = state.copyWith(
              onState: OnState.onSingUp,
              userName: userName,
              errorMessage:
                  "You should wait some time, someone is already validating your account",
              stateConfig: stateConfigMaps[OnState.onSingUp]!,
            );
          },
          internalError: (String err) {
            print("internal err - " + err);
            state = state.copyWith(
              errorMessage: err,
              userName: userName,
            );
          },
        )
      },
    );

    state = state.copyWith(
      onState: OnState.onSingUpPhoneValidation,
      stateConfig: stateConfigMaps[OnState.onSingUpPhoneValidation]!,
    );
  }

  void validate(
    String userPhoneNumber,
    List<String> userCode,
  ) async {
    var serviceResponse = await userService.validate(userPhoneNumber, userCode);

    serviceResponse.fold(
      (l) => {
        print("ID: " +
            l.userId +
            "Tokens" +
            "\n" +
            l.refreshToken +
            "\n" +
            l.shortLiveToken +
            "\n" +
            l.shortLiveTokenTTL +
            "\n")
      },
      (r) => {
        r.whenOrNull(
          wrongOTP: () {
            print("wrongOTP");
            state = state.copyWith(
              errorMessage: "Wrong code, try again",
            );
          },
          otpMaxTriesReached: () {
            print("otpMaxTriesReached");
            state = state.copyWith(
              errorMessage: "Ups, we reach the limit",
            );
          },
          internalError: (String err) {
            print("internal err - " + err);
            state = state.copyWith(
              errorMessage: err,
            );
          },
        )
      },
    );
  }

  void onUserAttemptToLogIn(String userPhoneNumber) async {
    var serviceResponse = await userService.initLoging(userPhoneNumber);

    serviceResponse.fold(
      (l) => {
        state = state.copyWith(
          onState: OnState.onLogin,
          errorMessage: "",
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
