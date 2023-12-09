import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:location/location.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/services/cmd/auth/local/local_auth_service.dart';
import 'package:now_v8/src/services/cmd/colors_service/colors_service.dart';
import 'package:now_v8/src/services/cmd/filter_service/fake/filterFakeService.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/filterService.dart';
import 'package:now_v8/src/services/cmd/location_service/fake/locationFakeService.dart';
import 'package:now_v8/src/services/cmd/location_service/service/locationService.dart';
import 'package:now_v8/src/services/cmd/storage/key_value/local_hive_storage.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

final locationServiceProvider = Provider<ILocationService>((ref) {
  return LocationFakeService();
});

final filterServiceProvider = Provider<IFilterService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return FilterService(apiConfig: apiConfig);
});

final fakeFilterServiceProvider = Provider<IFilterService>((ref) {
  return FilterFakeService();
});

final colorsServiceProvider = Provider<ColorsService>(
  (ref) => ColorsService(),
);

final apiConfigProvider = Provider<ApiConfig>(
  (ref) => ApiConfig.toProd(),
);

final keyValueProvider = Provider.family<IKeyValueStorage, String>(
  (ref, databaseName) {
    return HiveKeyValue<String>(boxName: databaseName);
  },
);

final authProvider = Provider<IAuthService>((ref) {
  var keyValue = ref.read(keyValueProvider("authSession"));
  return AuthLocalStorage(keyValueStorage: keyValue);
});
