import 'package:flutter/material.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';

class DoneOrCancelView extends StatelessWidget {
  final String state;
  const DoneOrCancelView({super.key, required this.state});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        children: [Text("DoneOrCancelView " + state)],
      ),
    );
  }
}
