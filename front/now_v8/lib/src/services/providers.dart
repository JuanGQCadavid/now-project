import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/services/colors_service/colors_service.dart';
import 'package:now_v8/src/services/filter_service/fake/filterFakeService.dart';
import 'package:now_v8/src/services/location_service/fake/locationFakeService.dart';

final spotServiceProvider = Provider((ref) {
  return 0;
});

final locationServiceProvider = Provider<ILocationService>((ref) {
  return LocationFakeService();
});

final filterServiceProvider = Provider<IFilterService>((ref) {
  return FilterFakeService();
});

final colorsServiceProvider = Provider<ColorsService>((ref) => ColorsService());