import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/contracts/colorService.dart';
import 'package:now_v8/src/core/contracts/filterService.dart';
import 'package:now_v8/src/core/contracts/locationService.dart';
import 'package:now_v8/src/core/models/spot.dart';

class GeneralViewModel {
  final IFilterService filterService;
  final IColorService colorService;
  final ILocationService locationService;

  List<Spot> spots = List.empty();

  GeneralViewModel({
    required this.filterService,
    required this.colorService,
    required this.locationService
  });

  List<Spot> getSpots(){
    LatLng userLocation = locationService.getUserCurrentLocation();

    spots = filterService.getByProximity(cpLat: userLocation.latitude, cpLng: userLocation.longitude);

    spots.forEach((spot) {
      spot.spotsColor = colorService.getColor();
    });

    return spots;
  }
}
