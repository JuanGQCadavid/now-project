import 'package:dartz/dartz.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/gcp_services.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/contracts/location_service.dart';
import 'package:now_v8/src/core/contracts/profile_service.dart';
import 'package:now_v8/src/core/contracts/spot_core_service.dart';
import 'package:now_v8/src/core/contracts/user_service.dart';
import 'package:now_v8/src/core/models/profile.dart';
import 'package:now_v8/src/core/models/user.dart';
import 'package:now_v8/src/services/cmd/auth/local/local_auth_service.dart';
import 'package:now_v8/src/services/cmd/colors_service/colors_service.dart';
import 'package:now_v8/src/services/cmd/filter_service/service/filter_service.dart';
import 'package:now_v8/src/services/cmd/gcp_service/fake/fake_google_cloud_services.dart';
import 'package:now_v8/src/services/cmd/location_service/fake/location_fake_service.dart';
import 'package:now_v8/src/services/cmd/profile_service/service/service.dart';
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

final userProfileStateProvider =
    StateNotifierProvider<UserProfileState, Either<UserProfile, None>>(
  (ref) {
    final authstateNotifier = ref.watch(authStateProvider.notifier);
    final userDetails = ref.watch(authStateProvider);
    final userProfileService = ref.read(userProfileServiceProvider);

    return UserProfileState(
      authState: authstateNotifier,
      userDetails: userDetails,
      userProfileService: userProfileService,
    );
  },
);

final authProvider = Provider<IAuthService>((ref) {
  var keyValue = HiveKeyValue<Map<dynamic, dynamic>?>(boxName: "authSession");
  return AuthLocalStorage(keyValueStorage: keyValue);
});

final authStateProvider =
    StateNotifierProvider<AuthState, Either<UserDetails, None>>((ref) {
  final authService = ref.read(authProvider);
  final userService = ref.read(userServiceProvider);
  return AuthState(
    authService: authService,
    userService: userService,
  );
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
  (ref) => ApiConfig.toStaging(),
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

final userProfileServiceProvider = Provider<IUserProfileService>((ref) {
  final ApiConfig apiConfig = ref.read(apiConfigProvider);
  return UserProfileService(
    apiConfig: apiConfig,
    caller: NowServicesCaller(
      baseUrl: apiConfig.getUserProfileEndpoint(),
    ),
  );
});
