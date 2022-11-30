part of 'DIALOG_FILENAME';

class _CLASS_NAMEDialogCtlr {
  static _CLASS_NAMEDialogCtlr? _instance;

  _CLASS_NAMEDialogCtlr._internal();

  factory _CLASS_NAMEDialogCtlr(){
    _instance ??= _CLASS_NAMEDialogCtlr._internal();
    return _instance!;
  }

  init(){}

  dispose(){
    _instance = null;
  }
}

