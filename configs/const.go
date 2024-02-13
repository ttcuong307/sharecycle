package configs

const (
	DnsFormat        string = "%s:%s@tcp(%s:%d)/%s?%s"
	MigrateDnsFormat string = "%s:%s@/%s?%s"
)

const (
	ModeLocal string = "local"
	ModeTest  string = "test"
)
