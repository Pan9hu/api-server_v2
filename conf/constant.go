package conf

const (
	ConfigEnv                      = "MELO_CONFIG"
	DefaultWindowsServerConfigFile = "\\AppData\\Local\\melo-cmdb\\api.properties" // windows系统
	DefaultUnixConfigFile          = "/etc/melo-cmdb/api.properties"               // unix系统
	DefaultOsxConfigFile           = "/etc/melo-cmdb/api.properties"               // darwin系统
)
