package main

type Config struct {
	RainService string
	RefreshTime int
}

var config = &Config{
	RainService: "http://127.0.0.1:7362/json",
	RefreshTime: 500,
}
