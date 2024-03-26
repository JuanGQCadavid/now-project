import 'package:dartz/dartz.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/token.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';

abstract class ISpotCoreService {
  Future<Either<LongSpot, BackendErrors>> createSpot(
    LongSpot spot,
    Token token,
  );
  Future<Either<LongSpot, BackendErrors>> fetchSpot(String spotID);
}
