import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/features/spots_creation/model/spot_creator_state.dart';

class SpotCreator extends StateNotifier<SpotCreatorState> {
  // SpotCreator(super.state);

  SpotCreator()
      : super(
          const SpotCreatorState(
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
                )),
          ),
        );

  void onNext(LongSpot spot) {}
}
