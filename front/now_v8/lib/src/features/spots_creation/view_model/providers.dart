import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/simple_state.dart';
import 'package:now_v8/src/features/spots_creation/model/core.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';
import 'package:now_v8/src/features/spots_creation/view_model/state_notifier.dart';
import 'package:now_v8/src/services/core/providers.dart';

final coreProvider = Provider<SpotsCreatorCore>((ref) {
  var gcpProvider = ref.read(gpcServicesProvider);
  var spotsCoreProvider = ref.read(spotsCoreSeriveProvider);
  var authState = ref.watch(authStateProvider.notifier);
  return SpotsCreatorCore(
    gpcService: gcpProvider,
    coreService: spotsCoreProvider,
    authState: authState,
  );
});

final spotsCreatorNotiferProvider =
    StateNotifierProvider.autoDispose<SpotCreator, SpotCreatorState>((ref) {
  var core = ref.read(coreProvider);
  return SpotCreator(
    core: core,
  );
});

final locationNotiferProvider =
    StateNotifierProvider<LocationState, SimpleState<PlaceInfo>>((ref) {
  var core = ref.read(coreProvider);
  var location = ref.read(locationServiceProvider);
  return LocationState(
    locationService: location,
    core: core,
  );
});

final tagNotifierProvider =
    StateNotifierProvider<TagsState, List<String>>((ref) => TagsState());
