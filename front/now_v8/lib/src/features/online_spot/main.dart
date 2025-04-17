import 'package:flutter/material.dart';
import 'package:now_v8/src/features/notifications/widgets/full_notificatios.dart';

class OnlineSpotFeature extends StatelessWidget {
  const OnlineSpotFeature({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: Colors.grey.shade100,
        child: Stack(
          children: [
            const Placeholder(),
            NotificationsFeature(),
          ],
        ),
      ),
    );
  }
}
