import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/features/granular_view/model/ganular_model.dart';
import 'package:now_v8/src/features/granular_view/model/granular_spot.dart';

class DetailedSpotsState extends StateNotifier<List<LongSpot>> {
  final GranularModel granularModel;

  DetailedSpotsState({
    required this.granularModel,
  }) : super([]);

  void refreshSpots() async{
    List<LongSpot> spots = await granularModel.getSpots();
    state = spots;
  }
}

// TODO: We are not thinking in the case that there is only 1, 2 or not events.
class OnSpotState extends StateNotifier<GranularSpot> {
  late List<LongSpot> actualList;
  final GranularModel granularModel;

  OnSpotState({this.actualList = const [], required this.granularModel})
      : super(GranularSpot.empty()) {
        if(actualList.isNotEmpty) {
          state = granularModel.generateNewModel(0, actualList); 
        }
      }

  void previousOne() {
    int spotIndex =
        granularModel.findSpotIndex(state.spot.eventInfo.id, actualList);
    int newIndex = 0;

    if ((spotIndex - 1) >= 0) {
      newIndex = spotIndex - 1;
    } else {
      newIndex = actualList.length - 1;
    }

    state = granularModel.generateNewModel(newIndex, actualList);
  }

  void nextOne() {
    int spotIndex =
        granularModel.findSpotIndex(state.spot.eventInfo.id, actualList);
    int newIndex = 0;

    if ((spotIndex + 1) > (actualList.length - 1)) {
      newIndex = 0;
    } else {
      newIndex = spotIndex + 1;
    }

    state = granularModel.generateNewModel(newIndex, actualList); 
  }


}
