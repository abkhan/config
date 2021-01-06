# Config package
This package provides functions to read a YAML file into a struct.
Any missing fields would be identified.

## Example:

```
	flag.Parse()

	c := sconf{}
	conf.Load(&c)

	if err := config.ValidateConf(c); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Config: %+v\n", c)
```

Parse the flag first to tell the package the config file name.
Load the file, and then validate.

