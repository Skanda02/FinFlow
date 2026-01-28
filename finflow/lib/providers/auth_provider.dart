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

    // 2. Simulate a network request
    await Future.delayed(const Duration(seconds: 1));

    // 3.
    // START: Backend login logic (commented out for development)
    // state = await AsyncValue.guard(() async {
    //   final dio = ref.read(dioProvider);

    //   final response = await dio.post(
    //     '/login',
    //     data: {'email': email, 'password': password},
    //   );

    //   if (response.statusCode != 200) {
    //     throw Exception("Login Failed: ${response.statusMessage}");
    //   }
    //   // If successful, you would usually store a JWT token here
    //   print("Login Success: ${response.data}");
    // });
    // END: Backend login logic

    // START: Mock login logic for development
    // In a real app, you would get a user object or token here
    state = const AsyncData(null);
    // END: Mock login logic
  }
}
