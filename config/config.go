package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)
type Config struct{
	Redis RedisConfig `yaml:"redis"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type RedisConfig struct{
	Addr  string `yaml:"addr"`
	Password string `yaml:"password"`
	DB int `yaml:"db"`
}

type PostgresConfig struct {  
    Host     string `yaml:"host"`  
    Port     int    `yaml:"port"`  
    User     string `yaml:"user"`  
    Password string `yaml:"password"`  
    DBName   string `yaml:"dbname"`  
}  	

func GetConfigs() Config{  
    var configs Config  
    dir, _ := os.Getwd()  
    yamlFile, err := os.ReadFile(dir + "/config/config-development.yaml")  
    if err != nil {  
        log.Fatal("yamlFile.Get err:", err)  
    }  
    if yaml.Unmarshal(yamlFile, &configs); err != nil {  
        log.Fatal("Error unmarshalling YAML:", err)  
    }
    return configs
}

func GetRedisConfigs() RedisConfig{
	return GetConfigs().Redis	
}

func GetPostgresConfigs() PostgresConfig{
	return GetConfigs().Postgres	
}