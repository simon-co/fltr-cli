import 'package:flutter/material.dart';
import 'package:PROJECT_NAME/src/app/app_theme.dart';

part 'c_app_settings.dart';

class AppSettingsDialog extends StatefulWidget {
  const AppSettingsDialog({super.key});

  @override
  State<AppSettingsDialog> createState() => _AppSettingsDialogState();
}

class _AppSettingsDialogState extends State<AppSettingsDialog> {
  @override
  Widget build(BuildContext context) {
    return const SimpleDialog(
      title: Text(
        "Settings",
        textAlign: TextAlign.center,
      ),
      contentPadding: EdgeInsets.all(16),
      titlePadding: EdgeInsets.all(16),
      children: [
        _SettingsTile("Theme"),
        _ThemeButton()
      ],
    );
  }
}

class _SettingsTile extends StatelessWidget {
  final String title;
  const _SettingsTile(this.title, {super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisSize: MainAxisSize.min,
      children: [
        const Divider(
          height: 12,
        ),
        Text(title)
      ],
    );
  }
}

class _ThemeButton extends StatelessWidget {
  const _ThemeButton({super.key});

  @override
  Widget build(BuildContext context) {
    return ValueListenableBuilder(
        valueListenable: AppTheme.notifier,
        builder: (_, mode, __) {
          return DropdownButton<ThemeMode>(items: const [
            DropdownMenuItem(
              value: ThemeMode.system,
              child: Text("System Theme"),
            ),
            DropdownMenuItem(
              value: ThemeMode.light,
              child: Text("Light Theme"),
            ),
            DropdownMenuItem(
              value: ThemeMode.dark,
              child: Text("Dark Theme"),
            )
          ], value: mode, onChanged: AppTheme.setTheme);
        });
  }
}

