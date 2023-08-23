package client

import apiLog "github.com/prometheus/common/log"

func InitClient() {
	InitMessaggeClient()
	InitRelationClient()
	InitPingClient()
	InitUserClient()
	InitFeedClient()
	InitPublishClient()
	InitFavoriteClient()
	apiLog.Info("All RPC client initialized")
}
