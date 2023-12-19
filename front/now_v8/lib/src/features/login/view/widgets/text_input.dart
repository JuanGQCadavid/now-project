import 'package:flutter/material.dart';

class TextInput extends StatelessWidget {
  final String? hint;
  final void Function(String)? onTextChanged;
  final TextInputType keyboardType;
  final int? minLines;
  final int? maxLines;

  const TextInput(
      {super.key,
      this.hint,
      this.onTextChanged,
      this.keyboardType = TextInputType.text,
      this.maxLines,
      this.minLines});

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      decoration: InputDecoration(
        hintText: hint,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(15.0),
        ),
      ),
      minLines: minLines,
      maxLines: maxLines,
      keyboardType: keyboardType,
      autocorrect: false,
      enableSuggestions: true,
      onTapOutside: (event) {
        WidgetsBinding.instance.focusManager.primaryFocus?.unfocus();
      },
      onChanged: onTextChanged,
    );
  }
}
