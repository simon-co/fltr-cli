import 'package:flutter/widgets.dart';
import 'package:go_router/go_router.dart';
import 'package:PROJECT_NAME/src/views/splash/v_splash.dart';
import 'package:PROJECT_NAME/src/views/start/v_start.dart';

class RootNav {
  static GoRouter router = GoRouter(initialLocation: SplashView.route, routes: [
    GoRoute(path: SplashView.route, builder: (context, state) => SplashView()),
    GoRoute(path: StartView.route, builder: (conetext, state) => StartView()),
  ]);
}
