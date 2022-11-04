// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'app_error.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

AppError _$AppErrorFromJson(Map<String, dynamic> json) => AppError(
      $enumDecode(_$AppErrorCodeEnumMap, json['errorCode']),
      json['message'] as String,
      json['data'] ?? "",
    )
      ..caller = json['caller'] as String
      ..className = json['className'] as String
      ..filename = json['filename'] as String
      ..line = json['line'] as String
      ..column = json['column'] as String
      ..trace =
          (json['trace'] as List<dynamic>).map((e) => e as String).toList();

Map<String, dynamic> _$AppErrorToJson(AppError instance) => <String, dynamic>{
      'errorCode': _$AppErrorCodeEnumMap[instance.errorCode]!,
      'message': instance.message,
      'caller': instance.caller,
      'className': instance.className,
      'filename': instance.filename,
      'line': instance.line,
      'column': instance.column,
      'trace': instance.trace,
      'data': instance.data,
    };

const _$AppErrorCodeEnumMap = {
  AppErrorCode.e404: 'e404',
  AppErrorCode.e500: 'e500',
  AppErrorCode.e409: 'e409',
};

