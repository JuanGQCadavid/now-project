import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/features/notifications/model/notifications.dart';
import 'package:now_v8/src/features/notifications/main.dart';
import 'package:now_v8/src/utils/date_utils.dart';

class NotificationWidget extends StatelessWidget {
  final Notifications notification;
  const NotificationWidget({super.key, required this.notification});

  @override
  Widget build(BuildContext context) {
    if (notification.type == NotificationType.systemNotification) {
      switch (notification.systemNotifications) {
        case SystemNotifications.eventConclude:
          return _NotificationCard(
            emoji: "👋",
            msg: "The event has ended",
            // emojiBackground: Colors.blueAccent.shade100,
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventStopped:
          return _NotificationCard(
            emoji: "✋",
            // emojiBackground: const Color.fromARGB(255, 202, 220, 229),
            msg: "The event has being stopped",
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventResumed:
          return _NotificationCard(
            emoji: "🙌",
            // emojiBackground: const Color.fromARGB(255, 220, 238, 200),
            msg: "The event has being resumed",
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventStarted:
          return _NotificationCard(
            emoji: "🥳",
            msg: "The event has started",
            // emojiBackground: const Color.fromARGB(255, 167, 201, 255),
            dateTime: notification.dateTime,
          );
        case SystemNotifications.empty:
          return _NotificationCard(
            emoji: "🤷‍♀️",
            msg: "Emmmmm... Well this is an empty notification from the system",
            dateTime: notification.dateTime,
          );
      }
    }

    return _NotificationCard(
      emoji: notification.emoji,
      msg: notification.message,
      dateTime: notification.dateTime,
    );
  }
}

class _NotificationCard extends StatelessWidget {
  final String emoji;
  final String msg;
  final DateTime dateTime;
  late Color? emojiBackground;

  var defaultColor = Colors.grey.shade100;

  _NotificationCard({
    super.key,
    required this.emoji,
    required this.msg,
    required this.dateTime,
    this.emojiBackground,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      mainAxisSize: MainAxisSize.max,
      children: [
        Container(
          decoration: BoxDecoration(
            color: emojiBackground ?? defaultColor,
            shape: BoxShape.circle,
          ),
          padding: const EdgeInsets.all(15),
          child: Text(
            emoji,
            style: const TextStyle(fontSize: 25),
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
                msg,
                style: Theme.of(context).textTheme.bodyMedium!.copyWith(
                      fontWeight: FontWeight.bold,
                    ),
                overflow: TextOverflow.visible,
              ),
              Text(GetDateString(dateTime))
            ],
          ),
        )
      ],
    );
  }
}
