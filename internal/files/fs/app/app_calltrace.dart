class Calltrace {
  final int depth;

  late final String filename;
  late final String callerName;
  late final String className;
  late final String line;
  late final String column;

  Calltrace({this.depth = 0}){
    _parse();
  }

  _parse(){
    final trace = StackTrace.current;

    final frame = trace.toString().split("\n")[2 + depth];

    final fnRe = RegExp(" (_?[A-Z][A-Za-z0-9]*\._?[A-Z a-z0-9]*) ");
    final fnMatch = fnRe.firstMatch(frame);

    final fnSplit = fnMatch!.group(1)?.split(".");
    className = fnSplit![0];
    callerName = fnSplit[1];

    final pkgRe = RegExp(":(.*.dart):([0-9]*):([0-9]*)");
    final pkgMatch = pkgRe.firstMatch(frame)!;

    filename = pkgMatch.group(1)!;
    line = pkgMatch.group(2)!;
    column = pkgMatch.group(3)!;
  }
}
