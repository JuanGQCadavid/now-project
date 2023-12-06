import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';

part 'user_service.freezed.dart';

abstract class IUserService {
  Future<Either<None, UserError>> initLoging(String userPhoneNumber);
  Future<Either<None, UserError>> initSingUp(
    String userPhoneNumber,
    String userName,
  );
  Future<Either<UserDetails, UserError>> validate(
    String userPhoneNumber,
    List<String> userCode,
  );
}

@freezed
class UserError with _$UserError {
  factory UserError.userDoesNotExist() = UserDoesNotExist;
  factory UserError.phoneNumberAlreadyTaken() = phoneNumberAlreadyTaken;
  factory UserError.otpAlive() = otpAlive;
  factory UserError.otpDied() = otpDied;
  factory UserError.wrongOTP() = wrongOTP;
  factory UserError.noPendingOTP() = noPendingOTP;
  factory UserError.otpMaxTriesReached() = otpMaxTriesReached;

  factory UserError.internalError(String error) = InternalError;
}

UserError mapMessageErrorToUserError(ErrorMessage errorMessage) {
  switch (errorMessage.id) {
    case "UserNotFound":
      return UserError.userDoesNotExist();
    case "Internal":
      return UserError.internalError(errorMessage.internalError);
    case "PhoneNumberAlreadyTaken":
      return UserError.phoneNumberAlreadyTaken();
    case "OTPAlive":
      return UserError.otpAlive();
    case "OTPDied":
      return UserError.otpDied();
    case "WrongOTP":
      return UserError.wrongOTP();
    case "NoPendingOTP":
      return UserError.noPendingOTP();
    case "OTPMaxTriesReached":
      return UserError.otpMaxTriesReached();
    default:
      return UserError.internalError(errorMessage.message);
  }
}
