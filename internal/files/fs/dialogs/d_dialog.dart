import 'package:flutter/material.dart';
import 'package:PRJOECT_NAME/src/app/app_theme.dart';

part 'CONTROLLER_FILENAME';

class DIALOG_CLASS_NAME extends StatefulWidget {
  const DIALOG_CLASS_NAME({super.key});

  @override
  State<DIALOG_CLASS_NAME> createState() => _DIALOG_CLASS_NAMEState();
}

class _DIALOG_CLASS_NAMEState extends State<DIALOG_CLASS_NAME> {
  final ctlr = CONTROLLER_CLASS_NAME();

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
    return const SimpleDialog(
      title: Text(
        DIALOG_TITLE,
        textAlign: TextAlign.center,
      ),
      contentPadding: EdgeInsets.all(16),
      titlePadding: EdgeInsets.all(16),
      children: [],
    );
  }
