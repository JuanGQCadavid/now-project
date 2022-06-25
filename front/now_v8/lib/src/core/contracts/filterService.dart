

import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/state_response.dart';

abstract class IFilterService {

  Future<List<Spot>> getByProximity({
    required double cpLat,
    required double cpLng,
    double radious = 10
  });

  Future<StateResponse<List<LongSpot>, String>> getByProximityWithState({
    required double cpLat,
    required double cpLng,
    double radious = 10,
    String token = "",
  });
}