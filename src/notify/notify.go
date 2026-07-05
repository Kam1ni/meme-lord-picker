package notify

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func Notify(title string, message string) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to dbus\n%s", err.Error()))
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0), "", title, message, []string{}, map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		panic(fmt.Sprintf("Failed to notify copy to clipboard\n%s", call.Err.Error()))
	}
}
