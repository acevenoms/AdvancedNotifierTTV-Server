package api

import "github.com/acevenoms/AdvancedNotifierTTV-Server/src/model"

// Idempotent
func addChannel(channelLogin string) (model.Channel, error) {
	//TODO: Check if we already have it and return the existing channel

	//TODO: If new, check to make sure it's real with Twitch, then return it
	return model.Channel{Name: channelLogin}, nil
}

// Idempotent
func addViewer(userId string) (model.Viewer, error) {
	//TODO: Make sure they're valid an populate the return

	//TODO: If they're new we need more info somehow. Maybe mark them as needing onboarding
	return model.Viewer{Name: userId}, nil
}

// Idempotent
func addViewerToChannel(channelLogin string, userId string) (model.Channel, error) {
	channel, err := addChannel(channelLogin)
	if err != nil {
		return channel, err
	}

	viewer, err := addViewer(userId)
	if err != nil {
		return channel, err
	}

	channel.AddViewer(viewer)
	return channel, nil
}

// Idempotent
func addNotificaton(channelLogin string, userId string, set string, must bool, filter model.IFilter) (model.Channel, error) {
	channel, err := addViewerToChannel(channelLogin, userId)
	if err != nil {
		return channel, err
	}

	viewer := channel.GetViewer(userId)
	viewer.Filters = model.FilterSet{Name: set}
	if must {
		viewer.Filters.All = append(viewer.Filters.All, filter)
	} else {
		viewer.Filters.Any = append(viewer.Filters.Any, filter)
	}

	//TODO: Send this as an update to the ORM

	return channel, nil
}
