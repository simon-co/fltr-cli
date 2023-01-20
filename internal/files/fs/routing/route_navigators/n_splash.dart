import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/splash/v_splash.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';


class SplashRouteNavigator extends StatelessWidget {
  SplashRouteNavigator(this.params, {super.key});
  static const route = "/welcome";
  Map<String, String> params;
  static final navKey = GlobalKey<NavigatorState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Navigator(
        key: navKey,
        initialRoute: SplashView.route,
        onGenerateRoute: (RouteSettings settings) {
          late Widget view;
          switch (settings.name) {
            case SplashView.route:
              view = SplashView();
              break;
          }
          return MaterialPageRoute(
            settings: settings,
            builder: (context) => view,
          );
        },
      ),
    );
  }

  static AppError? toSplashView() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushReplacementNamed(SplashView.route);
    } else {
      return AppError(AppErrorCode.e500, "navState is null");
    }
    return null;
  }

  static AppError? pushSplashView() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushNamed(SplashView.route);
    } else {
      return AppError(AppErrorCode.e500, "navState is null");
    }
    return null;
  }

  static AppError? pop() {
    final navState = navKey.currentState;
    if (navState != null && navState.canPop()) {
      navState.pop();
      return null;
    } else {
      return AppError(AppErrorCode.e500, "navState can't be popped");
    }
  }
}
