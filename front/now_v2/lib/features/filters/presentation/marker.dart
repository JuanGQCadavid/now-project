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

class _FullSpotMarker extends StatelessWidget {
  final Spot spotData;
  const _FullSpotMarker({
    Key? key,
    required this.spotData,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(10),
      decoration: const BoxDecoration(
          color: Colors.white38,
          borderRadius: BorderRadius.only(
              bottomLeft: Radius.circular(25),
              bottomRight: Radius.circular(25),
              topLeft: Radius.circular(25),
              topRight: Radius.circular(25))),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(spotData.eventInfo.emoji),
          const SizedBox(
            height: 5,
          ),
          Text(
            spotData.placeInfo.mapProviderId,
            overflow: TextOverflow.ellipsis,
            maxLines: 3,
          ),
        ],
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
    return !showFullInfo
        ? _EmojiBotton(
            spotData: spotData,
          )
        : _FullSpotMarker(
            spotData: spotData,
          );
  }
}
