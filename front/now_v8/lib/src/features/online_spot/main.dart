import 'package:flutter/material.dart';
import 'package:now_v8/src/utils/date_utils.dart';

class OnlineSpotFeature extends StatelessWidget {
  const OnlineSpotFeature({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        color: Colors.amber,
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

enum NotificationType { userNotification, systemNotification }

enum SystemNotifications {
  eventStopped,
  eventResumed,
  eventConclude,
  eventStarted,
  empty
}

class Notification {
  final String emoji;
  final DateTime dateTime;
  final String message;
  final NotificationType type;
  final SystemNotifications systemNotifications;

  const Notification({
    required this.dateTime,
    required this.type,
    this.message = "",
    this.emoji = "",
    this.systemNotifications = SystemNotifications.empty,
  });
}

class NotificationsFeature extends StatelessWidget {
  List<Notification> notifications = <Notification>[
    Notification(
      message: "I will cancel the meeting soon, I have to go in 10 mins",
      dateTime: DateTime.now(),
      emoji: "😅",
      type: NotificationType.userNotification,
    ),
    Notification(
      message: "The chair does not have good ligth, we move to table C.",
      dateTime: DateTime.now().subtract(const Duration(
        days: 1,
        minutes: 10,
      )),
      emoji: "💡",
      type: NotificationType.userNotification,
    ),
    Notification(
      dateTime: DateTime.now().subtract(const Duration(
        days: 2,
        minutes: 30,
      )),
      type: NotificationType.systemNotification,
      systemNotifications: SystemNotifications.eventResumed,
    ),
    Notification(
      dateTime: DateTime.now().subtract(const Duration(
        days: 10,
        minutes: 60,
      )),
      emoji: "",
      type: NotificationType.systemNotification,
      systemNotifications: SystemNotifications.eventStopped,
    ),
    Notification(
      message: "We are going to be close to the Dance floor",
      dateTime: DateTime.now().subtract(const Duration(
        days: 40,
        hours: 1,
        minutes: 40,
      )),
      emoji: "💃🏻",
      type: NotificationType.userNotification,
    ),
    Notification(
      dateTime: DateTime.now().subtract(const Duration(
        days: 368,
        hours: 2,
      )),
      type: NotificationType.systemNotification,
      systemNotifications: SystemNotifications.eventStarted,
    ),
  ];
  NotificationsFeature({super.key});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Container(
        margin: const EdgeInsets.all(15),
        color: Colors.white,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            Text(
              "Notifications",
              style: Theme.of(context).textTheme.titleLarge,
            ),
            SizedBox(
              height: 10,
            ),
            NotificationWidget(
              notification: notifications[0],
            ),
            NotificationWidget(
              notification: notifications[1],
            ),
            NotificationWidget(
              notification: notifications[2],
            ),
            NotificationWidget(
              notification: notifications[3],
            ),
            NotificationWidget(
              notification: notifications[4],
            ),
            NotificationWidget(
              notification: notifications[5],
            ),
          ],
        ),
      ),
    );
  }
}

class NotificationWidget extends StatelessWidget {
  final Notification notification;
  const NotificationWidget({super.key, required this.notification});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      mainAxisSize: MainAxisSize.max,
      children: [
        Container(
          decoration: BoxDecoration(
            color: Colors.grey.shade100,
            shape: BoxShape.circle,
          ),
          padding: const EdgeInsets.all(15),
          child: Text(
            notification.emoji,
            style: const TextStyle(fontSize: 40),
          ),
        ),
        const SizedBox(
          width: 10,
        ),
        Flexible(
          child: Column(
            mainAxisSize: MainAxisSize.max,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                notification.message,
                style: Theme.of(context).textTheme.bodyMedium!.copyWith(
                      fontWeight: FontWeight.bold,
                    ),
                overflow: TextOverflow.visible,
              ),
              Text(GetDateString(notification.dateTime))
            ],
          ),
        )
      ],
    );
  }
}
