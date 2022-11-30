import 'package:flutter/material.dart';
import 'package:isar/isar.dart';
import 'package:PROJECT_NAME/src/services/s_isar.dart';

part 'm_settings.g.dart';

@collection
class AppSettings {
  AppSettings();

  Id? id;

  @Enumerated(EnumType.name)
  ThemeMode? theme;

  AppSettings.standard() {
    id = 1;
    theme = ThemeMode.system;
  }

  static Future<AppSettings> load() async {
    final service = IsarService();
    AppSettings? settings = await service.getAppSettings();
    if (settings != null) {
      return settings;
    } else {
      return AppSettings.standard();
    }
  }

  static Future<void> saveTheme(ThemeMode mode) async {
    AppSettings settings = await AppSettings.load();
    settings.theme = mode;
    settings.save();
  }

  Future<void> save() async {
    final service = IsarService();
    await service.saveAppSettings(this);
  }
}
