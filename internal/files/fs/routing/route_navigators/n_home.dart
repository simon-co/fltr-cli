import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/start/v_start.dart';

class HomeRouteNavigator extends StatelessWidget {
  HomeRouteNavigator(this.params, {super.key});
  static const route = "/home";
  Map<String, String> params;
  static final navKey = GlobalKey<NavigatorState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Navigator(
        key: navKey,
        initialRoute: params["view"] ?? StartView.route,
        onGenerateRoute: (RouteSettings settings) {
          late Widget view;
          switch (settings.name) {
            case StartView.route:
              view = StartView(params);
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
