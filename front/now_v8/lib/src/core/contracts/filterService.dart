

import 'package:now_v8/src/core/models/spot.dart';

abstract class IFilterService {

  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10
  });
}