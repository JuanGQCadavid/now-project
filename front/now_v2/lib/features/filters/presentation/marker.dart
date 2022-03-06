import 'package:flutter/material.dart';
import 'package:now_v2/core/domain/models/spot.dart';
import 'package:now_v2/core/presentation/spots_granular_view.dart';

class _EmojiBotton extends StatelessWidget {
  final Spot spotData;
  const _EmojiBotton({Key? key, required this.spotData}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.only(
              bottomLeft: Radius.circular(50),
              bottomRight: Radius.circular(50),
              topLeft: Radius.circular(50),
              topRight: Radius.circular(50))),
      child: TextButton(
        child: Text(spotData.eventInfo.emoji),
        onPressed: () {
          Navigator.push(
              context,
              MaterialPageRoute(
                  builder: (context) =>
                      SpotGranularView(centerSpot: spotData)));
        },
      ),
    );
  }
}

class SpotMarker extends StatelessWidget {
  final Spot spotData;
  late bool showFullInfo;

  SpotMarker({
    Key? key,
    required this.spotData,
    this.showFullInfo = false,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final emojiIcon = _EmojiBotton(
      spotData: spotData,
    );

    return !showFullInfo
        ? emojiIcon
        : Column(
            children: [
              emojiIcon,
              Visibility(
                child: Center(
                  child: Text(
                    spotData.placeInfo.mapProviderId,
                    overflow: TextOverflow.ellipsis,
                    maxLines: 3,
                  ),
                ),
                visible: true,
              )
            ],
          );
  }
  // @override
  // Widget build(BuildContext context) {
  //   return Container(
  //     decoration: const BoxDecoration(
  //         color: Colors.white,
  //         borderRadius: BorderRadius.only(
  //             bottomLeft: Radius.circular(50),
  //             bottomRight: Radius.circular(50),
  //             topLeft: Radius.circular(50),
  //             topRight: Radius.circular(50))),
  //     child: TextButton(
  //       child: Text(spotData.eventInfo.emoji),
  //       onPressed: () {
  //         print(spotData.eventInfo.id);

  //         Navigator.push(
  //             context,
  //             MaterialPageRoute(
  //                 builder: (context) =>
  //                     SpotGranularView(centerSpot: spotData)));
  //       },
  //     ),
  //   );
  // }
}
