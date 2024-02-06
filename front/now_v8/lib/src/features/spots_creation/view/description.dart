import 'package:flutter/material.dart';
import 'package:now_v8/src/features/login/view/widgets/text_input.dart';

class SpotGeneralInfo extends StatelessWidget {
  const SpotGeneralInfo({super.key});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        children: [
          TextInput(
            hint: "Title",
            onTextChanged: (value) {},
          ),
          const SizedBox(
            height: 30,
          ),
          SizedBox(
            child: TextInput(
              hint: "Description hi",
              onTextChanged: (value) {},
              keyboardType: TextInputType.multiline,
              minLines: 5,
            ),
          ),
        ],
      ),
    );
  }
}
