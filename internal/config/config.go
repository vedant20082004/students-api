package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string
}

// env-default:"production"
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("Config","","path to configuration file")
		flag.Parse()

		configPath = *flags

		if configPath==""{
			log.Fatal("Config Path not set")
		}

	}

	if _,err :=os.Stat(configPath);os.IsNotExist(err){

		log.Fatalf("Config file doesnot exist : %s", configPath )
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg, )

	if err!=nil{
		log.Fatalf("cannot read coonfig file : %s",err.Error())
	}

	return &cfg
}