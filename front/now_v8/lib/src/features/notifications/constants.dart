import 'package:now_v8/src/features/notifications/model/notifications.dart';

List<Notifications> globalExmapleNotifications = <Notifications>[
  Notifications(
      dateTime: DateTime.now(),
      type: NotificationType.systemNotification,
      systemNotifications: SystemNotifications.eventConclude),
  Notifications(
    message: "I will cancel the meeting soon, I have to go in 10 mins",
    dateTime: DateTime.now().subtract(const Duration(minutes: 15)),
    emoji: "😅",
    type: NotificationType.userNotification,
  ),
  Notifications(
    message: "The chair does not have good ligth, we move to table C.",
    dateTime: DateTime.now().subtract(const Duration(
      days: 1,
      minutes: 10,
    )),
    emoji: "💡",
    type: NotificationType.userNotification,
  ),
  Notifications(
    dateTime: DateTime.now().subtract(const Duration(
      days: 2,
      minutes: 30,
    )),
    type: NotificationType.systemNotification,
    systemNotifications: SystemNotifications.eventResumed,
  ),
  Notifications(
    dateTime: DateTime.now().subtract(const Duration(
      days: 10,
      minutes: 60,
    )),
    emoji: "",
    type: NotificationType.systemNotification,
    systemNotifications: SystemNotifications.eventStopped,
  ),
  Notifications(
    message: "We are going to be close to the Dance floor",
    dateTime: DateTime.now().subtract(const Duration(
      days: 40,
      hours: 1,
      minutes: 40,
    )),
    emoji: "💃🏻",
    type: NotificationType.userNotification,
  ),
  Notifications(
    dateTime: DateTime.now().subtract(const Duration(
      days: 368,
      hours: 2,
    )),
    type: NotificationType.systemNotification,
    systemNotifications: SystemNotifications.eventStarted,
  ),
];
