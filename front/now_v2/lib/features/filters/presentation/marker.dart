import 'package:flutter/material.dart';
import 'package:now_v2/core/domain/models/spot.dart';
import 'package:now_v2/core/presentation/spots_granular_view.dart';

class SpotMarker extends StatelessWidget {
  final Spot spotData;

  const SpotMarker({
    Key? key,
    required this.spotData,
  }) : super(key: key);

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
          print(spotData.eventInfo.id);

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
