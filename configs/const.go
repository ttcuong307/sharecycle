package configs

const (
	Mode             string = "MODE"
	ConfigFileFormat string = "config.%s.yml"
	DnsFormat        string = "%s:%s@tcp(%s:%d)/%s?%s"
	MigrateDnsFormat string = "%s:%s@/%s?%s"
	DBDateTimeFormat string = "2006-01-02 15:04:05"
)

const (
	ModeLocal string = "local"
	ModeTest  string = "test"
)
