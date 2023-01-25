import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';

part 'CONTROLLER_FILE_NAME';

class CLASS_NAMEView extends StatefulWidget {
  const CLASS_NAMEView(this.params, {super.key});
  static const route = "ROUTE";
  final Map<String, String> params;

  @override
  State<CLASS_NAMEView> createState() => _CLASS_NAMEViewState();
}

class _CLASS_NAMEViewState extends State<CLASS_NAMEView> {
  final ctrl = _CLASS_NAMEViewCtlr();

  @override
  void initState() {
    super.initState();
    ctrl.init();
  }

  @override
  void dispose() {
    super.dispose();
    ctrl.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Center(child: Text("CLASS_NAME View"));
  }
}
