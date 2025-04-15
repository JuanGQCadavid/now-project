import 'dart:developer';
import 'dart:io';

import 'package:web_socket_channel/web_socket_channel.dart';

class NotificationsService {
  late WebSocketChannel channel;

  NotificationsService(String uri) {
    channel = WebSocketChannel.connect(Uri.parse(uri));
  }

  init() async {
    try {
      await channel.ready;
    } on SocketException catch (e) {
      log("we fail to connect to the socket on SocketException: ${e.message}");
      // Handle the exception.
    } on WebSocketChannelException catch (e) {
      log("we fail to connect to the socket on WebSocketChannelException: ${e.message}");
    }
  }

  void istening() {}
}
