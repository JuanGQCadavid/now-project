import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/auth_service.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/location_service.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';
import 'package:now_v8/src/features/general_view/model/filteredSpots.dart';
import 'package:now_v8/src/core/contracts/key_value_storage.dart';
import 'package:now_v8/src/core/models/state_response.dart';

class GeneralViewModel {
  final IFilterService filterService;
  final IColorService colorService;
  final ILocationService locationService;
  final IKeyValueStorage sessionDatabase;
  late SpotsColors defaultColor;

  final String searchSessionKey = "Generla-View-Session-Id";
  final double threshold = 11;
  final double radious = 0.03;

  LatLng? lastPositionKnown;
  double? lastZoomKnown;

  GeneralViewModel({
    required this.filterService,
    required this.colorService,
    required this.locationService,
    required this.sessionDatabase,
  }) {
    defaultColor = colorService.getColor();
    sessionDatabase.doInit();
  }

  Future<String> getSessionToken() async {
    await sessionDatabase.doInit();
    var token = sessionDatabase.getValue(searchSessionKey);

    return token.fold(
      (l) {
        return l as String;
      },
      (r) {
        return "";
      },
    );
  }

  bool doJumpToDetails() {
    if (lastZoomKnown != null && lastZoomKnown! >= threshold) {
      lastPositionKnown = null;
      lastZoomKnown = null;
      return true;
    }
    return false;
  }

  Future<List<Spot>> getSpots({LatLng? centralPosition, double? zoom}) async {
    lastPositionKnown = centralPosition;
    lastZoomKnown = zoom;

    LatLng searchPosition =
        centralPosition ?? await locationService.getUserCurrentLocation();

    String token = await getSessionToken();

    StateResponse<List<Spot>, String> filterResponse =
        await filterService.getSpotsByProximityWithState(
      cpLat: searchPosition.latitude,
      cpLng: searchPosition.longitude,
      token: token,
      radious: radious,
    );

    if (filterResponse.token.isEmpty || filterResponse.token != token) {
      sessionDatabase.save(filterResponse.token, searchSessionKey);
    }

    for (var spot in filterResponse.response) {
      if (spot.date.compareTo(DateTime.now()) > 0) {
        spot.spotsColor = colorService.getScheduleColor();
      } else {
        spot.spotsColor = colorService.getColor();
      }
    }

    return filterResponse.response;
  }

  /// 1. Check if tagsOn is empty, if so just return all spots.
  /// 2. To filter
  ///  1. Check if the spot princiapl tag is on tags on, if so then mark flag as true, if not then add the tag as not showed.
  ///  2. Check if the spot secondary tags contains one of the tags on, if so then mark flag as true, if not then add the tag as not showed.
  ///  3. if the mark if flagged if so then add the spot.

  FilteredSpots filterSpotsBaseOnTags(
    Set<String> tagsSelected,
    List<Spot> spots,
  ) {
    if (tagsSelected.isEmpty) {
      return FilteredSpots(
        spots: spots,
        tagsOff: {},
        tagsSelected: tagsSelected,
        onFilterColor: const SpotsColors.empty(),
      );
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
