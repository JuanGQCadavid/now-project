import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';

class GeneralViewModel {
  final IFilterService filterService;
  final IColorService colorService;
  final ILocationService locationService;
  late SpotsColors defaultColor;

  List<Spot> spots = List.empty();

  GeneralViewModel(
      {required this.filterService,
      required this.colorService,
      required this.locationService}) {
    defaultColor = colorService.getColor();
  }

  Future<List<Spot>> getSpots() async {
    print("getSpots");
    LatLng userLocation = await locationService.getUserCurrentLocation();
    print("userLocation done");

    spots = await filterService.getByProximity(
        cpLat: userLocation.latitude, 
        cpLng: userLocation.longitude
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

    spots.forEach((spot) {
      bool spotToShow = false;

      if (tagsSelected.contains(spot.principalTag)) {
        spotToShow = true;
      } else {
        result.tagsOff.add(spot.principalTag);
      }

      spot.secondaryTags.forEach((spotTag) {
        if (tagsSelected.contains(spotTag)) {
          spotToShow = true;
        } else {
          result.tagsOff.add(spotTag);
        }
      });

      if (spotToShow) {
        result.spots.add(spot);
      }
    });

    return result;
  }


}
