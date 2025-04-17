import 'package:now_v8/src/features/notifications/model/notifications.dart';

void sortNotifications(List<Notifications> notifications) {
  notifications.sort((a, b) => b.dateTime.compareTo(a.dateTime));
}
