part of 'v_start.dart';

class _StartViewCtrl {
  static _StartViewCtrl? _instance;

  _StartViewCtrl._internal();

  factory _StartViewCtrl() {
    _instance ??= _StartViewCtrl._internal();
    return _instance!;
  }

  dispose() {
    _instance = null;
  }

  init() {}

  Future<AppResult<void>> openSettings(BuildContext context) async {
    final result = AppResult<void>();
    try {
      await showDialog(
          context: context, builder: (context) => const AppSettingsDialog());
    } catch (err) {
      result.error = AppError.parse(err);
      result.error!.printToConsole();
    }
    return result;
  }
}
