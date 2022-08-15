package dborm

// 配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex"`
	Value       string
	Module      string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 用户会话

type Session struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	Token     string `gorm:"uniqueIndex"`
	CreatedAt int64
	UpdatedAt int64
}

// CAM 密钥

type Secret struct {
	Id          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"index"`
	SecretId    string `gorm:"uniqueIndex"`
	SecretKey   string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// SSH 密钥

type Sshkey struct {
	Id          uint `gorm:"primaryKey"`
	UserId      uint `gorm:"index"`
	PublicKey   string
	PrivateKey  string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 客户端 - 列表

type SlaveNode struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    string `gorm:"index"`
	HostId    string `gorm:"index"`
	HostName  string
	Address   string
	Status    string
	CreatedAt int64
	UpdatedAt int64
}

// 客户端 - 任务记录

type SlaveTask struct {
	Id        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"index"`
	HostId    string `gorm:"index"`
	HostName  string
	Subject   string
	Content   string
	Status    string
	Result    string
	CreatedAt int64
	UpdatedAt int64
}

// 自动化助手 - 脚本

type TATScript struct {
	Id               uint `gorm:"primaryKey"`
	UserId           uint `gorm:"index"`
	Name             string
	Username         string
	Content          string
	Description      string
	CommandType      string
	WorkingDirectory string
	Timeout          uint
	CreatedAt        int64
	UpdatedAt        int64
}

// 自动化助手 - 历史记录

type TATHistory struct {
	Id                   uint `gorm:"primaryKey"`
	UserId               uint `gorm:"index"`
	KeyId                uint `gorm:"index"`
	Name                 string
	Region               string
	InvocationId         string
	InvocationStatus     string
	InvocationResultJson string
}

// 用户

type User struct {
	Id          uint   `gorm:"primaryKey"`
	Username    string `gorm:"uniqueIndex"`
	Password    string
	AppToken    string `gorm:"uniqueIndex"`
	Description string `gorm:"default:什么也没有"`
	Secrets     []Secret
	Sessions    []Session
	CreatedAt   int64
	UpdatedAt   int64
}
