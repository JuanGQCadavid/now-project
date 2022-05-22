import 'package:flutter/material.dart';
import 'package:now_v8/src/features/general_view/views/widgets/header.dart';

class GeneralViewFeature extends StatelessWidget {
  const GeneralViewFeature({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: GeneralViewBody(),
        bottomNavigationBar: Container(
          color: Colors.blue,
          child: Text("Hi!"),
        ),
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
          child: Container(
            color: Colors.blue,
            child: Text("Here will be the map"),
          ),
        )
      ],
    );
  }
}
