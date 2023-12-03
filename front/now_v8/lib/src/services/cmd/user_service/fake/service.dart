import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/user.dart';

// class FakeUserService implements IUserService {
//   final String exist = "+57301";
//   final UserDetails existUser =
//       UserDetails(userId: "123", userName: "Juan", userToken: "myFuckingToken");
//   final String doesNotExist = "+57323";

//   @override
//   Future<Either<None, UserError>> login(String userPhoneNumber) async {
//     if (userPhoneNumber == exist) {
//       return left(None());
//     } else if (userPhoneNumber == doesNotExist) {
//       return right(UserError.userDoesNotExist());
//     } else {
//       return right(UserError.internalError(
//           "phone number is not the two options ${userPhoneNumber}"));
//     }
//   }
// }
