package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// redisdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// creds := &message.EmailCredentials{
	// 	Address:  "test_address",
	// 	Hostname: "test_hostname",
	// 	Name:     "test_name",
	// 	Port:     "test_port",
	// 	Password: "test_password",
	// }

	// data := message.NewHTMLMessage("Test Subject", "Test Body")

	// email := message.Email{
	// 	Credentials: creds,
	// 	Data:        data,
	// }

	// d, err := email.MarshalBinary()
	// if err != nil {
	// 	panic(err)
	// }

	// // err = redisdb.Set("email", d, 0).Err()
	// // if err != nil {
	// // 	panic(err)
	// // }

	// val, err := redisdb.Get("email").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// var _email message.Email
	// err = _email.UnmarshalBinary([]byte(val))
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = redisdb.Del("email").Result()
	// if err != nil {
	// 	panic(err)
	// }
}
