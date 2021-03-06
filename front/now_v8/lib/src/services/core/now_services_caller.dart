import 'dart:io';

import 'package:dio/dio.dart';
import 'package:dartz/dartz.dart';
import 'package:now_v8/src/services/core/models/backend_errors.dart';
import 'package:now_v8/src/services/core/models/methods.dart';

class NowServicesCaller {
  late Dio _dio;

  NowServicesCaller({required String baseUrl}) {
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
    Map<String, dynamic>? data,
  }) async {
    Response response;

    try {
      if (method == Method.GET) {
        response = await _dio.get(path, queryParameters: queryParameters);
      } else {
        response =
            await _dio.post(path, data: data, queryParameters: queryParameters);
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
        return left(
          const InternalError(),
        );
      }

      return left(
        const UnknownError(),
      );
    } catch (e) {
      return left(
        const UnknownError(),
      );
    }
  }
}
