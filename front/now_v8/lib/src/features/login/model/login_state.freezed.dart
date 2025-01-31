// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'login_state.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#adding-getters-and-methods-to-our-models');

/// @nodoc
mixin _$OnStateConfig {
  bool get showPhoneNumber => throw _privateConstructorUsedError;
  bool get showCodeInput => throw _privateConstructorUsedError;
  bool get showUserName => throw _privateConstructorUsedError;

  /// Create a copy of OnStateConfig
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $OnStateConfigCopyWith<OnStateConfig> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $OnStateConfigCopyWith<$Res> {
  factory $OnStateConfigCopyWith(
          OnStateConfig value, $Res Function(OnStateConfig) then) =
      _$OnStateConfigCopyWithImpl<$Res, OnStateConfig>;
  @useResult
  $Res call({bool showPhoneNumber, bool showCodeInput, bool showUserName});
}

/// @nodoc
class _$OnStateConfigCopyWithImpl<$Res, $Val extends OnStateConfig>
    implements $OnStateConfigCopyWith<$Res> {
  _$OnStateConfigCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of OnStateConfig
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? showPhoneNumber = null,
    Object? showCodeInput = null,
    Object? showUserName = null,
  }) {
    return _then(_value.copyWith(
      showPhoneNumber: null == showPhoneNumber
          ? _value.showPhoneNumber
          : showPhoneNumber // ignore: cast_nullable_to_non_nullable
              as bool,
      showCodeInput: null == showCodeInput
          ? _value.showCodeInput
          : showCodeInput // ignore: cast_nullable_to_non_nullable
              as bool,
      showUserName: null == showUserName
          ? _value.showUserName
          : showUserName // ignore: cast_nullable_to_non_nullable
              as bool,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$OnStateConfigImplCopyWith<$Res>
    implements $OnStateConfigCopyWith<$Res> {
  factory _$$OnStateConfigImplCopyWith(
          _$OnStateConfigImpl value, $Res Function(_$OnStateConfigImpl) then) =
      __$$OnStateConfigImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({bool showPhoneNumber, bool showCodeInput, bool showUserName});
}

/// @nodoc
class __$$OnStateConfigImplCopyWithImpl<$Res>
    extends _$OnStateConfigCopyWithImpl<$Res, _$OnStateConfigImpl>
    implements _$$OnStateConfigImplCopyWith<$Res> {
  __$$OnStateConfigImplCopyWithImpl(
      _$OnStateConfigImpl _value, $Res Function(_$OnStateConfigImpl) _then)
      : super(_value, _then);

  /// Create a copy of OnStateConfig
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? showPhoneNumber = null,
    Object? showCodeInput = null,
    Object? showUserName = null,
  }) {
    return _then(_$OnStateConfigImpl(
      showPhoneNumber: null == showPhoneNumber
          ? _value.showPhoneNumber
          : showPhoneNumber // ignore: cast_nullable_to_non_nullable
              as bool,
      showCodeInput: null == showCodeInput
          ? _value.showCodeInput
          : showCodeInput // ignore: cast_nullable_to_non_nullable
              as bool,
      showUserName: null == showUserName
          ? _value.showUserName
          : showUserName // ignore: cast_nullable_to_non_nullable
              as bool,
    ));
  }
}

/// @nodoc

class _$OnStateConfigImpl
    with DiagnosticableTreeMixin
    implements _OnStateConfig {
  const _$OnStateConfigImpl(
      {required this.showPhoneNumber,
      required this.showCodeInput,
      required this.showUserName});

  @override
  final bool showPhoneNumber;
  @override
  final bool showCodeInput;
  @override
  final bool showUserName;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'OnStateConfig(showPhoneNumber: $showPhoneNumber, showCodeInput: $showCodeInput, showUserName: $showUserName)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'OnStateConfig'))
      ..add(DiagnosticsProperty('showPhoneNumber', showPhoneNumber))
      ..add(DiagnosticsProperty('showCodeInput', showCodeInput))
      ..add(DiagnosticsProperty('showUserName', showUserName));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$OnStateConfigImpl &&
            (identical(other.showPhoneNumber, showPhoneNumber) ||
                other.showPhoneNumber == showPhoneNumber) &&
            (identical(other.showCodeInput, showCodeInput) ||
                other.showCodeInput == showCodeInput) &&
            (identical(other.showUserName, showUserName) ||
                other.showUserName == showUserName));
  }

  @override
  int get hashCode =>
      Object.hash(runtimeType, showPhoneNumber, showCodeInput, showUserName);

  /// Create a copy of OnStateConfig
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$OnStateConfigImplCopyWith<_$OnStateConfigImpl> get copyWith =>
      __$$OnStateConfigImplCopyWithImpl<_$OnStateConfigImpl>(this, _$identity);
}

abstract class _OnStateConfig implements OnStateConfig {
  const factory _OnStateConfig(
      {required final bool showPhoneNumber,
      required final bool showCodeInput,
      required final bool showUserName}) = _$OnStateConfigImpl;

  @override
  bool get showPhoneNumber;
  @override
  bool get showCodeInput;
  @override
  bool get showUserName;

  /// Create a copy of OnStateConfig
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$OnStateConfigImplCopyWith<_$OnStateConfigImpl> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
mixin _$LoginState {
  String get phoneNumber => throw _privateConstructorUsedError;
  String get userName => throw _privateConstructorUsedError;
  OnState get onState => throw _privateConstructorUsedError;
  String get errorMessage => throw _privateConstructorUsedError;
  OnStateConfig get stateConfig => throw _privateConstructorUsedError;

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  $LoginStateCopyWith<LoginState> get copyWith =>
      throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $LoginStateCopyWith<$Res> {
  factory $LoginStateCopyWith(
          LoginState value, $Res Function(LoginState) then) =
      _$LoginStateCopyWithImpl<$Res, LoginState>;
  @useResult
  $Res call(
      {String phoneNumber,
      String userName,
      OnState onState,
      String errorMessage,
      OnStateConfig stateConfig});

  $OnStateConfigCopyWith<$Res> get stateConfig;
}

/// @nodoc
class _$LoginStateCopyWithImpl<$Res, $Val extends LoginState>
    implements $LoginStateCopyWith<$Res> {
  _$LoginStateCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? phoneNumber = null,
    Object? userName = null,
    Object? onState = null,
    Object? errorMessage = null,
    Object? stateConfig = null,
  }) {
    return _then(_value.copyWith(
      phoneNumber: null == phoneNumber
          ? _value.phoneNumber
          : phoneNumber // ignore: cast_nullable_to_non_nullable
              as String,
      userName: null == userName
          ? _value.userName
          : userName // ignore: cast_nullable_to_non_nullable
              as String,
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as OnState,
      errorMessage: null == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String,
      stateConfig: null == stateConfig
          ? _value.stateConfig
          : stateConfig // ignore: cast_nullable_to_non_nullable
              as OnStateConfig,
    ) as $Val);
  }

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @pragma('vm:prefer-inline')
  $OnStateConfigCopyWith<$Res> get stateConfig {
    return $OnStateConfigCopyWith<$Res>(_value.stateConfig, (value) {
      return _then(_value.copyWith(stateConfig: value) as $Val);
    });
  }
}

/// @nodoc
abstract class _$$LoginStateImplCopyWith<$Res>
    implements $LoginStateCopyWith<$Res> {
  factory _$$LoginStateImplCopyWith(
          _$LoginStateImpl value, $Res Function(_$LoginStateImpl) then) =
      __$$LoginStateImplCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call(
      {String phoneNumber,
      String userName,
      OnState onState,
      String errorMessage,
      OnStateConfig stateConfig});

  @override
  $OnStateConfigCopyWith<$Res> get stateConfig;
}

/// @nodoc
class __$$LoginStateImplCopyWithImpl<$Res>
    extends _$LoginStateCopyWithImpl<$Res, _$LoginStateImpl>
    implements _$$LoginStateImplCopyWith<$Res> {
  __$$LoginStateImplCopyWithImpl(
      _$LoginStateImpl _value, $Res Function(_$LoginStateImpl) _then)
      : super(_value, _then);

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? phoneNumber = null,
    Object? userName = null,
    Object? onState = null,
    Object? errorMessage = null,
    Object? stateConfig = null,
  }) {
    return _then(_$LoginStateImpl(
      phoneNumber: null == phoneNumber
          ? _value.phoneNumber
          : phoneNumber // ignore: cast_nullable_to_non_nullable
              as String,
      userName: null == userName
          ? _value.userName
          : userName // ignore: cast_nullable_to_non_nullable
              as String,
      onState: null == onState
          ? _value.onState
          : onState // ignore: cast_nullable_to_non_nullable
              as OnState,
      errorMessage: null == errorMessage
          ? _value.errorMessage
          : errorMessage // ignore: cast_nullable_to_non_nullable
              as String,
      stateConfig: null == stateConfig
          ? _value.stateConfig
          : stateConfig // ignore: cast_nullable_to_non_nullable
              as OnStateConfig,
    ));
  }
}

/// @nodoc

class _$LoginStateImpl with DiagnosticableTreeMixin implements _LoginState {
  const _$LoginStateImpl(
      {required this.phoneNumber,
      required this.userName,
      required this.onState,
      required this.errorMessage,
      required this.stateConfig});

  @override
  final String phoneNumber;
  @override
  final String userName;
  @override
  final OnState onState;
  @override
  final String errorMessage;
  @override
  final OnStateConfig stateConfig;

  @override
  String toString({DiagnosticLevel minLevel = DiagnosticLevel.info}) {
    return 'LoginState(phoneNumber: $phoneNumber, userName: $userName, onState: $onState, errorMessage: $errorMessage, stateConfig: $stateConfig)';
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty('type', 'LoginState'))
      ..add(DiagnosticsProperty('phoneNumber', phoneNumber))
      ..add(DiagnosticsProperty('userName', userName))
      ..add(DiagnosticsProperty('onState', onState))
      ..add(DiagnosticsProperty('errorMessage', errorMessage))
      ..add(DiagnosticsProperty('stateConfig', stateConfig));
  }

  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$LoginStateImpl &&
            (identical(other.phoneNumber, phoneNumber) ||
                other.phoneNumber == phoneNumber) &&
            (identical(other.userName, userName) ||
                other.userName == userName) &&
            (identical(other.onState, onState) || other.onState == onState) &&
            (identical(other.errorMessage, errorMessage) ||
                other.errorMessage == errorMessage) &&
            (identical(other.stateConfig, stateConfig) ||
                other.stateConfig == stateConfig));
  }

  @override
  int get hashCode => Object.hash(
      runtimeType, phoneNumber, userName, onState, errorMessage, stateConfig);

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @JsonKey(includeFromJson: false, includeToJson: false)
  @override
  @pragma('vm:prefer-inline')
  _$$LoginStateImplCopyWith<_$LoginStateImpl> get copyWith =>
      __$$LoginStateImplCopyWithImpl<_$LoginStateImpl>(this, _$identity);
}

abstract class _LoginState implements LoginState {
  const factory _LoginState(
      {required final String phoneNumber,
      required final String userName,
      required final OnState onState,
      required final String errorMessage,
      required final OnStateConfig stateConfig}) = _$LoginStateImpl;

  @override
  String get phoneNumber;
  @override
  String get userName;
  @override
  OnState get onState;
  @override
  String get errorMessage;
  @override
  OnStateConfig get stateConfig;

  /// Create a copy of LoginState
  /// with the given fields replaced by the non-null parameter values.
  @override
  @JsonKey(includeFromJson: false, includeToJson: false)
  _$$LoginStateImplCopyWith<_$LoginStateImpl> get copyWith =>
      throw _privateConstructorUsedError;
}
