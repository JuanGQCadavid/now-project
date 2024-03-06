import 'package:flutter/material.dart';
import 'package:now_v8/src/core/widgets/text_input.dart';

class SpotGeneralInfo extends StatelessWidget {
  final void Function(String) onTitleChanged;
  final TextEditingController? titleController;

  final void Function(String) onDescriptionChange;
  final TextEditingController? descriptionController;

  final String? errMessage;

  const SpotGeneralInfo({
    super.key,
    required this.onDescriptionChange,
    required this.onTitleChanged,
    this.descriptionController,
    this.titleController,
    this.errMessage,
  });
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        children: [
          TextInput(
            hint: "Title",
            onTextChanged: onTitleChanged,
            controller: titleController,
          ),
          const SizedBox(
            height: 30,
          ),
          SizedBox(
            child: TextInput(
              hint: "Description hi",
              onTextChanged: onDescriptionChange,
              controller: descriptionController,
              keyboardType: TextInputType.multiline,
              minLines: 5,
            ),
          ),
          Visibility(
            visible: errMessage != null,
            child: Container(
              margin: const EdgeInsets.all(30),
              child: Text(
                errMessage ?? "",
                style: const TextStyle(color: Colors.red),
              ),
            ),
          )
        ],
      ),
    );
  }
}
