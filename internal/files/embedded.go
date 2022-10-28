package files

import "embed"

//go:embed fs
var efs embed.FS

var (
	appP    = "fs/app/"
	hiveP   = "fs/hive/"
	navP    = "fs/navigation/"
	utilP   = "fs/util/"
	viewsP  = "fs/views/"
	dialogP = "fs/dialog/"
	cubitP  = "fs/cubit/"
	modelP  = "fs/models/"
)

var EmbeddedFsPaths = struct {
	Main                 string
	// App                  string
	// AppConfig            string
	// AppError             string
	// AppErrorG            string
	// AppResult            string
	// AppSettings          string
	// AppTheme             string
	// AppTrace             string
	// HiveAppSettings      string
	// HiveAppSettingsG     string
	// HiveBoxes            string
	// NavMain              string
	// UtilListenable       string
	// UtilPrimativeWrapper string
	// SplashView           string
	// SplashCtlr           string
	// ViewView             string
	// ViewCtlr             string
	// DialogView           string
	// DialogCtlr           string
	// Cubit                string
	// CubitState           string
	// Navigator            string
	// Model                string
}{
	Main:                 "fs/main.dart",
	// App:                  appP + "app.dart",
	// AppConfig:            appP + "app_config.dart",
	// AppError:             appP + "app_error.dart",
	// AppErrorG:            appP + "app_error.g.dart",
	// AppResult:            appP + "app_result.dart",
	// AppTheme:             appP + "app_theme.dart",
	// AppSettings:          appP + "app_settings.dart",
	// AppTrace:             appP + "app_trace.dart",
	// HiveAppSettings:      hiveP + "h_app_settings.dart",
	// HiveAppSettingsG:     hiveP + "h_app_settings.g.dart",
	// HiveBoxes:            hiveP + "h_boxes.dart",
	// NavMain:              navP + "n_main.dart",
	// UtilListenable:       utilP + "u_app_listenable.dart",
	// UtilPrimativeWrapper: utilP + "u_primative_wrapper.dart",
	// SplashView:           viewsP + "v_splash.dart",
	// SplashCtlr:           viewsP + "c_splash.dart",
	// ViewView:             viewsP + "v_view.dart",
	// ViewCtlr:             viewsP + "c_view.dart",
	// DialogView:           dialogP + "d_dialog.dart",
	// DialogCtlr:           dialogP + "c_dialog.dart",
	// Cubit:                cubitP + "cubit.dart",
	// CubitState:           cubitP + "state.dart",
	// Navigator:            navP + "nav.dart",
	// Model:                modelP + "model.dart",
}
