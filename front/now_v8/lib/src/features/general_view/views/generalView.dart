import 'package:flutter/material.dart';
import 'package:now_v8/src/features/general_view/views/widgets/footbar.dart';
import 'package:now_v8/src/features/general_view/views/widgets/header.dart';
import 'package:now_v8/src/features/general_view/views/widgets/nowMap.dart';

class GeneralViewFeature extends StatelessWidget {
  const GeneralViewFeature({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: GeneralViewBody(),
        bottomNavigationBar: BottomBar(),
      ),
    );
  }
}

class GeneralViewBody extends StatelessWidget {
  const GeneralViewBody({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          child: GeneralViewHeader(),
          margin: const EdgeInsets.all(10),
        ),
        Expanded(
          child: ClipRRect(
            borderRadius: BorderRadius.only(topLeft: Radius.circular(25), topRight: Radius.circular(25),bottomLeft: Radius.circular(25) ),
            child: MapSample(),
          ),
        )
      ],
    );
  }
}
