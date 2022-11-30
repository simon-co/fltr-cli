import 'package:PROJECT_NAME/src/app/app_theme.dart';
import 'package:PROJECT_NAME/src/models/m_settings.dart';

class  AppConfig {
  static startUp() async {
    final settings = await AppSettings.load();
    AppTheme.setTheme(settings.theme);
  }
}
