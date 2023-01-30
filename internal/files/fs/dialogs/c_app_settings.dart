part of 'd_app_settings.dart';

class _AppSettingsDialogCtrl {
  static _AppSettingsDialogCtrl? _instance;

  _AppSettingsDialogCtrl._internal();

  factory _AppSettingsDialogCtrl() {
    _instance ??= _AppSettingsDialogCtrl._internal();
    return _instance!;
  }

  init() {}

  dispose() {
    _instance = null;
  }

  setThemeMode(ThemeMode? mode) async {
    if (mode != null) {
      await AppTheme.setTheme(mode);
      await AppSettings.saveTheme(mode);
    }
  }
}
