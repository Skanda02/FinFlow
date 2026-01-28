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

  final Map<String, double> monthlyData = {
    'Week 1': 800.0,
    'Week 2': 1200.0,
    'Week 3': 1000.0,
    'Week 4': 1500.0,
    'Week 5': 1300.0,
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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Summary')),
      body: Padding(
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
            Expanded(
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
          ],
        ),
      ),
    );
  }
}
