import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';

class SummaryScreen extends StatefulWidget {
  const SummaryScreen({super.key});

  @override
  State<SummaryScreen> createState() => _SummaryScreenState();
}

class _SummaryScreenState extends State<SummaryScreen> {
  int _selectedInterval = 0; // 0: Weekly, 1: Monthly, 2: Yearly

  // Placeholder data
  final Map<String, double> weeklyData = {
    'Mon': 120.0,
    'Tue': 200.0,
    'Wed': 150.0,
    'Thu': 300.0,
    'Fri': 250.0,
    'Sat': 180.0,
    'Sun': 220.0,
  };

  final Map<String, String> weeklyDescriptions = {
    'Mon': 'Groceries, Transport',
    'Tue': 'Restaurant, Shopping',
    'Wed': 'Utilities, Fuel',
    'Thu': 'Entertainment, Dining',
    'Fri': 'Shopping, Coffee',
    'Sat': 'Movie, Groceries',
    'Sun': 'Brunch, Gifts',
  };

  final Map<String, double> monthlyData = {
    'Week 1': 800.0,
    'Week 2': 1200.0,
    'Week 3': 1000.0,
    'Week 4': 1500.0,
    'Week 5': 1300.0,
  };

  final Map<String, String> monthlyDescriptions = {
    'Week 1': 'Daily expenses, utilities',
    'Week 2': 'Shopping, entertainment',
    'Week 3': 'Groceries, transport',
    'Week 4': 'Major shopping, dining',
    'Week 5': 'Mixed expenses',
  };

  final Map<String, double> yearlyData = {
    'Jan': 5000.0,
    'Feb': 4500.0,
    'Mar': 6000.0,
    'Apr': 5500.0,
    'May': 7000.0,
    'Jun': 6500.0,
    'Jul': 7500.0,
    'Aug': 8000.0,
    'Sep': 7200.0,
    'Oct': 6800.0,
    'Nov': 8500.0,
    'Dec': 9000.0,
  };

  final Map<String, String> yearlyDescriptions = {
    'Jan': 'New year expenses, bills',
    'Feb': 'Regular expenses',
    'Mar': 'Spring shopping',
    'Apr': 'Travel, entertainment',
    'May': 'Summer prep, shopping',
    'Jun': 'Vacation expenses',
    'Jul': 'Summer activities',
    'Aug': 'Back to school',
    'Sep': 'Fall season expenses',
    'Oct': 'Festival expenses',
    'Nov': 'Holiday shopping start',
    'Dec': 'Holiday season expenses',
  };

  List<FlSpot> _getSpots() {
    Map<String, double> data;
    if (_selectedInterval == 0) {
      data = weeklyData;
    } else if (_selectedInterval == 1) {
      data = monthlyData;
    } else {
      data = yearlyData;
    }
    return data.entries
        .map(
          (e) => FlSpot(data.keys.toList().indexOf(e.key).toDouble(), e.value),
        )
        .toList();
  }

  Map<String, double> _getCurrentData() {
    if (_selectedInterval == 0) {
      return weeklyData;
    } else if (_selectedInterval == 1) {
      return monthlyData;
    } else {
      return yearlyData;
    }
  }

  Map<String, String> _getCurrentDescriptions() {
    if (_selectedInterval == 0) {
      return weeklyDescriptions;
    } else if (_selectedInterval == 1) {
      return monthlyDescriptions;
    } else {
      return yearlyDescriptions;
    }
  }

  String _getPeriodLabel() {
    if (_selectedInterval == 0) {
      return 'Weekly';
    } else if (_selectedInterval == 1) {
      return 'Monthly';
    } else {
      return 'Yearly';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Summary')),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Center(
                child: ToggleButtons(
                  isSelected: List.generate(
                    3,
                    (index) => index == _selectedInterval,
                  ),
                  onPressed: (index) {
                    setState(() {
                      _selectedInterval = index;
                    });
                  },
                  children: const [
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 16.0),
                      child: Text('Weekly'),
                    ),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 16.0),
                      child: Text('Monthly'),
                    ),
                    Padding(
                      padding: EdgeInsets.symmetric(horizontal: 16.0),
                      child: Text('Yearly'),
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 24),
              Text(
                'Total Summary',
                style: Theme.of(context).textTheme.headlineSmall,
              ),
              const SizedBox(height: 16),
              Card(
                child: ListTile(
                  title: const Text('Total Income'),
                  trailing: Text(
                    '+ \$5000.00',
                    style: const TextStyle(color: Colors.green, fontSize: 18),
                  ),
                ),
              ),
              Card(
                child: ListTile(
                  title: const Text('Total Expense'),
                  trailing: Text(
                    '- \$${_getSpots().map((e) => e.y).reduce((a, b) => a + b).toStringAsFixed(2)}',
                    style: const TextStyle(color: Colors.red, fontSize: 18),
                  ),
                ),
              ),
              const SizedBox(height: 24),
              Text(
                'Expense Variation',
                style: Theme.of(context).textTheme.headlineSmall,
              ),
              const SizedBox(height: 16),
              SizedBox(
                height: 300,
                child: LineChart(
                  LineChartData(
                    gridData: FlGridData(show: false),
                    titlesData: FlTitlesData(
                      leftTitles: AxisTitles(
                        sideTitles: SideTitles(showTitles: false),
                      ),
                      bottomTitles: AxisTitles(
                        sideTitles: SideTitles(
                          showTitles: true,
                          getTitlesWidget: (value, meta) {
                            List<String> titles;
                            if (_selectedInterval == 0) {
                              titles = weeklyData.keys.toList();
                            } else if (_selectedInterval == 1) {
                              titles = monthlyData.keys.toList();
                            } else {
                              titles = yearlyData.keys.toList();
                            }
                            final int index = value.toInt();
                            if (index < 0 || index >= titles.length) {
                              return const SizedBox.shrink();
                            }
                            return SideTitleWidget(
                              axisSide: meta.axisSide,
                              child: Text(titles[index]),
                            );
                          },
                          reservedSize: 30,
                          interval: 1,
                        ),
                      ),
                      topTitles: AxisTitles(
                        sideTitles: SideTitles(showTitles: false),
                      ),
                      rightTitles: AxisTitles(
                        sideTitles: SideTitles(showTitles: false),
                      ),
                    ),
                    borderData: FlBorderData(
                      show: true,
                      border: Border.all(
                        color: const Color(0xff37434d),
                        width: 1,
                      ),
                    ),
                    minX: 0,
                    maxX: _getSpots().length.toDouble() - 1,
                    minY: 0,
                    maxY:
                        _getSpots()
                            .map((e) => e.y)
                            .reduce((a, b) => a > b ? a : b) *
                        1.2,
                    lineBarsData: [
                      LineChartBarData(
                        spots: _getSpots(),
                        isCurved: true,
                        color: Colors.red,
                        barWidth: 5,
                        isStrokeCapRound: true,
                        dotData: FlDotData(show: false),
                        belowBarData: BarAreaData(
                          show: true,
                          color: Colors.red.withOpacity(0.3),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 32),
              Text(
                '${_getPeriodLabel()} Breakdown',
                style: Theme.of(context).textTheme.headlineSmall,
              ),
              const SizedBox(height: 16),
              Card(
                elevation: 2,
                child: Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: Column(
                    children: [
                      // Table Header
                      Container(
                        decoration: BoxDecoration(
                          color: Theme.of(context).colorScheme.primaryContainer,
                          borderRadius: const BorderRadius.vertical(
                            top: Radius.circular(8),
                          ),
                        ),
                        child: Padding(
                          padding: const EdgeInsets.symmetric(
                            vertical: 12.0,
                            horizontal: 16.0,
                          ),
                          child: Row(
                            children: [
                              Expanded(
                                flex: 2,
                                child: Text(
                                  'Period',
                                  style: Theme.of(context).textTheme.titleMedium
                                      ?.copyWith(fontWeight: FontWeight.bold),
                                ),
                              ),
                              Expanded(
                                flex: 3,
                                child: Text(
                                  'Description',
                                  style: Theme.of(context).textTheme.titleMedium
                                      ?.copyWith(fontWeight: FontWeight.bold),
                                ),
                              ),
                              Expanded(
                                flex: 2,
                                child: Text(
                                  'Amount',
                                  textAlign: TextAlign.right,
                                  style: Theme.of(context).textTheme.titleMedium
                                      ?.copyWith(fontWeight: FontWeight.bold),
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                      const Divider(height: 1),
                      // Table Rows
                      ..._buildTableRows(),
                      const Divider(height: 1, thickness: 2),
                      // Total Row
                      Padding(
                        padding: const EdgeInsets.symmetric(
                          vertical: 12.0,
                          horizontal: 16.0,
                        ),
                        child: Row(
                          children: [
                            Expanded(
                              flex: 2,
                              child: Text(
                                'Total',
                                style: Theme.of(context).textTheme.titleMedium
                                    ?.copyWith(fontWeight: FontWeight.bold),
                              ),
                            ),
                            const Expanded(flex: 3, child: SizedBox()),
                            Expanded(
                              flex: 2,
                              child: Text(
                                '\$${_calculateTotal().toStringAsFixed(2)}',
                                textAlign: TextAlign.right,
                                style: Theme.of(context).textTheme.titleMedium
                                    ?.copyWith(
                                      fontWeight: FontWeight.bold,
                                      color: Colors.red,
                                    ),
                              ),
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 16),
            ],
          ),
        ),
      ),
    );
  }

  List<Widget> _buildTableRows() {
    final data = _getCurrentData();
    final descriptions = _getCurrentDescriptions();
    final List<Widget> rows = [];

    data.forEach((key, value) {
      rows.add(
        Padding(
          padding: const EdgeInsets.symmetric(vertical: 12.0, horizontal: 16.0),
          child: Row(
            children: [
              Expanded(
                flex: 2,
                child: Text(key, style: Theme.of(context).textTheme.bodyLarge),
              ),
              Expanded(
                flex: 3,
                child: Text(
                  descriptions[key] ?? '',
                  style: Theme.of(
                    context,
                  ).textTheme.bodyMedium?.copyWith(color: Colors.grey[600]),
                ),
              ),
              Expanded(
                flex: 2,
                child: Text(
                  '\$${value.toStringAsFixed(2)}',
                  textAlign: TextAlign.right,
                  style: Theme.of(
                    context,
                  ).textTheme.bodyLarge?.copyWith(color: Colors.red),
                ),
              ),
            ],
          ),
        ),
      );
      if (data.keys.last != key) {
        rows.add(const Divider(height: 1));
      }
    });

    return rows;
  }

  double _calculateTotal() {
    final data = _getCurrentData();
    return data.values.fold(0.0, (sum, value) => sum + value);
  }
}
