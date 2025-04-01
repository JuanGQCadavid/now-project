final Map<int, String> months = {
  1: "Jan",
  2: "Feb",
  3: "Mar",
  4: "Apr",
  5: "May",
  6: "Jun",
  7: "Jul",
  8: "Aug",
  9: "Sep",
  10: "Oct",
  11: "Nov",
  12: "Dec",
};

final Map<int, String> longMonths = {
  1: "January",
  2: "February",
  3: "March",
  4: "April",
  5: "May",
  6: "June",
  7: "July",
  8: "August",
  9: "September",
  10: "October",
  11: "November",
  12: "December",
};

final Map<int, String> days = {
  0: "Today",
  1: "Monday",
  2: "Tuesday",
  3: "Wednesday",
  4: "Thursday",
  5: "Friday",
  6: "Saturday",
  7: "Sunday",
};

enum Today { Morning, Afternoon, Evening, Night }

final Map<Today, String> today = {
  Today.Morning: "morning",
  Today.Afternoon: "afternoon",
  Today.Evening: "evening",
  Today.Night: "night",
};

String GetDateDiffString(DateTime dateTime) {
  var now = DateTime.now();

  if (dateTime.year != now.year) {
    return '${dateTime.year}';
  }
  int delta = now.day - dateTime.day;

  if (dateTime.month != now.month) {
    return '${longMonths[dateTime.month]}';
  }

  if (delta == 0) {
    return '${days[0]}';
  }

  if (delta <= 7) {
    return '${days[dateTime.weekday]}';
  }

  return '${days[dateTime.weekday]}, ${dateTime.day}';
}

String GetDateString(DateTime dateTime) {
  var now = DateTime.now();

  if (dateTime.year != now.year) {
    return '${days[dateTime.weekday]}, ${dateTime.day} ${months[dateTime.month]} ${dateTime.year}';
  }

  String minutes = "${dateTime.minute}";

  if (dateTime.minute < 10) {
    minutes = "0${dateTime.minute}";
  }

  String atTime = '${dateTime.hour}:$minutes';

  if (dateTime.month != now.month) {
    return '${days[dateTime.weekday]}, ${dateTime.day} ${months[dateTime.month]} at $atTime';
  }

  int delta = now.day - dateTime.day;
  String timeOfToday = "Today";

  if (dateTime.hour >= 5 && dateTime.hour < 12) {
    timeOfToday = 'This ${today[Today.Morning]}';
  } else if (dateTime.hour >= 12 && dateTime.hour < 17) {
    timeOfToday = 'This ${today[Today.Afternoon]}';
  } else if (dateTime.hour >= 17 && dateTime.hour < 19) {
    timeOfToday = 'This ${today[Today.Evening]}';
  } else {
    timeOfToday = 'This ${today[Today.Night]}';
  }

  if (delta == 0) {
    return '$timeOfToday at $atTime';
  }

  if (delta <= 7) {
    return '${days[dateTime.weekday]} at $atTime';
  }

  return '${days[dateTime.weekday]}, ${dateTime.day} at $atTime';
}
