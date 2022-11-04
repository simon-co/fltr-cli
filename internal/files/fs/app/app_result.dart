import 'package:PROJECT_NAME/src/app/app_error.dart';

class AppResult<T> {
  T? data;
  AppError? error;
  List<AppError> errorList;
  AppResult({this.data, this.error, this.errorList = const []});
}
