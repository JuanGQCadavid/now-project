import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/profile.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';

abstract class IUserProfileService {
  Future<Either<UserProfile, BackendErrors>> getUserProfile(
    String userId,
    Token token,
  );
}
