# Golang Tools Config

Golang Tools Config Support：

- Json
- Yaml
- Toml
- Xml

## Install

```
go get github.com/dingdayu/deawood/config
```


## Example

### json

```
	
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.json", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

Run:

```
go run ../example/config/main.go -c ../example/config/conf.json
{8080 [{gitbook /gitbook [ll ls] dingdayu}]}
```

### xml


```

	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.xml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```



### toml


```
	
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.toml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```


### yaml


```
	
	type Instance struct {
		Name 	string
		Path	string
		Cmd		[]string
		User	string
	}

	
	type Config struct {
		Port int
		Instance []Instance
	}

	var conf Config
	err := config.New("conf.yaml", &conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
```

