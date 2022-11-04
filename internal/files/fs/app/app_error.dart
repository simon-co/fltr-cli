import 'dart:convert';
import 'package:json_annotation/json_annotation.dart';
import 'package:PROJECT_NAME/src/app/app_calltrace.dart';

part 'app_error.g.dart';

enum AppErrorCode { e404, e500, e409 }

@JsonSerializable()
class AppError {
  AppErrorCode errorCode;
  String message;
  String caller;
  String className;
  String filename;
  String line;
  String column;
  List<String> trace;
  dynamic data;

  AppError._internal(
      {required this.errorCode,
      required this.message,
      required this.caller,
      required this.className,
      required this.filename,
      required this.line,
      required this.column,
      this.data})
      : trace = [
          "\n{\"Filename\": \"$filename\" \"className\": \"$className\" \"caller\": \"$caller\" \"line\": \"$line\" \"column\": \"$column\"}\n"
        ];

  factory AppError(AppErrorCode errorCode, String message,
      [dynamic data = ""]) {
    final trace = Calltrace(depth: 1);
    return AppError._internal(
        errorCode: errorCode,
        message: message,
        caller: trace.callerName,
        className: trace.className,
        filename: trace.filename,
        line: trace.line,
        column: trace.column,
        data: data);
  }

  factory AppError.parse(
    dynamic error,
  ) {
    final trace = Calltrace(depth: 1);
    if (error is AppError) {
      error.addTrace(trace);
      return error;
    } else {
      return AppError._internal(
          errorCode: AppErrorCode.e500,
          message: "Unexpected Error",
          filename: trace.filename,
          className: trace.className,
          caller: trace.callerName,
          line: trace.line,
          column: trace.column,
          data: error);
    }
  }

  Map<String, dynamic> toJson() => _$AppErrorToJson(this);

  addTrace(Calltrace trace) {
    this.trace.add(
        "\n{\"Filename\": \"${trace.filename}\" \"className\": \"${trace.className}\" \"caller\": \"${trace.callerName}\" \"line\": \"${trace.line}\" \"column\": \"${trace.column}\"}\n");
  }

  printToConsole() {
    const encoder = JsonEncoder.withIndent('  ');
    print(encoder.convert(toJson()));
  }
}
