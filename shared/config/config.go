package config

type Config struct {
	App    App
	Server Server
	DB     DB
	Jwt    Jwt
	Redis  Redis
	Front  Front
}

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type Jwt struct {
	Secret          string
	JwtValidityTime int
}

type Redis struct {
	Host                      string
	Port                      string
	Username                  string
	Password                  string
	TcpNameSpace              string
	TcpJobQueue               string
	TcpJobQueueProgressionSub string
}

type Front struct {
	Host string
	Port string
}
