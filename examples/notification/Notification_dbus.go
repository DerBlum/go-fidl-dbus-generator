// Code generated by Go-Fidl-Generator. DO NOT EDIT.
// see https://github.com/SourceFellows/go-fidl-dbus-generator
package notification

import (
	"context"
	"github.com/godbus/dbus/v5"
)

func NewNotificationsSender(dest, path string) (*notificationsSender, error) {

	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, err
	}

	broadcastMatchOptions := []dbus.MatchOption{
		dbus.WithMatchObjectPath(dbus.ObjectPath(path)),
		dbus.WithMatchInterface("org.freedesktop.Notifications"),
	}

	return &notificationsSender{
		dbusConnection:        conn,
		destination:           dest,
		path:                  dbus.ObjectPath(path),
		broadcastMatchOptions: broadcastMatchOptions,
	}, nil
}

type notificationsSender struct {
	dbusConnection        *dbus.Conn
	destination           string
	path                  dbus.ObjectPath
	broadcastMatchOptions []dbus.MatchOption
}

func (impl *notificationsSender) Close() error {
	return impl.dbusConnection.Close()
}

func (impl *notificationsSender) Notify(ctx context.Context, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints Hint, expire_timeout int32) (uint8, error) {

	var result uint8

	err := impl.dbusConnection.Object(impl.destination, impl.path).
		CallWithContext(ctx, "org.freedesktop.Notifications.Notify", 0, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout).
		Store(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
