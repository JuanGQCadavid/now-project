import 'package:flutter/material.dart';
import 'package:now_v8/src/features/notifications/constants.dart';
import 'package:now_v8/src/features/notifications/model/notifications.dart';
import 'package:now_v8/src/features/notifications/widgets/cards.dart';
import 'package:now_v8/src/utils/sorting.dart';

class NotifcationsResume extends StatelessWidget {
  final List<Notifications> notifications;
  NotifcationsResume({super.key, required this.notifications}) {
    sortNotifications(notifications);
  }
  @override
  Widget build(BuildContext context) {
    Widget notificationsResume;

    if (notifications.isEmpty) {
      notificationsResume = Container(
        height: 60,
        margin: const EdgeInsets.only(top: 5),
        decoration: BoxDecoration(
            color: Colors.grey.shade300,
            borderRadius: const BorderRadius.all(Radius.circular(10))),
        child: const Expanded(
            child: Center(
          child: Text("All quiet on the notification front."),
        )),
      );
    } else {
      notificationsResume = NotificationPlusCTA(
        notification: notifications[0],
        numberOfMessages: notifications.length,
      );
    }

    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 14),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            "Notifications",
            style: Theme.of(context).textTheme.titleLarge,
          ),
          notificationsResume,
        ],
      ),
    );
  }
}

class NotificationPlusCTA extends StatelessWidget {
  final Notifications notification;
  final int numberOfMessages;
  const NotificationPlusCTA(
      {super.key, required this.notification, required this.numberOfMessages});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Expanded(
          flex: 3,
          child: Padding(
            padding: const EdgeInsets.only(top: 5),
            child: NotificationWidget(
              notification: notification,
            ),
          ),
        ),
        if (numberOfMessages > 1)
          Expanded(
            flex: 1,
            child: SeeMoreCTA(
              numberOfMessages: numberOfMessages,
            ),
          )
      ],
    );
  }
}

class SeeMoreCTA extends StatelessWidget {
  final int numberOfMessages;
  final String seeMoreTXT = "See all";
  const SeeMoreCTA({
    super.key,
    required this.numberOfMessages,
  });

  @override
  Widget build(BuildContext context) {
    int toDisplay = numberOfMessages - 1;

    if (numberOfMessages > 10) {
      toDisplay = 9;
    }

    return FilledButton.tonal(
      onPressed: () {},
      child: Padding(
        padding: const EdgeInsets.symmetric(vertical: 5),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisSize: MainAxisSize.min,
          children: [
            Center(
              child: Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  const Icon(Icons.add),
                  Text(
                    toDisplay.toString(),
                    style: Theme.of(context).textTheme.titleLarge,
                  )
                ],
              ),
            ),
            Text(
              seeMoreTXT,
              style: Theme.of(context).textTheme.bodySmall,
            ),
          ],
        ),
      ),
    );
  }
}
