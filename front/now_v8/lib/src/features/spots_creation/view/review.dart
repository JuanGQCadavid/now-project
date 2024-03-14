import 'package:flutter/material.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/widgets/tags.dart';

class ReviewView extends StatelessWidget {
  late LongSpot spot = const LongSpot(
    dateInfo: DateInfo(
      dateTime: "",
      id: "",
      startTime: "",
      durationApproximatedInSeconds: 0,
    ),
    eventInfo: EventInfo(
      name: "This is the event name",
      id: "",
      description:
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec quis libero nec eros ultricies volutpat. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Aliquam vel imperdiet mi, ac fermentum leo. Mauris ut ante eget ante dignissim consequat. Donec id est ornare, condimentum tortor porttitor",
      maximunCapacty: 0,
      emoji: ":p",
    ),
    hostInfo: HostInfo(
      name: "Bolis, this is secret",
    ),
    placeInfo: PlaceInfo(
      name: "Ranua de Juana -#AT#- Calle 1 # 1-1",
      lat: 171.0,
      lon: 16.0,
      mapProviderId: "ThisIsTheProviderId",
    ),
    topicInfo: TopicsInfo(
      principalTag: "",
      secondaryTags: [
        "#A",
        "#B",
        "#C",
        "#D",
        "#A",
        "#B",
        "#C",
        "#D",
        "#A",
        "#B",
        "#C",
        "#D"
      ],
    ),
  );
  ReviewView({super.key});
  @override
  Widget build(BuildContext context) {
    List<Widget> tags = [];
    for (var i = 0; i < spot.topicInfo.secondaryTags.length; i++) {
      tags.add(
        Container(
          margin: const EdgeInsets.symmetric(horizontal: 5, vertical: 5),
          child: TagStringView(
            tagValue: spot.topicInfo.secondaryTags[i],
            onDeleteTagPressed: (a) {},
            showDeleteButton: false,
          ),
        ),
      );
    }

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(spot.eventInfo.name),
          SizedBox(
            height: 15,
          ),
          Text(spot.eventInfo.description),
          SizedBox(
            height: 15,
          ),
          tags.length > 0
              ? Wrap(
                  children: tags,
                )
              : const Text("Tags: No tags were selected"),
        ],
      ),
    );
  }
}
