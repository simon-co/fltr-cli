package files

import "embed"

//go:embed fs
var efs embed.FS

var (
	appP    = "fs/app/"
	navP    = "fs/navigation/"
	utilP   = "fs/util/"
	viewsP  = "fs/views/"
	dialogP = "fs/dialog/"
	stateP  = "fs/state/"
	modelP  = "fs/models/"
)

var EmbeddedFsPaths = struct {
	Main         string
	App          string
	AppError     string
	AppErrorG    string
	AppResult    string
	AppTheme     string
	AppCalltrace string
	RootNav      string
	// UtilListenable       string
	// UtilPrimativeWrapper string
	SplashView string
	SplashCtlr string
	StartView  string
	StartCtlr  string
	ViewView   string
	ViewCtlr   string
	// DialogView           string
	// DialogCtlr           string
	// CubitState           string
	// Model                string
}{
	Main:         "fs/main.dart",
	App:          appP + "app.dart",
	AppError:     appP + "app_error.dart",
	AppErrorG:    appP + "app_error.g.dart",
	AppResult:    appP + "app_result.dart",
	AppTheme:     appP + "app_theme.dart",
	AppCalltrace: appP + "app_trace.dart",
	RootNav:      navP + "n_root.dart",
	// UtilListenable:       utilP + "u_app_listenable.dart",
	// UtilPrimativeWrapper: utilP + "u_primative_wrapper.dart",
	SplashView: viewsP + "v_splash.dart",
	SplashCtlr: viewsP + "c_splash.dart",
	StartView:  viewsP + "v_start.dart",
	StartCtlr:  viewsP + "c_start.dart",
	ViewView:   viewsP + "v_view.dart",
	ViewCtlr:   viewsP + "c_view.dart",
	// DialogView:           dialogP + "d_dialog.dart",
	// DialogCtlr:           dialogP + "c_dialog.dart",
	// Cubit:                cubitP + "cubit.dart",
	// CubitState:           cubitP + "state.dart",
	// Navigator:            navP + "nav.dart",
	// Model:                modelP + "model.dart",
}
