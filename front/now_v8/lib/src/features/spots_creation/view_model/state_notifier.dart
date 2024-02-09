import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/features/spots_creation/model/core.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';

class SpotCreator extends StateNotifier<SpotCreatorState> {
  late Map<OnState, Function(bool, LongSpot spot)> mapStates;
  final SpotsCreatorCore core;

  SpotCreator({required this.core})
      : super(
          const SpotCreatorState(
            actualStep: 0,
            totalSteps: 4,
            onState: OnState.onLocation,
            onError: "",
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

  void onNext(LongSpot spot) {
    Function(bool, LongSpot) func = mapStates[super.state.onState]!;
    func(true, spot);
  }

  void onBack() {
    Function(bool, LongSpot) func = mapStates[super.state.onState]!;
    func(false, state.spot);
  }

  void onDescription(bool next, LongSpot spot) {
    print("onDescription");
    print(spot.eventInfo.description);
    print(spot.eventInfo.name);
    if (next) {
      if (spot.eventInfo.description.isEmpty || spot.eventInfo.name.isEmpty) {
        state = state.copyWith(onError: "Title and description are required");
        return;
      }

      var newEvents = state.spot.eventInfo.copyWith(
          name: spot.eventInfo.name, description: spot.eventInfo.description);

      state = state.copyWith(
        onState: OnState.onLocation,
        actualStep: 1,
        spot: spot.copyWith(eventInfo: newEvents),
        onError: "",
      );
    }
  }

  Future onLocation(bool next, LongSpot spot) async {
    if (next) {
      var called = await core.getOptions("WeWork");

      called.fold((l) {
        print("Places");
        print(l);
      }, (r) {
        print("Oh fuck");
      });
      // state = state.copyWith(
      //   onState: OnState.onTags,
      //   actualStep: 2,
      //   onError: "",
      // );
    } else {
      state = state.copyWith(
        onState: OnState.onDescription,
        actualStep: 0,
        onError: "",
      );
    }
  }

  void onTags(bool next, LongSpot spot) {
    if (next) {
      state = state.copyWith(
        onState: OnState.onReview,
        actualStep: 3,
        onError: "",
      );
    } else {
      state = state.copyWith(
        onState: OnState.onLocation,
        actualStep: 1,
        onError: "",
      );
    }
  }

  void onReview(bool next, LongSpot spot) {
    if (next) {
      state = state.copyWith(onState: OnState.onDone);
    } else {
      state = state.copyWith(
        onState: OnState.onTags,
        actualStep: 2,
        onError: "",
      );
    }
  }

  void onCancelle(bool next, LongSpot spot) {}
  void onDone(bool next, LongSpot spot) {}
}
