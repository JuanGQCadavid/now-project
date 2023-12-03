import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:flutter/foundation.dart';

part 'login_state.freezed.dart';

enum OnState {
  onLogin,
  onSingUp,
  onInit,
  onSingUpPhoneValidation,
  onErrorState,
  onDone,
}

@freezed
class OnStateConfig with _$OnStateConfig {
  const factory OnStateConfig({
    required bool showPhoneNumber,
    required bool showCodeInput,
    required bool showUserName,
  }) = _OnStateConfig;
}

@freezed
class LoginState with _$LoginState {
  const factory LoginState({
    required String phoneNumber,
    required String userName,
    required OnState onState,
    required String errorMessage,
    required OnStateConfig stateConfig,
  }) = _LoginState;
}

Map<OnState, OnStateConfig> stateConfigMaps = {
  OnState.onInit: const OnStateConfig(
    showCodeInput: false,
    showPhoneNumber: true,
    showUserName: false,
  ),
  OnState.onLogin: const OnStateConfig(
    showCodeInput: true,
    showPhoneNumber: true,
    showUserName: false,
  ),
  OnState.onSingUp: const OnStateConfig(
    showCodeInput: false,
    showPhoneNumber: true,
    showUserName: true,
  ),
  OnState.onSingUpPhoneValidation: const OnStateConfig(
    showCodeInput: true,
    showPhoneNumber: true,
    showUserName: true,
  ),
  OnState.onErrorState: const OnStateConfig(
    showCodeInput: false,
    showPhoneNumber: false,
    showUserName: false,
  ),
  OnState.onDone: const OnStateConfig(
    showCodeInput: false,
    showPhoneNumber: false,
    showUserName: false,
  ),
};
