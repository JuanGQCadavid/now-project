import 'package:flutter/material.dart';

class LocationInfo extends StatelessWidget {
  final double size = 11;
  final Color color = Colors.green;
  final String locationAddress = "Medellin, Antioquia.";

  const LocationInfo({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Row(mainAxisSize: MainAxisSize.min, children: [
        Container(
          height: size,
          width: size,
          decoration: BoxDecoration(
              color: color, borderRadius: BorderRadius.circular(100)),
        ),
        const SizedBox(
          width: 10,
        ),
        Text(locationAddress)
      ]),
    );
  }
}
