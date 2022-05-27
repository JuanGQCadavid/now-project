

import 'package:now_v8/src/core/models/spot.dart';

abstract class IFilterService {
  List<Spot> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10
  });
}