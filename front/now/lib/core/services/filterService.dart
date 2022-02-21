import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:now/core/domain/errors/backend_errors.dart';
import 'package:now/core/domain/models/spot.dart';
import 'package:now/core/domain/ports/i_now_filter_service.dart';

import 'package:http/http.dart' as http;

class FilterService implements INowFIlterService {
  final String FILTER_URL =
      "https://4co5utcub8.execute-api.us-east-2.amazonaws.com/prod";

  @override
  Future<Either<BackendErrors, Locations>> fetchByProximity({
    required double lat,
    required double lon,
    double radious = 0.5,
  }) async {
    final uri = Uri.parse('$FILTER_URL/filter/proximity?cpLat=$lat&cpLon=$lon');
    print(uri);
    print("\n+++++++++++");

    final closestSpots = await http.get(uri);

    print(closestSpots.statusCode);
    print("\n+++++++++++");
    print(closestSpots.body);
    print("\n------------");

    if (closestSpots.statusCode >= 200 && closestSpots.statusCode < 400) {
      final locations = Locations.fromJson(json.decode(closestSpots.body));

      print(locations.places.first.eventInfo.name);
      print(locations.places.first.eventInfo.emoji);
      print(locations.places.first.placeInfo.lat);
      print(locations.places.first.placeInfo.lon);
      return right(locations);
    }

    return left(const BackendErrors.internalError());
  }
}
