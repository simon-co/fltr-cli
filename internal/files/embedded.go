package files

import "embed"

//go:embed fs
var efs embed.FS

var (
	appP      = "fs/app/"
	routingP  = "fs/routing/"
	utilP     = "fs/util/"
	viewsP    = "fs/views/"
	dialogP   = "fs/dialogs/"
	stateP    = "fs/state/"
	modelP    = "fs/models/"
	servicesP = "fs/services/"
	routeNavP = routingP + "route_navigators/"
)

var EmbeddedFsPaths = struct {
	Main               string
	App                string
	AppConfig          string
	AppError           string
	AppErrorG          string
	AppResult          string
	AppTheme           string
	AppCalltrace       string
	SettingsModel      string
	SettingsModelG     string
	IsarService        string
	Router             string
	SplashView         string
	SplashCtrl         string
	StartView          string
	StartCtrl          string
	ViewView           string
	ViewCtrl           string
	SettingsDialogView string
	SettingsDialogCtrl string
	DialogView         string
	DialogCtrl         string
	HomeNav            string
	SplashNav          string
	RouteNav           string
	Model              string
	IsarModel          string
	// UtilListenable       string
	// UtilPrimativeWrapper string
	// CubitState           string
}{
	Main:               "fs/main.dart",
	App:                appP + "app.dart",
	AppConfig:          appP + "app_config.dart",
	AppError:           appP + "app_error.dart",
	AppErrorG:          appP + "app_error.g.dart",
	AppResult:          appP + "app_result.dart",
	AppTheme:           appP + "app_theme.dart",
	AppCalltrace:       appP + "app_calltrace.dart",
	SettingsModel:      modelP + "m_settings.dart",
	SettingsModelG:     modelP + "m_settings.g.dart",
	IsarService:        servicesP + "s_isar.dart",
	Router:             routingP + "router.dart",
	SplashView:         viewsP + "v_splash.dart",
	SplashCtrl:         viewsP + "c_splash.dart",
	StartView:          viewsP + "v_start.dart",
	StartCtrl:          viewsP + "c_start.dart",
	ViewView:           viewsP + "v_view.dart",
	ViewCtrl:           viewsP + "c_view.dart",
	SettingsDialogView: dialogP + "d_app_settings.dart",
	SettingsDialogCtrl: dialogP + "c_app_settings.dart",
	DialogView:         dialogP + "d_dialog.dart",
	DialogCtrl:         dialogP + "c_dialog.dart",
	HomeNav:            routeNavP + "n_home.dart",
	SplashNav:          routeNavP + "n_splash.dart",
	RouteNav:           routeNavP + "route_nav.dart",
	Model:              modelP + "model.dart",
	IsarModel:          modelP + "isar_model.dart",
	// UtilListenable:       utilP + "u_app_listenable.dart",
	// UtilPrimativeWrapper: utilP + "u_primative_wrapper.dart",
	// Cubit:                cubitP + "cubit.dart",
	// CubitState:           cubitP + "state.dart",
	// Navigator:            navP + "nav.dart",
}
