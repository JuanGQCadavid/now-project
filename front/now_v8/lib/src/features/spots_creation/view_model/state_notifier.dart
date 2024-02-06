import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';

class SpotCreator extends StateNotifier<SpotCreatorState> {
  late Map<OnState, Function(bool)> mapStates;

  SpotCreator()
      : super(
          const SpotCreatorState(
            actualStep: 0,
            totalSteps: 4,
            onState: OnState.onDescription,
            spot: LongSpot(
              dateInfo: DateInfo(
                dateTime: "",
                id: "",
                startTime: "",
                durationApproximatedInSeconds: 0,
              ),
              eventInfo: EventInfo(
                name: "",
                id: "",
                description: "",
                maximunCapacty: 0,
                emoji: ":p",
              ),
              hostInfo: HostInfo(
                name: "",
              ),
              placeInfo: PlaceInfo(
                name: "",
                lat: 0.0,
                lon: 0.0,
                mapProviderId: "",
              ),
              topicInfo: TopicsInfo(
                principalTag: "",
                secondaryTags: [],
              ),
            ),
          ),
        ) {
    mapStates = {
      OnState.onDone: onDone,
      OnState.onDescription: onDescription,
      OnState.onLocation: onLocation,
      OnState.onTags: onTags,
      OnState.onReview: onReview,
      OnState.onCancelle: onCancelle,
    };
  }

  // void onNext(LongSpot spot) {}

  // void onBack(LongSpot spot) {}

  void onNext() {
    Function(bool) func = mapStates[super.state.onState]!;
    func(true);
  }

  void onBack() {
    Function(bool) func = mapStates[super.state.onState]!;
    func(false);
  }

  void onDescription(bool next) {
    if (next) {
      state = state.copyWith(onState: OnState.onLocation, actualStep: 1);
    }
  }

  void onLocation(bool next) {
    if (next) {
      state = state.copyWith(onState: OnState.onTags, actualStep: 2);
    } else {
      state = state.copyWith(onState: OnState.onDescription, actualStep: 0);
    }
  }

  void onTags(bool next) {
    if (next) {
      state = state.copyWith(onState: OnState.onReview, actualStep: 3);
    } else {
      state = state.copyWith(onState: OnState.onLocation, actualStep: 1);
    }
  }

  void onReview(bool next) {
    if (next) {
      state = state.copyWith(onState: OnState.onDone);
    } else {
      state = state.copyWith(onState: OnState.onTags, actualStep: 2);
    }
  }

  void onCancelle(bool next) {}
  void onDone(bool next) {}
}
