package config

type Config struct {
	Kafka Kafka
	Db    Db
	Api   Api
}

type Kafka struct {
	BrokerUrl       string
	ConsumerGroupId string
	Topic           string
}

type Db struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type Api struct {
	Port string
	Jwt  string
}
