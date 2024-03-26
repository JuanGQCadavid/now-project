import 'package:dartz/dartz.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:location/location.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/auth/local/local_auth_service.dart';
import 'package:now_v8/src/services/cmd/colors_service/colors_service.dart';
import 'package:now_v8/src/services/cmd/filter_service/fake/filterFakeService.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/filterService.dart';
import 'package:now_v8/src/services/cmd/gcp_service/fake/fake_google_cloud_services.dart';
import 'package:now_v8/src/services/cmd/gcp_service/services/google_cloud_services.dart';
import 'package:now_v8/src/services/cmd/location_service/fake/locationFakeService.dart';
import 'package:now_v8/src/services/cmd/location_service/service/locationService.dart';
import 'package:now_v8/src/services/cmd/spot_core_service/service/service.dart';
import 'package:now_v8/src/services/cmd/storage/key_value/local_hive_storage.dart';
import 'package:now_v8/src/services/cmd/user_service/service/service.dart';
import 'package:now_v8/src/services/core/notifiers.dart';
import 'package:now_v8/src/services/core/now_services_caller.dart';
import 'package:now_v8/src/services/core/services_api_configuration.dart';

//////////////////////////////////////////////////////////////////////////////
///
/// Internal Services
///
//////////////////////////////////////////////////////////////////////////////

final locationServiceProvider = Provider<ILocationService>((ref) {
  return LocationFakeService();
});

final colorsServiceProvider = Provider<ColorsService>(
  (ref) => ColorsService(),
);

final keyValueProvider = Provider.family<IKeyValueStorage, String>(
  (ref, databaseName) {
    return HiveKeyValue<String>(boxName: databaseName);
  },
);

final authProvider = Provider<IAuthService>((ref) {
  var keyValue = HiveKeyValue<Map<dynamic, dynamic>?>(boxName: "authSession");
  return AuthLocalStorage(keyValueStorage: keyValue);
});

final authStateProvider =
    StateNotifierProvider<AuthState, Either<UserDetails, None>>((ref) {
  final authService = ref.read(authProvider);
  return AuthState(authService: authService);
});

//////////////////////////////////////////////////////////////////////////////
///
/// Third Party Services
///
//////////////////////////////////////////////////////////////////////////////

final gpcServicesProvider = Provider<IGCPServices>((ref) {
  // return GoogleCloudServices();
  return FakeGoogleServices();
});

//////////////////////////////////////////////////////////////////////////////
///
/// Backend Services
///
//////////////////////////////////////////////////////////////////////////////

final apiConfigProvider = Provider<ApiConfig>(
  (ref) => ApiConfig.toProd(),
);

final userServiceProvider = Provider<IUserService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return UserService(
    apiConfig: apiConfig,
  );
});

final filterServiceProvider = Provider<IFilterService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return FilterService(apiConfig: apiConfig);
});

final spotsCoreSeriveProvider = Provider<ISpotCoreService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return SpotCoreService(
    apiConfig: apiConfig,
    caller: NowServicesCaller(
      baseUrl: apiConfig.getSpotCoreEndpoint(),
    ),
  );
});

final fakeFilterServiceProvider = Provider<IFilterService>((ref) {
  return FilterFakeService();
});
