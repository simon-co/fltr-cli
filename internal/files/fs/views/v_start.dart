import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';
import 'package:PROJECT_NAME/src/dialogs/app_settings/d_app_settings.dart';

part 'c_start.dart';

class StartView extends StatefulWidget {
  StartView(this.params, {super.key});
  static const route = "start";
  final Map<String, String> params;

  @override
  State<StartView> createState() => _StartViewState();
}

class _StartViewState extends State<StartView> {
  final ctrl = _StartViewCtlr();

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
    return Center(child: Text("Start View"));
  }
}
