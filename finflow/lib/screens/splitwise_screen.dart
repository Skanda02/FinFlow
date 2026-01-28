import 'package:flutter/material.dart';

class SplitwiseScreen extends StatefulWidget {
  const SplitwiseScreen({super.key});

  @override
  State<SplitwiseScreen> createState() => _SplitwiseScreenState();
}

class _SplitwiseScreenState extends State<SplitwiseScreen> {
  final _formKey = GlobalKey<FormState>();
  final _emailController = TextEditingController();
  final _amountController = TextEditingController();
  final _descriptionController = TextEditingController();

  @override
  void dispose() {
    _emailController.dispose();
    _amountController.dispose();
    _descriptionController.dispose();
    super.dispose();
  }

  void _submitSplit() {
    if (_formKey.currentState!.validate()) {
      final email = _emailController.text;
      final amount = double.tryParse(_amountController.text);
      final description = _descriptionController.text;
      final timestamp = DateTime.now();

      if (amount != null) {
        // TODO: Send data to the database
        print('Email: $email');
        print('Amount: $amount');
        print('Description: $description');
        print('Timestamp: $timestamp');

        // Clear the form
        _formKey.currentState!.reset();
        _emailController.clear();
        _amountController.clear();
        _descriptionController.clear();

        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Split request sent successfully!')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Form(
        key: _formKey,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Split Expense',
              style: Theme.of(context).textTheme.headlineSmall,
            ),
            const SizedBox(height: 24),
            TextFormField(
              controller: _emailController,
              decoration: const InputDecoration(
                labelText: 'User Email',
                prefixIcon: Icon(Icons.email),
                border: OutlineInputBorder(),
              ),
              keyboardType: TextInputType.emailAddress,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter an email';
                }
                if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(value)) {
                  return 'Please enter a valid email';
                }
                return null;
              },
            ),
            const SizedBox(height: 16),
            TextFormField(
              controller: _amountController,
              decoration: const InputDecoration(
                labelText: 'Amount',
                prefixIcon: Icon(Icons.attach_money),
                border: OutlineInputBorder(),
              ),
              keyboardType: TextInputType.number,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter an amount';
                }
                if (double.tryParse(value) == null) {
                  return 'Please enter a valid number';
                }
                return null;
              },
            ),
            const SizedBox(height: 16),
            TextFormField(
              controller: _descriptionController,
              decoration: const InputDecoration(
                labelText: 'Description',
                prefixIcon: Icon(Icons.description),
                border: OutlineInputBorder(),
              ),
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter a description';
                }
                return null;
              },
            ),
            const SizedBox(height: 24),
            SizedBox(
              width: double.infinity,
              child: ElevatedButton(
                onPressed: _submitSplit,
                child: const Text('Send Split Request'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
