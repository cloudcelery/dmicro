package constant

const (
	TOPIC_USER_LOGIN   string = "topic.user.login"
	TOPIC_USER_CREATED string = "topic.user.created"

	TOPIC_L2A_PREFIX string = "topic.L2A:%d" // %d为appid
	TOPIC_L2G_PREFIX string = "topic.L2G:%s" // %s为gateid
	//TOPIC_L2A_REQ_PREFIX     string = "topic.L2A.req:%d"     // %d为appid
	//TOPIC_L2G_RSP_PREFIX     string = "topic.L2G.rsp:%s"     // %d为gateid
	//TOPIC_L2G_RSP_BIN_PREFIX string = "topic.L2G.rsp.bin:%s" // %d为gateid
	TOPIC_L2G_PUSH_PREFIX string = "topic.L2G.push:%s" // %d为gateid

)

const (
	REDIS_KEY_CONNID = "ums:appid:%d:uid:%d:plat:%d"
)
