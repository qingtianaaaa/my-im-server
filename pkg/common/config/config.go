package config

const (
	ETCD      = "etcd"
	KUBERNETS = "kubernets"
)

var (
	FileName                 = "config.yml"
	DiscoveryConfigFileName  = "discovery.yml"
	LogConfigName            = "log.yml"
	OpenIMRPCAuthCfgFileName = "rpc-auth.yml"
)

type API struct {
	Api struct {
		ListenIP         string `mapstructure:"listenIP"`
		Port             int    `mapstructure:"port"`
		CompressionLevel string `mapstructure:"compressionLevel"`
	} `mapstructure:"api"`
	Prometheus struct {
		Enable       bool   `mapstructure:"enable"`
		AutoSetPorts bool   `mapstructure:"autoSetPorts"`
		Ports        []int  `mapstructure:"ports"`
		GrafanaURL   string `mapstructure:"grafanaURL"`
	} `mapstructure:"prometheus"`
}

type Discovery struct {
	Enable     string     `mapstructure:"enable"`
	Etcd       Etcd       `mapstructure:"etcd"`
	Kubernets  Kubernets  `mapstructure:"kubernets"`
	RpcService RpcService `mapstructure:"rpcService"`
}

type Kubernets struct {
	Namespace string `mapstructure:"namespace"`
}

type Etcd struct {
	RootDirectory string   `mapstructure:"rootDirectory"`
	Address       []string `mapstructure:"address"`
	Username      string   `mapstructure:"username"`
	Password      string   `mapstructure:"password"`
}

type RpcService struct {
	User           string `mapstructure:"user"`
	Friend         string `mapstructure:"friend"`
	Msg            string `mapstructure:"msg"`
	Push           string `mapstructure:"push"`
	MessageGateway string `mapstructure:"messageGateway"`
	Group          string `mapstructure:"group"`
	Auth           string `mapstructure:"auth"`
	Conversation   string `mapstructure:"conversation"`
	Third          string `mapstructure:"third"`
}

type Auth struct {
	RPC struct {
		RegisterIP   string `mapstructure:"registerIP"`
		ListenIP     string `mapstructure:"listenIP"`
		AutoSetPorts bool   `mapstructure:"autoSetPorts"`
		Ports        []int  `mapstructure:"ports"`
	} `mapstructure:"rpc"`
	Prometheus  Prometheus `mapstructure:"prometheus"`
	TokenPolicy struct {
		Expire int64 `mapstructure:"expire"`
	} `mapstructure:"tokenPolicy"`
}

type Log struct {
	StoragePath string `mapstructure:"storagePath"`
	IsStdout    bool   `mapstructure:"isStdout"`
	WithStack   bool   `mapstructure:"withStack"`
}

type Prometheus struct {
	Enable bool  `mapstructure:"enable"`
	Ports  []int `mapstructure:"ports"`
}

type AllConfig struct {
	Discovery Discovery
	Log       Log
	Auth      Auth
	API       API
}
