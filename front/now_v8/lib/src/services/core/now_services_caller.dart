import 'dart:io';

import 'package:dio/dio.dart';
import 'package:dartz/dartz.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';

class NowServicesCaller {
  late Dio _dio;
  late String baseUrl;

  NowServicesCaller({required this.baseUrl}) {
    _dio = Dio(
      BaseOptions(
        baseUrl: baseUrl,
      ),
    );
  }

  Future<Either<BackendErrors, dynamic>> request(
    Method method,
    String path, {
    Map<String, dynamic>? queryParameters,
    Map<String, dynamic>? headers,
    Map<String, dynamic>? data,
  }) async {
    Response response;

    print("~~~~~~~~~~~~~~~~~~~~~~~ Backend caller ~~~~~~~~~~~~~~~~~~~~~~~");
    print("URL: ${baseUrl} ${path}");
    print("");

    try {
      Options options = Options();
      if (headers != null) {
        options = Options(headers: headers);
      }

      if (method == Method.GET) {
        print("Method: GET");
        response = await _dio.get(path,
            queryParameters: queryParameters, options: options);
      } else {
        print("Method: POST");
        response = await _dio.post(path,
            data: data, queryParameters: queryParameters, options: options);
      }

      // Checkers on 500
      if (response.statusCode == 500) {
        return left(
          const InternalError(),
        );
      } else if (response.statusCode == 503) {
        return left(
          const ServiceUnavailable(),
        );
      }

      print(response.data);
      return right(response.data);
    } on SocketException catch (e) {
      return left(
        const NoInternetConnection(),
      );
    } on FormatException catch (e) {
      return left(
        const BadResponseFormat(),
      );
    } on DioError catch (e) {
      if (e.type == DioErrorType.response) {
        Println(e.message);
        var data = e.response?.data ?? "";
        if (data is String || data.runtimeType == String) {
          Println(data);
          return left(
            ClientError(ErrorMessage("NONE", e.message, "NONE")),
          );
        } else {
          return left(
            ClientError(ErrorMessage.fromJson(e.response?.data ?? "")),
          );
        }
      }
      print(e.toString());
      return left(
        const UnknownError(),
      );
    } catch (e) {
      print(e.toString());
      return left(
        const UnknownError(),
      );
    }
  }
}
