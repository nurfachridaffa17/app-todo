package env

import (
	"flag"
	"log"
	"os"
	"strconv"

	"app-todo/pkg/file"

	"github.com/joho/godotenv"
)

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (*env) Load() {
	var err error
	mode := flag.String("mode", "dev", "dev, prod or stage")
	flag.Parse()
	rootPath := file.GetRootDirectory()

	switch *mode {
	case "dev":
		err = godotenv.Load(rootPath + "/.env.development")
	case "prod":
		err = godotenv.Load(rootPath + "/.env.production")
	case "stage":
		err = godotenv.Load(rootPath + "/.env.staging")
	default:
		err = godotenv.Load(rootPath + "/.env.development")
	}

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func (e *env) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetBool(name string) bool {
	s := e.GetString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (e *env) GetFloat(name string) float64 {
	s := e.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}
