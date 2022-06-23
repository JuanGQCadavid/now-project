import 'package:flutter/material.dart';


class PlaceLabel extends StatelessWidget {
  final String placeName;
  const PlaceLabel({Key? key, required this.placeName}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(10),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          const Icon(
            Icons.place,
            color: Colors.red,
          ),
          const SizedBox(
            width: 10,
          ),
          Text(
            placeName,
          )
        ],
      ),
    );
  }
}