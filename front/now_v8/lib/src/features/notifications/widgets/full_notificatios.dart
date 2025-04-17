import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:now_v8/src/features/notifications/model/notifications.dart';
import 'package:now_v8/src/features/notifications/constants.dart';
import 'package:now_v8/src/features/notifications/widgets/cards.dart';
import 'package:now_v8/src/utils/date_utils.dart';
import 'package:now_v8/src/utils/sorting.dart';

class FullNotificationsView extends StatelessWidget {
  final List<Notifications> notifications;

  FullNotificationsView({
    super.key,
    required this.notifications,
  }) {
    sortNotifications(notifications);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white.withOpacity(0.8),
      body: NotificationsFeature(
        notifications: notifications,
      ),
    );
  }
}

class NotificationsFeature extends StatelessWidget {
  final EdgeInsets padding;
  final List<Notifications> notifications;

  NotificationsFeature({
    super.key,
    this.padding = const EdgeInsets.symmetric(
      horizontal: 20,
      vertical: 8,
    ),
    required this.notifications,
  }) {
    sortNotifications(notifications);
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
            Row(
              children: [
                IconButton(
                  onPressed: () => Navigator.of(context).pop(),
                  icon: const Icon(Icons.navigate_before),
                ),
                const SizedBox(
                  width: 10,
                ),
                Text(
                  "Notifications",
                  style: Theme.of(context).textTheme.displaySmall!,
                ),
              ],
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
