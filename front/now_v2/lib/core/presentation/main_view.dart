import 'package:flutter/material.dart';
import 'package:now_v2/core/presentation/spots_general_view.dart';

class MainView extends StatelessWidget {
  const MainView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const SpotGeneralView(),
    );
  }
}
