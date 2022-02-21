import 'package:dartz/dartz.dart';
import 'package:now/core/domain/errors/backend_errors.dart';
import 'package:now/core/domain/models/spot.dart';

abstract class INowFIlterService {
  Future<Either<BackendErrors, Locations>> fetchByProximity({
    required double lat,
    required double lon,
    double radious = 0.5,
  });
}
