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

//User struct
type User struct {
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

	// Demo the strings
	client.Do("SET", "language", "Golang")
	language, err := redis.String(client.Do("GET", "language"))
	if checkError(err) {
		//Print our key if it exists
		fmt.Println("language: " + language)
	}

	// Demo the sets
	client.Do("SADD", "golangList", "value1", "value2", "value3", "value4")
	client.Do("SADD", "PowerList", "value1", "value5", "value6", "value7")

	intercept, err := redis.Strings(client.Do("SINTER", "golangList", "PowerList"))
	if checkError(err) {
		fmt.Println("intercept: ", intercept)
	}

	union, err := redis.Strings(client.Do("SUNION", "PowerList", "golangList"))
	if checkError(err) {
		fmt.Println("union: ", union)
	}

	members, err := redis.Strings(client.Do("SMEMBERS", "golangList"))
	if checkError(err) {
		fmt.Println("members: ", members)
	}

	scard, err := redis.Int64(client.Do("SCARD", "golangList"))
	if checkError(err) {
		fmt.Println("scard: ", scard)
	}

	client.Do("HSET", "Hero", "Name", "Draw Ranger")
	client.Do("HSET", "Hero", "Health", "600")
	client.Do("HSET", "Hero", "Mana", "200")
	// client.Do("DEL", "PYTHON")

	strMap, err := redis.StringMap(client.Do("HGETALL", "Hero"))
	if checkError(err) {
		fmt.Println("hash: ", strMap)
	}

	// client.Do("DEL", "projects")

	var unencoded *User

	rangL, _ := redis.Strings(client.Do("LRANGE", "Languages", "0", "-1"))

	json.Unmarshal([]byte(rangL[0]), &unencoded)

	// rangD, err := redis.ByteSlices(client.Do("LRANGE", "Descriptions", "0", "-1"))
	// if checkError(err) {
	// 	fmt.Println()
	// 	fmt.Println("rangeL: ", rangD)
	// }

	strMap, err = redis.StringMap(client.Do("HGETALL", "daily_status"))
	if checkError(err) {
		fmt.Println()
		str, _ := json.Marshal(strMap)
		fmt.Println("strmap: ", string(str))
	}

	exist := ""

	exist, _ = redis.String(client.Do("HGET", "Hero", "Name"))
	if exist == "" {
		fmt.Println("HGET doesn't  work with hashes")
	}
	exist = ""
	exist, _ = redis.String(client.Do("HGET", "Hero", "FooDoO"))
	if exist == "" {
		fmt.Println(fmt.Sprintf("HGET works with %s  hashes", language))
	}
	exist, _ = redis.String(client.Do("HGET", "emails", "rgulden7@gmail.com"))
	fmt.Println("Email: ", exist)

	defer client.Close()
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println("There was an error! ", err)
		return false
	}
	return true
}
