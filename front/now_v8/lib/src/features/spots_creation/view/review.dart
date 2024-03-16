import 'dart:async';

import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import 'package:now_v8/src/core/models/long_spot.dart';
import 'package:now_v8/src/core/models/spot.dart';
import 'package:now_v8/src/core/widgets/nowMap.dart';
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
      lat: 6.2379578,
      lon: -75.5626034,
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
    var adresss = spot.placeInfo.name.split("-#AT#-");

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
          const Text("Resume"),
          const SizedBox(
            height: 15,
          ),
          const Divider(),
          Text(
            spot.eventInfo.name,
            style: Theme.of(context).textTheme.titleLarge,
          ),
          const SizedBox(
            height: 15,
          ),
          Text(spot.eventInfo.description),
          const SizedBox(
            height: 15,
          ),
          tags.isNotEmpty
              ? Wrap(
                  children: tags,
                )
              : const Text("Tags: No tags were selected"),
          const SizedBox(
            height: 15,
          ),
          Text("${adresss[0]}, ${adresss[1]}"),
          const SizedBox(
            height: 15,
          ),
          SizedBox(
            height: 200,
            width: double.infinity,
            child: NowMapV2(
              centerMapOnSpots: true,
              includeUserLocation: false,
              camaraPosition: LatLng(
                spot.placeInfo.lat,
                spot.placeInfo.lon,
              ),
              mapController: Completer(),
              myLocationButtonEnabled: false,
              blockMap: true,
              spots: [
                Spot.withOutSpotColors(
                  principalTag: "",
                  secondaryTags: [],
                  latLng: LatLng(
                    spot.placeInfo.lat,
                    spot.placeInfo.lon,
                  ),
                  spotId: "",
                  date: DateTime.now(),
                )
              ],
            ),
          ),
        ],
      ),
    );
  }
}
