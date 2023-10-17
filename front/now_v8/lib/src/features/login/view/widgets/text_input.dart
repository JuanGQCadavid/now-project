import 'package:flutter/material.dart';

class TextInput extends StatelessWidget {
  final String? hint;
  final void Function(String)? onTextChanged;

  const TextInput({super.key, this.hint, this.onTextChanged});

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      decoration: InputDecoration(
        hintText: hint,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(15.0),
        ),
      ),
      keyboardType: TextInputType.text,
      autocorrect: false,
      enableSuggestions: true,
      onTapOutside: (event) {
        WidgetsBinding.instance.focusManager.primaryFocus?.unfocus();
      },
      onChanged: onTextChanged,
    );
  }
}
