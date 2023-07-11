## envconfig

Load values from environment variables

	cfg := envconfig.New()

	cfg.LoadInt("HTTP_PORT", 8080)
	cfg.LoadString("HTTP_HOST", "localhost")
	cfg.LoadString("HOME", "localhost")

	fmt.Println(cfg.Ints["HTTP_PORT"])
	fmt.Println(cfg.Strings["HTTP_HOST"])
	fmt.Println(cfg.Strings["HOME"])

To get

    2003/07/11 05:42:32 <HTTP_PORT> is not set, using default value <8080>
    2003/07/11 05:42:32 <HTTP_HOST> is not set, using default value <localhost>
    8080
    localhost
    /home/user

Change the message shown when a default value is used

	cfg.UsingDefaultValueFunc(func(name string, v interface{}) {
		log.Printf("using default value for %s: %v", name, v)
	})

	cfg.LoadInt("HTTP_PORT", 8080)

    2023/07/11 15:42:32 using default value for HTTP_PORT: 8080

Features

- load configuration values from environment variables
- support for int and string
- assign a default value
- show a message when a variable uses the default value. You can overload the function that shows that message
