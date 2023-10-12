package email

type Email struct {
	Enabled       bool     `env:"EMAIL_ENABLED" long:"enabled" description:"enables email support" yaml:"enabled"`
	Hostname      string   `env:"EMAIL_HOSTNAME" long:"hostname" description:"hostname of mail server" yaml:"hostname"`
	Port          int      `env:"EMAIL_PORT" long:"port" description:"port of mail server" yaml:"port"`
	Username      string   `env:"EMAIL_USERNAME" long:"username" description:"username to authenticate to mail server" yaml:"username"`
	Password      string   `env:"EMAIL_PASSWORD" long:"password" description:"password to authenticate to mail server" yaml:"password"`
	FromAddr      string   `env:"EMAIL_FROM_ADDR" long:"from-addr" description:"address to use as 'From'" yaml:"from_addr"`
	SendAddrs     []string `env:"EMAIL_SEND_ADDRS" long:"send-addrs" description:"addresses to send notifications to" yaml:"send_addrs"`
	TLSSkipVerify bool     `env:"EMAIL_TLS_SKIP_VERIFY" long:"tls-skip-verify" description:"skip SMTP TLS certificate validation" yaml:"tls_skip_verify"`
	MandatoryTLS  bool     `env:"EMAIL_MANDATORY_TLS" long:"mandatory-tls" description:"require TLS for SMTP connections. Defaults to opportunistic." yaml:"mandatory_tls"`
} `group:"Email Options" namespace:"email" yaml:"email"`
