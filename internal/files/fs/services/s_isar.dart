import 'package:flutter/material.dart';
import 'package:isar/isar.dart';
import 'package:PROJECT_NAME/src/models/m_settings.dart';

class IsarService {
  late Future<Isar> db;

  IsarService(){
    db = openDB();
  }

  Future<AppSettings?> getAppSettings() async {
    final isar = await db;
    return await isar.appSettings.get(1);
  }

  Future<void> saveAppSettings(AppSettings settings) async {
    final isar = await db;
    await isar.writeTxn(() async {
      await isar.appSettings.put(settings);
    });
  }

  Future<Isar> openDB() async {
    if (Isar.instanceNames.isEmpty) {
      return await Isar.open([AppSettingsSchema]);
    } else {
      return Future.value(Isar.getInstance());
    }
  }
}
