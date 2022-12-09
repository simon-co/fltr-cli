import 'package:go_router/go_router.dart';
import 'package:PROJECT_NAME/src/routing/route_navigators/n_home.dart';
import 'package:PROJECT_NAME/src/routing/route_navigators/n_splash.dart';

class AppRouter {
  static GoRouter router = GoRouter(
    initialLocation: SplashRouteNavigator.route, 
    routes: [
    GoRoute(
        name: SplashRouteNavigator.route,
        path: SplashRouteNavigator.route,
        builder: (context, state) => SplashRouteNavigator(state.queryParams)),
    GoRoute(
        name: HomeRouteNavigator.route,
        path: HomeRouteNavigator.route,
        builder: (conext, state) => HomeRouteNavigator(state.queryParams)),
  ]);
}
