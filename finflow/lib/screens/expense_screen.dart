import 'package:flutter/material.dart';
import 'package:math_expressions/math_expressions.dart';

class ExpenseScreen extends StatefulWidget {
  const ExpenseScreen({super.key});

  @override
  State<ExpenseScreen> createState() => _ExpenseScreenState();
}

class _ExpenseScreenState extends State<ExpenseScreen> {
  final _formKey = GlobalKey<FormState>();
  final _amountController = TextEditingController();
  final _descriptionController = TextEditingController();

  @override
  void dispose() {
    _amountController.dispose();
    _descriptionController.dispose();
    super.dispose();
  }

  void _submitExpense() {
    if (_formKey.currentState!.validate()) {
      try {
        Parser p = Parser();
        Expression exp = p.parse(_amountController.text);
        ContextModel cm = ContextModel();
        final amount = exp.evaluate(EvaluationType.REAL, cm);
        final description = _descriptionController.text;
        final timestamp = DateTime.now();

        // TODO: Send data to the database
        print('Expense Amount: $amount');
        print('Description: $description');
        print('Timestamp: $timestamp');

        // Clear the form
        _formKey.currentState!.reset();
        _amountController.clear();
        _descriptionController.clear();

        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Expense added successfully!')),
        );
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Invalid amount expression')),
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
              'Add Expense',
              style: Theme.of(context).textTheme.headlineSmall,
            ),
            const SizedBox(height: 24),
            TextFormField(
              controller: _amountController,
              decoration: const InputDecoration(
                labelText: 'Amount',
                prefixIcon: Icon(Icons.attach_money),
                border: OutlineInputBorder(),
              ),
              keyboardType: TextInputType.text,
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please enter an amount';
                }
                try {
                  Parser p = Parser();
                  p.parse(value);
                  return null;
                } catch (e) {
                  return 'Please enter a valid expression';
                }
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
                onPressed: _submitExpense,
                child: const Text('Add Expense'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
