import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/storage/key_value/local_hive_storage.dart';

class AuthLocalStorage implements IAuthService {
  final HiveKeyValue<UserDetails> keyValueStorage;
  final String key = "UserDetails";

  AuthLocalStorage({required this.keyValueStorage}) {
    keyValueStorage.doInit();
  }

  @override
  Either<UserDetails, None> getUserDetails() {
    var user = keyValueStorage.getValue(key);

    if (user == Null) {
      return right(const None());
    }

    return left(user);
  }

  @override
  storeUserDetails(UserDetails user) {
    keyValueStorage.save(user, key);
  }
}
