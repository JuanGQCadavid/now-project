enum NotificationType { userNotification, systemNotification }

enum SystemNotifications {
  eventStopped,
  eventResumed,
  eventConclude,
  eventStarted,
  empty
}

class Notifications {
  final String emoji;
  final DateTime dateTime;
  final String message;
  final NotificationType type;
  final SystemNotifications systemNotifications;

  const Notifications({
    required this.dateTime,
    required this.type,
    this.message = "",
    this.emoji = "",
    this.systemNotifications = SystemNotifications.empty,
  });
}
