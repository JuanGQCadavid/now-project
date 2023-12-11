import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:now_v8/src/services/core/providers.dart';

class ProfileFeature extends ConsumerWidget {
  const ProfileFeature({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    var authService = ref.read(authProvider);

    return Scaffold(
      body: Center(
          child: TextButton(
        child: const Text("Log out"),
        onPressed: () async {
          await authService.removeUserDetails();
          Navigator.pop(context);
        },
      )),
    );
  }
}
