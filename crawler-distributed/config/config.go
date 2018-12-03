package config

const (
	// Parser names
	ParseCity = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile = "ParseProfile"
	NilParser = "NilParser"

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRPC = "ItemSaveService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate limiting
	Qps = 20
)