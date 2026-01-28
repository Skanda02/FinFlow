import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

// --- REST CLIENT ---
final dioProvider = Provider<Dio>((ref) {
  return Dio(
    BaseOptions(
      baseUrl:
          'https://api.your-go-backend.com', // Replace with your team's Go server URL
      connectTimeout: const Duration(seconds: 5),
    ),
  );
});

// --- STATE MANAGEMENT ---
final authNotifierProvider = AsyncNotifierProvider<AuthNotifier, void>(() {
  return AuthNotifier();
});

class AuthNotifier extends AsyncNotifier<void> {
  @override
  Future<void> build() async {} // Initial state: doing nothing.

  Future<void> login(String email, String password) async {
    // 1. Set state to loading (triggers spinner in UI)
    state = const AsyncLoading();

    // 2. Perform the POST request to your Go Backend
    state = await AsyncValue.guard(() async {
      final dio = ref.read(dioProvider);

      final response = await dio.post(
        '/login',
        data: {'email': email, 'password': password},
      );

      if (response.statusCode != 200) {
        throw Exception("Login Failed: ${response.statusMessage}");
      }

      // If successful, you would usually store a JWT token here
      print("Login Success: ${response.data}");
    });
  }
}
