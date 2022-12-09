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
  final ctlr = _StartViewCtlr();

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
    ThemeData theme = Theme.of(context);
    return Scaffold(
      backgroundColor: theme.backgroundColor,
      appBar: _AppBar.build(context),
      body: Center(child: Text("Start View")),
    );
  }
}

class _AppBar {
  static AppBar build(BuildContext context) {
    final ctlr = _StartViewCtlr();
    return AppBar(
      actions: <Widget>[
        IconButton(
            onPressed: () => ctlr.openSettings(context),
            icon: const Icon(Icons.settings))
      ],
    );
  }
}
