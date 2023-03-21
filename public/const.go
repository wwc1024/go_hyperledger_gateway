package public

const (
	ValidatorKey        = "ValidatorKey"
	TranslatorKey       = "TranslatorKey"
	AdminSessionInfoKey = "AdminSessionInfoKey"

	LoadType1 = "传感器类"
	LoadType2 = "控制类"
	LoadType3 = "其他类"

	HTTPRuleTypePrefixURL = 0
	HTTPRuleTypeDomain    = 1

	// RedisFlowDayKey  = "all_service_count"
	// RedisFlowHourKey = "flow_hour_count"

	FlowTotal         = "flow_total"
	FlowServicePrefix = "flow_service_"
	FlowPeerPrefix    = "flow_peer_"

	JwtSignKey = "my_sign_key"
	JwtExpires = 60 * 60
)

var (
	LoadTypeMap = map[string]string{
		LoadType1: "传感器类",
		LoadType2: "控制类",
		LoadType3: "其他类",
	}
	NowSample = int64(1)
	DhtId     = int64(2)
	DhtIdInt  = int64(2)
	SdkConfig = ".config2.yaml"
)
