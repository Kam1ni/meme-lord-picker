package windowrules

type desktopEnvironment interface {
	setWindowPositionRule()
}

var desktopEnvironments = map[string]desktopEnvironment{
	"Hyprland": hyprland{},
}
