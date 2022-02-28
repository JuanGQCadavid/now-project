import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:now_v2/core/domain/models/spot.dart';
import 'spots_detail_view.dart';

class SpotGranularView extends StatelessWidget {
  final Spot centerSpot;

  const SpotGranularView({
    Key? key,
    required this.centerSpot,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _Body(),
      appBar: AppBar(
        title: Center(
          child: Text(centerSpot.eventInfo.name),
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: () {
        Navigator.push(context,
            MaterialPageRoute(builder: (context) => const SpotDetaillView()));
      }),
    );
  }
}

class _Page extends StatefulWidget {
  final Spot centerSpot;
  const _Page({
    Key? key,
    required this.centerSpot,
  }) : super(key: key);

  @override
  __PageState createState() => __PageState();
}

class __PageState extends State<_Page> {
  @override
  Widget build(BuildContext context) {
    return Container();
  }
}

class _Body extends StatelessWidget {
  const _Body({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        color: Colors.blueGrey,
      ),
    );
  }
}
