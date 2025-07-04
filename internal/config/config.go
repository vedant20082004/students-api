package config

type HttpServer struct{
	Addr string
}

type Config struct{

	Env string `yaml:"env" env:"ENV" env-required:"true"` 
	StoragePath string `yaml:"storage_path`
	HttpServer 
}