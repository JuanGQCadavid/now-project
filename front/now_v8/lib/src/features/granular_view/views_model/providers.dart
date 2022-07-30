import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/features/granular_view/model/ganular_model.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';
import 'package:now_v8/src/features/granular_view/views_model/state_notifiers.dart';
import 'package:now_v8/src/services/core/providers.dart';
// fakeFilterServiceProvider
// filterServiceProvider
final granularModelProvider = Provider<GranularModel>((ref) {
  final IFilterService filterService = ref.read(fakeFilterServiceProvider);
  final ILocationService locationService = ref.read(locationServiceProvider);

  return GranularModel(
      filterService: filterService, locationService: locationService);
});

final detailedSpotProvider =
    StateNotifierProvider<DetailedSpotsState, List<LongSpot>>((ref) {
  final granularModel = ref.read(granularModelProvider);

  return DetailedSpotsState(
    granularModel: granularModel,
  );
});

final onSpotProvider = StateNotifierProvider<OnSpotState, GranularSpot>((ref) {
  final granularModel = ref.read(granularModelProvider);
  final detailedSpot = ref.watch(detailedSpotProvider);

  return OnSpotState(granularModel: granularModel, actualList: detailedSpot);
});
