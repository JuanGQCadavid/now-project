import 'package:flutter/material.dart';

class MyWidget extends StatelessWidget {
  const MyWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}

class TagStringView extends StatelessWidget {
  final String tagValue;
  final bool showDeleteButton;
  final void Function(String) onDeleteTagPressed;

  const TagStringView({
    super.key,
    required this.tagValue,
    required this.onDeleteTagPressed,
    this.showDeleteButton = true,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      child: Row(
        mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          Text(
            tagValue,
            style: Theme.of(context)
                .textTheme
                .bodyLarge!
                .copyWith(color: Colors.white),
          ),
          Visibility(
            visible: showDeleteButton,
            child: IconButton(
              onPressed: () {
                onDeleteTagPressed(tagValue);
              },
              icon: const Icon(
                Icons.close,
                color: Colors.white,
              ),
            ),
          )
        ],
      ),
      padding: const EdgeInsets.symmetric(horizontal: 15, vertical: 5),
      decoration: BoxDecoration(
        color: Theme.of(context).colorScheme.primary,
        borderRadius: BorderRadius.circular(50),
      ),
    );
  }
}
