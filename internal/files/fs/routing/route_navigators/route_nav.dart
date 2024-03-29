import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/VIEW_PATH';
import 'package:PROJECT_NAME/src/app/app_error.dart';
import 'package:PROJECT_NAME/src/app/app_result.dart';
import 'package:PROJECT_NAME/src/dialogs/app_settings/d_app_settings.dart';

class NAV_CLASSNAME extends StatelessWidget {
  NAV_CLASSNAME(this.params, {super.key});
  static const route = "ROUTE";
  Map<String, String> params;
  static final navKey = GlobalKey<NavigatorState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: _AppBar.build(context),
      body: Navigator(
        key: navKey,
        initialRoute: params["view"] ?? VIEW_CLASS_NAME.route,
        onGenerateRoute: (RouteSettings settings) {
          late Widget view;
          switch (settings.name) {
            case VIEW_CLASS_NAME.route:
              view = VIEW_CLASS_NAME(params);
              break;
          }
          return NAV_CLASSNAMEPage(
            view: view,
            settings: settings,
          );
        },
      ),
    );
  }

  static AppError? toVIEW_CLASS_NAME() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushReplacementNamed(VIEW_CLASS_NAME.route);
    } else {
      return AppError(AppErrorCode.e500, "navState is null");
    }
    return null;
  }

  static AppError? pushVIEW_CLASS_NAME() {
    final navState = navKey.currentState;
    if (navState != null) {
      navState.pushNamed(VIEW_CLASS_NAME.route);
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

class NAV_CLASSNAMEPage extends PageRouteBuilder {
  final Widget view;
  final RouteSettings settings;
  NAV_CLASSNAMEPage({required this.view, required this.settings})
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
