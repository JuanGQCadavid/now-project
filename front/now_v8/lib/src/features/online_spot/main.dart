import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/utils/date_utils.dart';

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
        dateTime: DateTime.now(),
        type: NotificationType.systemNotification,
        systemNotifications: SystemNotifications.eventConclude),
    Notification(
      message: "I will cancel the meeting soon, I have to go in 10 mins",
      dateTime: DateTime.now().subtract(const Duration(minutes: 15)),
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

  final EdgeInsets padding;

  NotificationsFeature({
    super.key,
    this.padding = const EdgeInsets.symmetric(
      horizontal: 20,
      vertical: 8,
    ),
  }) {
    notifications.sort((a, b) => b.dateTime.compareTo(a.dateTime));
  }

  @override
  Widget build(BuildContext context) {
    String actualTitle = "";
    return SafeArea(
      child: Container(
        margin: const EdgeInsets.all(15),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              "Notifications",
              style: Theme.of(context).textTheme.displayMedium!,
            ),
            const SizedBox(
              height: 10,
            ),
            Expanded(
              child: Container(
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(15),
                ),
                child: CupertinoScrollbar(
                  thickness: 1.5,
                  // thumbVisibility: true,
                  child: ListView.builder(
                    itemCount: notifications.length,
                    itemBuilder: (context, index) {
                      bool needSeparator = false;
                      if (actualTitle.isNotEmpty) {
                        needSeparator = true;
                      }

                      Widget data = NotificationWidget(
                        notification: notifications[index],
                      );

                      String newTitle = GetDateDiffString(
                        notifications[index].dateTime,
                      );

                      if (actualTitle != newTitle) {
                        actualTitle = newTitle;
                        return Padding(
                          padding: padding,
                          child: Column(
                            crossAxisAlignment: CrossAxisAlignment.start,
                            children: [
                              Visibility(
                                visible: needSeparator,
                                child: const Divider(),
                              ),
                              Container(
                                margin:
                                    const EdgeInsets.symmetric(vertical: 10),
                                child: Text(
                                  actualTitle,
                                  style: Theme.of(context)
                                      .textTheme
                                      .bodyMedium!
                                      .copyWith(fontWeight: FontWeight.bold),
                                ),
                              ),
                              data,
                            ],
                          ),
                        );
                      }
                      return Padding(
                        padding: padding,
                        child: data,
                      );
                    },
                  ),
                ),
              ),
            )
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
    if (notification.type == NotificationType.systemNotification) {
      switch (notification.systemNotifications) {
        case SystemNotifications.eventConclude:
          return _NotificationCard(
            emoji: "👋",
            msg: "The event has ended",
            emojiBackground: Colors.blueAccent.shade100,
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventStopped:
          return _NotificationCard(
            emoji: "🟦",
            emojiBackground: const Color.fromARGB(255, 202, 220, 229),
            msg: "The event has being stopped",
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventResumed:
          return _NotificationCard(
            emoji: "🟩",
            emojiBackground: const Color.fromARGB(255, 220, 238, 200),
            msg: "The event has being resumed",
            dateTime: notification.dateTime,
          );
        case SystemNotifications.eventStarted:
          return _NotificationCard(
            emoji: "🥳",
            msg: "The event has started",
            emojiBackground: const Color.fromARGB(255, 167, 201, 255),
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
