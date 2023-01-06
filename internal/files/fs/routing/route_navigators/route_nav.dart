import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/VIEW_PATH';

class NAV_CLASSNAME extends StatelessWidget {
  NAV_CLASSNAME(this.params, {super.key});
  static const route = "ROUTE";
  Map<String, String> params;
  static final navKey = GlobalKey<NavigatorState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Navigator(
        key: navKey,
        initialRoute: VIEW_CLASS_NAME.route,
        onGenerateRoute: (RouteSettings settings) {
          late Widget view;
          switch (settings.name) {
            case VIEW_CLASS_NAME.route:
              view = VIEW_CLASS_NAME();
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
