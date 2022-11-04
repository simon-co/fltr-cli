import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/views/start/v_start.dart';
import 'package:go_router/go_router.dart';

part 'c_splash.dart';

class SplashView extends StatefulWidget {
  static const route = "/splash";
  const SplashView({super.key});

  @override
  State<SplashView> createState() => _SplashViewState();
}

class _SplashViewState extends State<SplashView> {
  final ctlr = SplashViewController();

  @override
  void initState() {
    super.initState();
    ctlr.init();
    ctlr.navigateAfterDelay(context, 750);
  }

  @override
  void dispose() {
    super.dispose();
    ctlr.dispose();
  }

  @override
  Widget build(BuildContext context) {
    ThemeData theme = Theme.of(context);
    return Container(
        color: theme.backgroundColor,
        child: Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: <Widget>[
              Text("PROJECT_NAME", style: theme.textTheme.headline1),
              Icon(
                Icons.thumb_up_rounded,
                color: theme.textTheme.headline1!.color,
                size: 100,
              )
            ],
          ),
        ));
  }
}

