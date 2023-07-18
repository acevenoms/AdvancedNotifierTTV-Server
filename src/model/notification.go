package model

// Abstract notif
type INotification interface {
	Value(string) string
}

type TwitchNotification struct {
	Data map[string]string
}

func (tn TwitchNotification) Value(key string) string {
	return tn.Data[key]
}
