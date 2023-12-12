import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/services/core/providers.dart';

class ProfileFeature extends ConsumerWidget {
  const ProfileFeature({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var userDetails = ref.read(userDetailsProvider);
    var autNotifier = ref.read(userDetailsProvider.notifier);

    return Scaffold(
      body: Center(
          child: TextButton(
        child: const Text("Log out"),
        onPressed: () async {
          await autNotifier.userLogOut();
          Navigator.pop(context);
        },
      )),
    );
  }
}
