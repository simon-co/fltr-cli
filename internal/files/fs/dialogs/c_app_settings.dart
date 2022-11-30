part of 'd_app_settings.dart';

class _AppSettingsDialogCtlr {
  static _AppSettingsDialogCtlr? _instance;

  _AppSettingsDialogCtlr._internal();

  factory _AppSettingsDialogCtlr() {
    _instance ??= _AppSettingsDialogCtlr._internal();
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
