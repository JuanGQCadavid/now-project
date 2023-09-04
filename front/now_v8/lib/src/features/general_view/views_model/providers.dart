// import 'dart:developer';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/features/general_view/model/generalViewModel.dart';
import 'package:now_v8/src/features/general_view/views_model/spotsStateNotifier.dart';
import 'package:now_v8/src/services/core/providers.dart';

final generalViewModelProvider = Provider<GeneralViewModel>((ref) {
  final IColorService colorsService = ref.read(colorsServiceProvider);
  final IFilterService filterService = ref.read(filterServiceProvider);
  final ILocationService locationService = ref.read(locationServiceProvider);

  return GeneralViewModel(
    colorService: colorsService,
    filterService: filterService,
    locationService: locationService,
  );
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

  return generalViewModel.filterSpotsBaseOnTags(tagsSelected, spots);
});
