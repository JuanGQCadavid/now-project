import 'package:flutter/material.dart';


class PlaceLabel extends StatelessWidget {
  final String placeName;
  final Color appColor;
  const PlaceLabel({Key? key, required this.placeName, required this.appColor}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.all(10),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(
            Icons.place,
            color: appColor,
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