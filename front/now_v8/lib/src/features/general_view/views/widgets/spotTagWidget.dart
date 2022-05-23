import 'package:flutter/material.dart';

class SpotTag extends StatelessWidget {
  final Color color = Colors.green.shade900;
  final String tag = "CoffeLovers";

  SpotTag({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.all(5),
      decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(5),
          border: Border.all(color: color)),
      child: Text(
        "#${tag}",
        style: TextStyle(color: color),
      ),
    );
  }
}


