import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/state_response.dart';

class GeneralViewModel {
  final IFilterService filterService;
  final IColorService colorService;
  final ILocationService locationService;
  final IKeyValueStorage sessionDatabase;

  late SpotsColors defaultColor;

  List<Spot> spots = List.empty();

  final String searchSessionKey = "Generla-View-Session-Id";

  GeneralViewModel(
      {required this.filterService,
      required this.colorService,
      required this.locationService,
      required this.sessionDatabase,
    }) {
    defaultColor = colorService.getColor();
  }

  Future<List<Spot>> getSpots([LatLng? centralPosition]) async {
    LatLng searchPosition = centralPosition??  await locationService.getUserCurrentLocation();

    String token = sessionDatabase.getValue(searchSessionKey);

    StateResponse<List<LongSpot>, String> filterResponse =
        await filterService.getByProximityWithState(
            cpLat: searchPosition.latitude, cpLng: searchPosition.longitude, token: token);
      

    spots = await filterService.getByProximity(
        cpLat: searchPosition.latitude, 
        cpLng: searchPosition.longitude
      );

    for (var spot in spots) {
      if (spot.date.compareTo(DateTime.now()) > 0){
        spot.spotsColor = colorService.getScheduleColor();
      }else {
        spot.spotsColor = colorService.getColor();
      }
      
    }

    return spots;
  }

  /// 1. Check if tagsOn is empty, if so just return all spots.
  /// 2. To filter
  ///  1. Check if the spot princiapl tag is on tags on, if so then mark flag as true, if not then add the tag as not showed.
  ///  2. Check if the spot secondary tags contains one of the tags on, if so then mark flag as true, if not then add the tag as not showed.
  ///  3. if the mark if flagged if so then add the spot.
  
  FilteredSpots filterSpotsBaseOnTags(
      Set<String> tagsSelected, 
      List<Spot> spots
  ) {
    if (tagsSelected.isEmpty) {
      return FilteredSpots(
          spots: spots,
          tagsOff: {},
          tagsSelected: tagsSelected,
          onFilterColor: const SpotsColors.empty());
    }

    FilteredSpots result = FilteredSpots.empty();
    result.onFilterColor = defaultColor;
    result.tagsSelected = tagsSelected;

    for (var spot in spots) {
      bool spotToShow = false;

      if (tagsSelected.contains(spot.principalTag)) {
        spotToShow = true;
      } else {
        result.tagsOff.add(spot.principalTag);
      }

      
      for (var spotTag in spot.secondaryTags) {
        if (tagsSelected.contains(spotTag)) {
          spotToShow = true;
        } else {
          result.tagsOff.add(spotTag);
        }
      }

      if (spotToShow) {
        result.spots.add(spot);
      }
    }

    return result;
  }


}
