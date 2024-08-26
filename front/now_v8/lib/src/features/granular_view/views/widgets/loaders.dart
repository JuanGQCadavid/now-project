import 'package:flutter/material.dart';

class FindingSpotsLoadingScreen extends StatelessWidget {
  const FindingSpotsLoadingScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          const SizedBox(
            height: 90,
            width: 90,
            child: CircularProgressIndicator(),
          ),
          const SizedBox(
            height: 30,
          ),
          SizedBox(
            width: 300,
            child: Text(
              "We are finding something amazing for you",
              style: Theme.of(context).textTheme.titleLarge,
              textAlign: TextAlign.center,
            ),
          ),
        ],
      ),
    );
  }
}
