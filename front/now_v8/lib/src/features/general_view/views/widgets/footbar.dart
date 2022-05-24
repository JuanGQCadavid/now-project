import 'package:flutter/material.dart';
import 'package:now_v8/src/core/widgets/nowTextCTA.dart';
import 'package:now_v8/src/features/general_view/views/widgets/spotTagWidget.dart';


class BottomBar extends StatelessWidget {
  const BottomBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.all(15),
      //color: Colors.blue,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          NowTextCTA(ctaText: "Filter", onPressed: (){},),
          NowTextCTA(ctaText: "Create", onPressed: (){},),
          NowTextCTA(ctaText: "Zoom in!", onPressed: (){},)
        ],
      ),
    );
  }
}