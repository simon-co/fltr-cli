import 'package:PROJECT_NAME/material.dart';

class AppTheme {
  static ThemeMode themeMode = ThemeMode.system;

  static final ValueNotifier<ThemeMode> notifier = ValueNotifier(themeMode);

  static dark() => ThemeData(
      colorScheme: ColorScheme.fromSwatch(
          primarySwatch: Colors.orange, brightness: Brightness.dark),
      iconTheme: IconThemeData(color: Colors.orange[500]),
      toggleableActiveColor: Colors.orange[500]);

  static setTheme(ThemeMode? mode) {
    if (mode != null) {
      themeMode = mode;
      notifier.value = themeMode;
    }
  }
}

