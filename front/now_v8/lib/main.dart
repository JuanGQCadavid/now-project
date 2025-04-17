import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:now_v8/src/features/general_view/views/main.dart';
import 'package:now_v8/src/features/notifications/constants.dart';
import 'package:now_v8/src/features/notifications/model/notifications.dart';
import 'package:now_v8/src/features/notifications/widgets/resume_notifications.dart';
import 'package:now_v8/src/features/online_spot/main.dart';

void main() async {
  await Hive.initFlutter();
  runApp(
    ProviderScope(
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      // theme: ThemeData(
      //   colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      //   useMaterial3: true,
      // ),
      theme: ThemeData(
        scaffoldBackgroundColor: Colors.white,
        useMaterial3: true,
      ),
      home:
          MyWidget(), //GeneralViewFeature() //GeneralViewFeature MapsClusterDemoTwo SpotsCreationFeature HomeTest OnlineSpotFeature
    );
  }
}

class MyWidget extends StatelessWidget {
  const MyWidget({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: Colors.grey.shade100,
        child: Column(
          children: [
            const Placeholder(),
            NotifcationsResume(notifications: <Notifications>[
              Notifications(
                dateTime: DateTime.now().subtract(
                  const Duration(
                    days: 1,
                    minutes: 15,
                  ),
                ),
                type: NotificationType.systemNotification,
                systemNotifications: SystemNotifications.eventConclude,
              ),
              Notifications(
                message:
                    "The chair does not have good ligth, we move to table C.",
                dateTime: DateTime.now().subtract(
                  const Duration(
                    days: 1,
                    minutes: 10,
                  ),
                ),
                emoji: "💡",
                type: NotificationType.userNotification,
              ),
            ] //List.from(globalExmapleNotifications),
                ),
            // Placeholder(),
          ],
        ),
      ),
    );
  }
}
