import 'package:dartz/dartz_unsafe.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/core/widgets/text_input.dart';
import 'package:now_v8/src/features/spots_creation/view_model/providers.dart';

class TagsSelectorView extends ConsumerWidget {
  
  const TagsSelectorView({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var state = ref.watch(tagNotifierProvider);
    var notifier = ref.read(tagNotifierProvider.notifier);

    List<Widget> tags = [];
    for (var i = 0; i < state.length; i++) {
      tags.add(TagStringView(tagValue: state[i]));
    }

    return  Padding(
      padding: const EdgeInsets.symmetric(horizontal: 15.0),
      child: Column(
        children: [
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


class TagStringView extends StatelessWidget {
  final String tagValue;

  const TagStringView({super.key, required this.tagValue});

  @override
  Widget build(BuildContext context) {
    return Text(tagValue);
  }
}
