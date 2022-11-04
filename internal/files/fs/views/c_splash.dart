part of 'v_splash.dart';

class SplashViewController {
  static SplashViewController? _instance;

  SplashViewController._internal();

  factory SplashViewController() {
    _instance ??= SplashViewController._internal();
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

