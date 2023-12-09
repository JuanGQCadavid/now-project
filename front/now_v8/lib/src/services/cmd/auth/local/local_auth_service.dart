import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/models/user.dart';

class AuthLocalStorage implements IAuthService {
  final IKeyValueStorage<dynamic, dynamic> keyValueStorage;
  final String key = "UserDetails";

  AuthLocalStorage({required this.keyValueStorage}) {
    keyValueStorage.doInit();
  }

  @override
  Either<UserDetails, None> getUserDetails() {
    var user = keyValueStorage.getValue(key);

    if (user.isRight()) {
      return right(const None());
    }

    return user.fold((l) {
      var user = l as UserDetails;
      return left(user);
    }, (r) {
      return right(const None());
    });
  }

  @override
  storeUserDetails(UserDetails user) {
    keyValueStorage.save(user, key);
  }
}
