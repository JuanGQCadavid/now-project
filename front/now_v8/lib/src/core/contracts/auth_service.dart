import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/user.dart';

abstract class IAuthService {
  Future<Either<UserDetails, None>> getUserDetails();
  Future storeUserDetails(UserDetails user);
}
