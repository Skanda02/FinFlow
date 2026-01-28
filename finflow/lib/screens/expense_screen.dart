import 'package:flutter/material.dart';

class ExpenseScreen extends StatelessWidget {
  const ExpenseScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Update Expense')),
      body: const Center(child: Text('Expense Screen')),
    );
  }
}
