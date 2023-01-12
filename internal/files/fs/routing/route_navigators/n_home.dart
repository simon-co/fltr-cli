import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';
import 'package:PROJECT_NAME/src/dialogs/app_settings/d_app_settings.dart';
import 'package:PROJECT_NAME/src/views/start/v_start.dart';

class HomeRouteNavigator extends StatelessWidget {
  HomeRouteNavigator(this.params, {super.key});
  static const route = "/home";
  Map<String, String> params;
  static final navKey = GlobalKey<NavigatorState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: _AppBar.build(context),
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
          return HomeRouteNavigatorPage(
              view: view,
              settings: settings,
          );
        },
      ),
    );
  }
}

class HomeRouteNavigatorPage extends PageRouteBuilder {
  final Widget view;
  final RouteSettings settings;
  HomeRouteNavigatorPage({required this.view, required this.settings})
      : super(
            settings: settings,
            pageBuilder: (
              BuildContext context,
              Animation animation,
              Animation secondaryAnimation,
            ) =>
                view,
            transitionsBuilder:
                (context, animation, secondaryAnimation, child) {
              return FadeTransition(
                opacity: animation,
                child: SlideTransition(
                  position: Tween<Offset>(
                          begin: const Offset(0, 0.25), end: Offset.zero)
                      .animate(animation),
                  child: SlideTransition(
                    position:
                        Tween<Offset>(begin: Offset.zero, end: Offset(0, -1))
                            .animate(secondaryAnimation),
                    child: child,
                  ),
                ),
              );
            });
}

class _AppBar {
  static AppBar build(BuildContext context) {
    return AppBar(
      backgroundColor: Colors.transparent,
      elevation: 0,
      actions: <Widget>[
        IconButton(
            color: Colors.grey,
            hoverColor: Colors.transparent,
            onPressed: () => openSettings(context),
            icon: const Icon(Icons.settings))
      ],
    );
  }

  static Future<AppResult<void>> openSettings(BuildContext context) async {
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
