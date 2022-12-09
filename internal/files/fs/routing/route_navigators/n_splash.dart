import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/splash/v_splash.dart';

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
}
