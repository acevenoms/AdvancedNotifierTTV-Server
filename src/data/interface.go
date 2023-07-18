package data

import "github.com/acevenoms/AdvancedNotifierTTV-Server/src/model"

type IDataSource interface {
	InsertChannel(model.Channel) uint64
	SelectChannel(key string) model.Channel
	QueryChannels(map[string]string) []model.Channel

	InsertViewer(model.Viewer) uint64
	SelectViewer(key string) model.Viewer
	QueryViewers(map[string]string) []model.Viewer
}
