package main

const CONFIG = "./../etc/app.yaml"

func main()  {
	config := new(Config)
	config.LoadConfig(CONFIG)

	app := new(App)
	app.Bootstrap(config).Process();
}
