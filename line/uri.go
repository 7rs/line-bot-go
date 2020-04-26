package line

const (
	APIHost     = "https://api.line.me/v2"
	APIDataHost = "https://api-data.line.me/v2"
)

// host: https://api.line.me/v2
const (
	// SECTION: Message
	// Send reply message.
	Reply = "/bot/message/reply"

	// Send push message.
	Push = "/bot/message/push"

	// Send multicast message.
	Multicast = "/bot/message/multicast"

	// Send narrowcast message.
	Narrowcast = "/bot/message/narrowcast"

	// Get narrowcast message status.
	NarrowcastProgress = "/bot/message/progress/narrowcast"

	// Send broadcast message.
	Broadcast = "/bot/message/broadcast"

	// Get the target limit for additional messages.
	Quota = "/bot/message/quota"

	// Get number of messages sent this month.
	QuotaConsumption = "/bot/message/quota/consumption"

	// Get number of sent reply messages.
	// query: ?date={yyyyMMdd}
	DeliveryReply = "/bot/message/delivery/reply"

	// Get number of sent push messages.
	// query: ?date={yyyyMMdd}
	DeliveryPush = "/bot/message/delivery/push"

	// Get number of sent multicast messages.
	// query: ?date={yyyyMMdd}
	DeliveryMulticast = "/bot/message/delivery/multicast"

	// Get number of sent broadcast messages.
	// query: ?date={yyyyMMdd}
	DeliveryBroadcast = "/bot/message/delivery/broadcast"

	// Get number of message deliveries.
	// query: ?date=20190418
	InsightMessageDelivery = "/bot/insight/message/delivery"

	// Get number of followers.
	// query: ?date=20190418
	InsightFollowers = "/bot/insight/followers"

	// Get friend demographics.
	InsightDemographic = "/bot/insight/demographic"

	// Get user interaction statistics.
	// query: ?requestId=f70dd685-499a-4231-a441-f24b8d4fba21
	// NOTE: This API is temporarily suspended.
	InsightMessageEvent = "/bot/insight/message/event"

	// SECTION: Profile
	// Gets user profile
	// path: /{userId}
	Profile = "/bot/profile"

	// SECTION: Group
	// Gets group member ids.
	// path: /{groupId}/members/ids
	// query: ?start={continuationToken}
	//
	// Gets group member profiles.
	// path: /{groupId}/member/{userId}
	//
	// Leaves group.
	// path: /{groupId}/leave
	Group = "/bot/group"

	// SECTION: Room
	// Gets room member ids.
	// path: {roomId}/members/ids
	// query: ?start={continuationToken}
	//
	// Gets room member profiles.
	// path: /{roomId}/member/{userId}
	//
	// Leaves room.
	// path: /{roomId}/leave
	Room = "/bot/room"
)

// host: https://api-data.line.me/v2
// Send broadcast message.
// path: /{messageId}/content
const Content = "https://api-data.line.me/v2/bot/message"
