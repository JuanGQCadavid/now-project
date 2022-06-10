import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/models/spotColors.dart';

class FilteredSpots {
  late Set<String> tagsSelected;
  late Set<String> tagsOff;
  late List<Spot> spots;
  late SpotsColors onFilterColor;

  FilteredSpots({
    required this.spots,
    required this.tagsOff,
    required this.tagsSelected,
    required this.onFilterColor,
  });

  FilteredSpots.empty(){
    tagsOff = {};
    spots = [];
    tagsSelected = {};
    onFilterColor = const SpotsColors.empty();
  }

}
