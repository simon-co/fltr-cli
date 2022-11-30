import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_config.dart';
import 'package:PROJECT_NAME/src/app/app.dart';

void main() async {
  await AppConfig.startUp();
  runApp(App());
}
