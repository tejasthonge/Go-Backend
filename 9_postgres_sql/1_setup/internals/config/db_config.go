

package config 

type DbConfig struct{
	UserName string  `json:"userName"`
	Password string  `json:"password"` 
	DbName string    `json:"dbName"`
}	