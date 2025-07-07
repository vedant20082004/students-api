package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("config","","path to configuration file")
		flag.Parse()

		configPath = *flags

		if configPath==""{
			log.Fatal("config Path not set")
		}

	}

	if _,err :=os.Stat(configPath);os.IsNotExist(err){

		log.Fatalf("config file doesnot exist : %s", configPath )
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg, )

	if err!=nil{
		log.Fatalf("cannot read coonfig file : %s",err.Error())
	}

	return &cfg
}