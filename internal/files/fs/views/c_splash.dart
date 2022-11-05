part of 'v_splash.dart';

class _SplashViewCtlr {
  static _SplashViewCtlr? _instance;

  _SplashViewCtlr._internal();

  factory _SplashViewCtlr() {
    _instance ??= _SplashViewCtlr._internal();
    return _instance!;
  }

  dispose() {
    _instance = null;
  }

  init() {}

  navigateAfterDelay(BuildContext context, int delay) async {
    Future.delayed(
      Duration(milliseconds: delay),
      () => context.go(StartView.route)
    );
  }
}

