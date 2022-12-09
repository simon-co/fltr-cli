import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_theme.dart';
import 'package:PROJECT_NAME/src/routing/router.dart';

class App extends StatefulWidget {
  const App({super.key});

  @override
  State<App> createState() => _AppState();
}

class _AppState extends State<App> {
  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder(
        valueListenable: AppTheme.notifier,
        builder: (_, mode, __) {
          return MaterialApp.router(
            debugShowCheckedModeBanner: false,             
            theme: ThemeData.light(),
            darkTheme: ThemeData.dark(),
            themeMode: mode,
            routerConfig: AppRouter.router,
          );
        });
  }
}
