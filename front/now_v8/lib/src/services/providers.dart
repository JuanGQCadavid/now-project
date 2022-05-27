import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/services/colors_service/colors_service.dart';
import 'package:now_v8/src/services/filter_service/filterService.dart';
import 'package:now_v8/src/services/location_service/locationService.dart';

final spotServiceProvider = Provider((ref) {
  return 0;
});

final locationServiceProvider = Provider<ILocationService>((ref) {
  return LocationService();
});

final filterServiceProvider = Provider<IFilterService>((ref) {
  return FilterService();
});

final colorsServiceProvider = Provider<ColorsService>((ref) => ColorsService());