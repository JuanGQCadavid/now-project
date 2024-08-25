import 'package:dartz/dartz_unsafe.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/widgets/tags.dart';
import 'package:now_v8/src/core/widgets/text_input.dart';
import 'package:now_v8/src/features/spots_creation/view_model/providers.dart';

class TagsSelectorView extends ConsumerWidget {
  final void Function(List<String>) tagsSelected;

  const TagsSelectorView({super.key, required this.tagsSelected});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var state = ref.watch(tagNotifierProvider);
    var notifier = ref.read(tagNotifierProvider.notifier);
    notifier.setCallback(tagsSelected);

    List<Widget> tags = [];
    for (var i = 0; i < state.length; i++) {
      tags.add(
        TagStringView(
          tagValue: state[i],
          onDeleteTagPressed: notifier.removeTag,
        ),
      );
    }

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        mainAxisSize: MainAxisSize.max,
        children: [
          Text(
            "Do you want to add some tags?",
            style: Theme.of(context).textTheme.titleLarge,
          ),
          SizedBox(height: 25),
          Wrap(
            children: tags,
            spacing: 10,
            runSpacing: 10,
          ),
          TextField(
            onSubmitted: notifier.addTag,
            controller: notifier.controller,
            autofocus: true,
            focusNode: notifier.focus,
          )
        ],
      ),
    );
  }
}
