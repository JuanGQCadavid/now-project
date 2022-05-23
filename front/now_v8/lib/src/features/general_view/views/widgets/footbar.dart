import 'package:flutter/material.dart';
import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';


class BottomBar extends StatelessWidget {
  const BottomBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      //color: Colors.blue,
      child: Row(
        children: [
          SpotTag()
        ],
      ),
    );
  }
}