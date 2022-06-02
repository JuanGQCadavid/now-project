import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/services/cmd/colors_service/colors_service.dart';
import 'package:now_v8/src/services/cmd/filter_service/fake/filterFakeService.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/filterService.dart';
import 'package:now_v8/src/services/cmd/location_service/fake/locationFakeService.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

final spotServiceProvider = Provider((ref) {
  return 0;
});

final locationServiceProvider = Provider<ILocationService>((ref) {
  return LocationFakeService();
});

final filterServiceProvider = Provider<IFilterService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return FilterService(apiConfig: apiConfig);
});

final colorsServiceProvider = Provider<ColorsService>(
  (ref) => ColorsService(),
);

final apiConfigProvider = Provider<ApiConfig>(
  (ref) => ApiConfig.toProd(),
);
