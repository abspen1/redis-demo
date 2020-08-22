package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

/*
redis.Bool() – converts a single reply to a bool
redis.Bytes() – converts a single reply to a byte slice ([]byte)
redis.Float64() – converts a single reply to a float64
redis.Int() – converts a single reply to a int
redis.String() – converts a single reply to a string
redis.Values() – converts an array reply to an slice of individual replies
redis.Strings() – converts an array reply to an slice of strings ([]string)
redis.ByteSlices() – converts an array reply to an slice of byte slices ([][]byte)
redis.StringMap() – converts an array of strings (alternating key, value) into a map[string]string. Useful for HGETALL etc
*/

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

//Project struct
type Project struct {
	Language    string
	Description string
}

func main() {
	secret := goDotEnvVariable("REDIS")

	client, err := redis.Dial("tcp", "10.10.10.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do("AUTH", secret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!", response)
	worked := "It didn't work"
	client.Do("SET", "foodoo", 1)
	exists, _ := redis.Bool(client.Do("EXISTS", "foodoo"))
	if exists {
		worked = "It worked!"
	}
	fmt.Println(worked)
	client.Do("DEL", "projects")

	proj1 := Project{"PYTHON", "Description 1"}
	proj2 := Project{"JAVA", "Description 2"}
	proj3 := Project{"PYTHON", "Description 3"}

	proj1En, _ := json.Marshal(proj1)
	proj2En, _ := json.Marshal(proj2)
	proj3En, _ := json.Marshal(proj3)

	client.Do("RPUSH", "projects", proj1En)
	client.Do("RPUSH", "projects", proj2En)
	client.Do("RPUSH", "projects", proj3En)

	var unencoded *Project

	project1, _ := redis.Strings(client.Do("LRANGE", "projects", 0, -1))

	len, _ := redis.Int(client.Do("LLEN", "projects"))

	i := 0

	var s []string

	for i < len {
		json.Unmarshal([]byte(project1[i]), &unencoded)
		fmt.Println(unencoded.Language)
		fmt.Println(unencoded.Description)
		i++
		s = append(s, unencoded.Language)
		s = append(s, unencoded.Description)
	}
	fmt.Println(s)
}
