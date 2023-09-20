// import 'dart:developer';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/features/general_view/model/general_view_model.dart';
import 'package:now_v8/src/features/general_view/views_model/spotsStateNotifier.dart';
import 'package:now_v8/src/services/core/providers.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';

final generalViewModelProvider = Provider<GeneralViewModel>((ref) {
  final IColorService colorsService = ref.read(colorsServiceProvider);
  final IFilterService filterService = ref.read(filterServiceProvider);
  final ILocationService locationService = ref.read(locationServiceProvider);
  final IKeyValueStorage keyValueStorage = ref.read(keyValueProvider("sessionDataGV"));

  return GeneralViewModel(
    colorService: colorsService,
    filterService: filterService,
    locationService: locationService,
    sessionDatabase: keyValueStorage
  );
});

final mapInteractionProvider = StateNotifierProvider<MapInteractions, MapState>((ref) => MapInteractions());

final mapSpotsBrigde = Provider((ref) {
  ref.listen<MapState>(mapInteractionProvider, (MapState? previousState, MapState newState) {
    if(previousState!.status == MapStatus.movingOnMap()  && newState!.status == MapStatus.movingIdle()){
      print("Refreshing!!!!!!!!!!!!!!!!!!!!!!!");
      var notifier = ref.read(spotsStateProvider.notifier);

      notifier.refreshSpots(latLng: newState.lastPositionKnowed);
    }
  });
  // ...
});


final spotsStateProvider = StateNotifierProvider<SpotsNotifer, List<Spot>>(
  ((ref) {

    final generalViewModel = ref.read(generalViewModelProvider);

    return SpotsNotifer(
      generalViewModel: generalViewModel,
    );
  }),
);

final tagsSelectedProvider =
    StateNotifierProvider<TagsNotifier, Set<String>>(((ref) {
  return TagsNotifier();
}));



final filteredSpotsProvider = StateProvider<FilteredSpots>((ref) {
  final tagsSelected = ref.watch(tagsSelectedProvider);
  final generalViewModel = ref.read(generalViewModelProvider);
  final spots = ref.watch(spotsStateProvider);
  ref.watch(mapSpotsBrigde);

  return generalViewModel.filterSpotsBaseOnTags(tagsSelected, spots);
});
