part of 'v_splash.dart';

class _SplashViewCtrl {
  static _SplashViewCtrl? _instance;

  _SplashViewCtrl._internal();

  factory _SplashViewCtrl() {
    _instance ??= _SplashViewCtrl._internal();
    return _instance!;
  }

  dispose() {
    _instance = null;
  }

  init() {}

  navigateAfterDelay(BuildContext context, int delay) async {
    Future.delayed(
      Duration(milliseconds: delay),
      () => context.go(HomeRouteNavigator.route)
    );
  }
}

