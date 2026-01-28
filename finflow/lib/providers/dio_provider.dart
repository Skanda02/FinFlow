import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

// Replace with your Go backend URL
final dioProvider = Provider<Dio>((ref) {
  return Dio(
    BaseOptions(
      baseUrl: 'https://your-go-backend-api.com',
      connectTimeout: const Duration(seconds: 5),
    ),
  );
});
