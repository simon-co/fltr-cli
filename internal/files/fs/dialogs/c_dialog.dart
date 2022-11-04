part of 'DIALOG_FILENAME';

class _DIALOG_CONTROLLER_NAME {
  static _DIALOG_CONTROLLER_NAME? _instance;

  _DIALOG_CONTROLLER_NAME._internal();

  factory _DIALOG_CONTROLLER_NAME(){
    _instance ??= _DIALOG_CONTROLLER_NAME._internal();
    return _instance!;
  }

  init(){}

  dispose(){
    _instance = null;
  }
}

