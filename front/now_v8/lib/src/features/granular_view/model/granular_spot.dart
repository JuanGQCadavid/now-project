import 'package:now_v8/src/core/models/long_spot.dart';

class GranularSpot {
  late SpotWindow window;
  late LongSpot spot;

  GranularSpot({required this.spot, required this.window});
  GranularSpot.empty() {
    window = SpotWindow.emptyWith();
    spot = const LongSpot(
      eventInfo: EventInfo(
        description: "",
        emoji: "",
        id: "",
        maximunCapacty: 0,
        name: "",
      ),
      hostInfo: HostInfo(
        name: "",
      ),
      placeInfo: PlaceInfo(
        lat: 0.0,
        lon: 0.0,
        mapProviderId: "",
        name: "",
      ),
      topicInfo: TopicsInfo(
        principalTopic: "",
        secondaryTopics: [],
      ),
      dateInfo: DateInfo(
        dateTime: "",
        id: "id",
        startTime: "",
        durationApproximatedInSeconds: 0,
      ),
    );
  }
}

class SpotWindow {
  final String nextOne;
  final String actualOne;
  final String previousOne;

  SpotWindow({
    required this.nextOne,
    required this.actualOne,
    required this.previousOne,
  });
  SpotWindow.emptyWith({
    this.actualOne = "",
    this.nextOne = "",
    this.previousOne = "",
  });

  bool isEmpty() {
    return nextOne.isEmpty && actualOne.isEmpty && previousOne.isEmpty;
  }
}
