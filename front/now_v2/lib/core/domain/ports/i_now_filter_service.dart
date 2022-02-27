import 'package:dartz/dartz.dart';

import '../errors/backend_errors.dart';
import '../models/spot.dart';

abstract class INowFIlterService {
  Future<Either<BackendErrors, Locations>> fetchByProximity({
    required double lat,
    required double lon,
    double radious = 0.5,
  });
}
