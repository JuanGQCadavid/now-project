import 'package:flutter/material.dart';
import 'package:now_v8/src/features/general_view/views/widgets/locationInfo.dart';

class GeneralViewHeader extends StatelessWidget {
  final greetingMessage = "Welcome back, \n Juan Gonzalo!";

  const GeneralViewHeader({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisSize: MainAxisSize.max,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              IconButton(onPressed: () {}, icon: Icon(Icons.menu)),
              Text(greetingMessage),
              IconButton(onPressed: () {}, icon: Icon(Icons.person))
            ],
          ),
          LocationInfo(),
        ],
      ),
    );
  }
}
