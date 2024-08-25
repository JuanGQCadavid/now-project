import 'package:flutter/material.dart';

class HostUpdates extends StatelessWidget {
  final String message = "No updates from host";
  const HostUpdates({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Container(
        margin: const EdgeInsets.only(left: 15, right: 15, bottom: 15),
        color: Colors.grey.shade100,
        child: Center(
          child: Text(message),
        ),
      ),
    );
  }
}
