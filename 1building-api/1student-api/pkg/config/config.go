package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"adress" env-required:"true"`
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"` // if we not set defoutlt value then alos done but we mut have the require value for it
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// we setting these function name as must load that means we are loading all the conig files in this
func MustLoad() *Config {

	//here we trying to get the config patha form the envarmetal varible if any one can set it
	//like export CONFIG_PATH=<PATH>
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		//if not getted the path from the envarmetal varible then

		//we geting form the flag passed at time of runing the our porject
		//here we geting
		//this can pass as
		// go run main.go --config filepath
		flags := flag.String("config", "", "Path to the configration")
		flag.Parse()

		// now we assing new config paht value to configPath vriable
		configPath = *flags

		//then we chencking ,
		//then also not getting path

		if configPath == "" {
			log.Fatal("Not geting config path form enevarmetal varible as well as form flag\nplease set it or pass at time of runing server")
		}
	}

	//now chenkig the that path any file is present or not
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("File not exit at path :%s", configPath)
	}

	var cfg Config

	//now we have to assing thee all config file to the Config struct by using cleanenv package metheod

	err := cleanenv.ReadConfig(configPath, &cfg) //here we have to pass fist parametor as confif patha secor paremtor were is to write the config baisclly in whic struct we have to pass the address of it

	if err != nil {
		log.Fatalf("Can not Read the the file %s", err.Error())
	}

	return &cfg

}
