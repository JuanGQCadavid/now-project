import 'package:flutter/material.dart';

class DoneOrCancelView extends StatelessWidget {
  final String state;
  final String id;
  const DoneOrCancelView({super.key, required this.state, required this.id});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        children: [Text("DoneOrCancelView " + state + " id: " + id)],
      ),
    );
  }
}
