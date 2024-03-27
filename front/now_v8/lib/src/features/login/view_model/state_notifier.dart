import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';
import 'package:now_v8/src/services/core/notifiers.dart';

class LoginStateNotifer extends StateNotifier<LoginState> {
  final AuthState auth;

  LoginStateNotifer({
    required this.auth,
  }) : super(
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

    var mapStates = {
      OnState.onInit: () {
        onUserAttemptToLogIn(userPhoneNumber);
      },
      OnState.onLogin: () {
        validate(userPhoneNumber, userCode);
      },
      OnState.onSingUp: () {
        initSignUp(userPhoneNumber, userName);
      },
      OnState.onSingUpPhoneValidation: () {
        validate(userPhoneNumber, userCode);
      },
    };

    var a = mapStates[state.onState];
    a!();
  }

  void initSignUp(
    String userPhoneNumber,
    String userName,
  ) async {
    var serviceResponse = await auth.initSingUp(userPhoneNumber, userName);

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
    var serviceResponse = await auth.validate(
      state.userName,
      userPhoneNumber,
      userCode,
    );

    serviceResponse.fold(
      (l) {
        state = state.copyWith(
          onState: OnState.onDone,
        );
      },
      (r) => r.whenOrNull(
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
      ),
    );
  }

  void onUserAttemptToLogIn(String userPhoneNumber) async {
    var serviceResponse = await auth.initLoging(userPhoneNumber);

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
