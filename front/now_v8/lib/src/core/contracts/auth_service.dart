import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/user.dart';

abstract class IAuthService {
  Either<UserDetails, None> getUserDetails();
  storeUserDetails(UserDetails user);
}
