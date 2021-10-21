package models

type Todo struct{
	ID int 			`json:"id"`
	Message string 	`json:"message"`
	Status string 	`json:"status"`
}

// Config 后面mapstructure是放前面的标题
type Config struct{
	*Mysqlconf 		`mapstructure:"mysql"`
	*Logconf   		`mapstructure:"log"`
}

type Mysqlconf struct{
	Host string 	`mapstructure:"host"`
	Port int 		`mapstructure:"port"`
	User string 	`mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Logconf struct{
	Level string 	`mapstructure:"level"`
	Maxsize int		`mapstructure:"max_size"`
	Maxage int 		`mapstructure:"max_age"`
	Maxbackups int 	`mapstructure:"max_backups"`
}

