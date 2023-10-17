import 'package:flutter/foundation.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/features/login/model/login_state.dart';

part 'state_notifier.freezed.dart';

class UserDetails {
  final String userName;
  final String userId;
  final String userToken;

  UserDetails({
    required this.userName,
    required this.userId,
    required this.userToken,
  });
}

@freezed
class UserError with _$UserError {
  factory UserError.userDoesNotExist() = UserDoesNotExist;
  factory UserError.internalError(String error) = InternalError;
}

abstract class IUserService {
  Future<Either<UserDetails, UserError>> login(String userPhoneNumber);
}

class FakeUserService implements IUserService {
  final String exist = "+57301";
  final UserDetails existUser =
      UserDetails(userId: "123", userName: "Juan", userToken: "myFuckingToken");
  final String doesNotExist = "+57323";

  @override
  Future<Either<UserDetails, UserError>> login(String userPhoneNumber) async {
    if (userPhoneNumber == exist) {
      return left(existUser);
    } else if (userPhoneNumber == doesNotExist) {
      return right(UserError.userDoesNotExist());
    } else {
      return right(UserError.internalError(
          "phone number is not the two options ${userPhoneNumber}"));
    }
  }
}

class LoginStateNotifer extends StateNotifier<LoginState> {
  final IUserService userService = FakeUserService();

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

  void onNext(String userPhoneNumber, String userName, String userCode) async {
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
    String userCode,
    String userName,
  ) {
    state = state.copyWith(
      onState: OnState.onDone,
    );
  }

  void preSignUser(
    String userPhoneNumber,
    String userCode,
  ) async {
    state = state.copyWith(
      onState: OnState.onSingUpPhoneValidation,
    );
  }

  void validateLogin(
    String userPhoneNumber,
    String userCode,
  ) async {
    state = state.copyWith(
      onState: OnState.onDone,
    );
  }

  void onUserAttemptToLogIn(String userPhoneNumber) async {
    var serviceResponse = await userService.login(userPhoneNumber);

    serviceResponse.fold(
      (l) => {
        state = state.copyWith(
          onState: OnState.onLogin,
          phoneNumber: userPhoneNumber,
        )
      },
      (r) => {
        r.when(
          userDoesNotExist: () {
            state = state.copyWith(
              onState: OnState.onSingUp,
              phoneNumber: userPhoneNumber,
            );
          },
          internalError: (String err) {
            state = state.copyWith(
              // onState: OnState.onErrorState,
              errorMessage: err,
              phoneNumber: userPhoneNumber,
            );
          },
        )
      },
    );
  }
}
