import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/models/user.dart';

class AuthLocalStorage implements IAuthService {
  final IKeyValueStorage<dynamic, dynamic> keyValueStorage;
  final String key = "UserDetails";

  AuthLocalStorage({required this.keyValueStorage}) {
    keyValueStorage.doInit().then((value) => {print("We are init!")});
  }

  @override
  Future<Either<UserDetails, None>> getUserDetails() async {
    await keyValueStorage.doInit();

    var user = keyValueStorage.getValue(key);

    if (user.isRight()) {
      return right(const None());
    }

    return user.fold((l) {
      var data = Map<String, dynamic>.from(l as Map);
      var user = UserDetails.fromJson(data);
      return left(user);
    }, (r) {
      return right(const None());
    });
  }

  @override
  Future storeUserDetails(UserDetails user) async {
    await keyValueStorage.doInit();

    var userDumped = user.toJson();
    keyValueStorage.save(userDumped, key);
  }

  @override
  Future removeUserDetails() async {
    await keyValueStorage.doInit();

    await keyValueStorage.delete(key);
  }
}
