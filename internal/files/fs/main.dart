import 'package:PROJECT_NAME/src/app/app.dart';
import 'package:PROJECT_NAME/src/app/app_config.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:flutter/material.dart';

void main() async {
  final startConfig = await AppConfig.onStartUp();
  if (startConfig.error != null) {
    AppError error = startConfig.error!;
    error.log();
  }
  runApp(App());
}
