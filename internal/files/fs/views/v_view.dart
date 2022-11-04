import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';
import 'package:PROJECT_NAME/src/dialogs/d_app_settings.dart';

part 'CONTROLLER_FILE_NAME';

class CLASS_NAMEView extends StatefulWidget {
  static const route = "VIEW_ROUTE";
  const VIEW_CLASS_NAME({super.key});

  @override
  State<CLASS_NAMEView> createState() => _VIEW_CLASS_NAMEState();
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
