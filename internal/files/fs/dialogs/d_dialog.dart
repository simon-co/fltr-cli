import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_theme.dart';

part 'CONTROLLER_FILENAME';

class CLASS_NAMEDialog extends StatefulWidget {
  const CLASS_NAMEDialog({super.key});

  @override
  State<CLASS_NAMEDialog> createState() => _CLASS_NAMEDialogState();
}

class _CLASS_NAMEDialogState extends State<CLASS_NAMEDialog> {
  final ctrl = _CLASS_NAMEDialogCtlr();

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
    return const SimpleDialog(
      title: Text(
        "CLASS_NAMEDialog",
        textAlign: TextAlign.center,
      ),
      contentPadding: EdgeInsets.all(16),
      titlePadding: EdgeInsets.all(16),
      children: [],
    );
  }
}
