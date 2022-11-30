import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';

part 'CONTROLLER_FILE_NAME';

class CLASS_NAMEView extends StatefulWidget {
  static const route = "ROUTE";
  const CLASS_NAMEView({super.key});

  @override
  State<CLASS_NAMEView> createState() => _CLASS_NAMEViewState();
}

class _CLASS_NAMEViewState extends State<CLASS_NAMEView> {
  final ctlr = _CLASS_NAMEViewCtlr();

  @override
  void initState() {
    super.initState();
    ctlr.init();
  }

  @override
  void dispose() {
    super.dispose();
    ctlr.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Center(child: Text("CLASS_NAME View"));
  }
}
