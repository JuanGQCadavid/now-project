import 'package:flutter/material.dart';

// ignore: must_be_immutable
class CodeInputV2 extends StatelessWidget {
  late List<TextEditingController> controllers;
  late List<FocusNode> focusNodes;
  final int size;

  CodeInputV2({super.key, required this.size}) {
    controllers = [];
    focusNodes = [];
    for (var i = 0; i < size; i++) {
      controllers.add(TextEditingController());
      focusNodes.add(FocusNode());
    }

    focusNodes[0].requestFocus();
  }

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 80,
      child: ListView.builder(
        itemCount: size,
        shrinkWrap: true,
        scrollDirection: Axis.horizontal,
        itemBuilder: (builderContext, index) {
          return _CodeNumnber(
              controller: controllers[index],
              myFocusNode: focusNodes[index],
              nextFocusNode: index == size - 1 ? null : focusNodes[index + 1]);
        },
      ),
    );
  }
}

// ignore: must_be_immutable
class _CodeNumnber extends StatelessWidget {
  TextEditingController controller;
  FocusNode? nextFocusNode;
  FocusNode myFocusNode;

  _CodeNumnber({
    super.key,
    required this.controller,
    required this.myFocusNode,
    this.nextFocusNode,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(left: 10),
      child: SizedBox(
        width: 60,
        child: TextFormField(
          textAlign: TextAlign.center,
          focusNode: myFocusNode,
          maxLength: 1,
          controller: controller,
          keyboardType: TextInputType.number,
          decoration: InputDecoration(
            border: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
            ),
          ),
          onChanged: (value) {
            if (value.isNotEmpty) {
              if (nextFocusNode != null) {
                nextFocusNode!.requestFocus();
              }
            }
          },
          onTapOutside: (event) {
            WidgetsBinding.instance.focusManager.primaryFocus?.unfocus();
          },
        ),
      ),
    );
  }
}
