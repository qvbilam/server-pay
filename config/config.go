package config

type ServerConfig struct {
	Name                 string               `mapstructure:"name" json:"name"`
	Port                 int                  `mapstructure:"port" json:"port"`
	Tags                 []string             `mapstructure:"tags" json:"tags"`
	DBConfig             DBConfig             `mapstructure:"db" json:"db"`
	RedisConfig          RedisConfig          `mapstructure:"redis" json:"redis"`
	UserServerConfig     UserServerConfig     `mapstructure:"user-server" json:"user-server"`
	ContactServerConfig  ContactServerConfig  `mapstructure:"contact-server" json:"contact-server"`
	RabbitMQServerConfig RabbitMQServerConfig `mapstructure:"rabbit-server" json:"rabbit-server"`
}

type DBConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database int    `mapstructure:"database" json:"database"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type ContactServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type RabbitMQServerConfig struct {
	Host                string `mapstructure:"host" json:"host"`
	Port                int64  `mapstructure:"port" json:"port"`
	Name                string `mapstructure:"name" json:"name"`
	User                string `mapstructure:"user" json:"user"`
	Password            string `mapstructure:"password" json:"password"`
	Exchange            string `mapstructure:"exchange" json:"exchange"`
	ExchangeChatPrivate string `mapstructure:"exchange_chat_private" json:"exchange_chat_private"`
	ExchangeChatRoom    string `mapstructure:"exchange_chat_room" json:"exchange_chat_room"`
	ExchangeChatGroup   string `mapstructure:"exchange_chat_group" json:"exchange_chat_group"`
	QueuePrefix         string `mapstructure:"queue_suffix" json:"queue_suffix"`
}
