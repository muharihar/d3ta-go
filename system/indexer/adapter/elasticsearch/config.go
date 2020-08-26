package elasticsearch

import (
	"fmt"

	es6 "github.com/elastic/go-elasticsearch/v6"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

// ConfigParserES6 config parser for elastic search 6
func ConfigParserES6(conf interface{}) es6.Config {
	cfg := make(map[string]interface{})
	if fmt.Sprintf("%T", conf) != "map[string]interface {}" {
		cfg["addresses"] = []string{"http://localhost:9200"}
		cfg["username"] = ""
		cfg["password"] = ""
		cfg["cloudID"] = ""
		cfg["APIKey"] = ""
	} else {
		conf := conf.(map[string]interface{})
		cfg["addresses"] = conf["addresses"]
		cfg["username"] = conf["username"]
		cfg["password"] = conf["password"]
		cfg["cloudID"] = conf["cloudid"]
		cfg["APIKey"] = conf["apikey"]
	}
	var addresses []string
	for _, v := range cfg["addresses"].([]interface{}) {
		addresses = append(addresses, v.(string))
	}
	opt := es6.Config{
		Addresses: addresses,
		Username:  fmt.Sprintf("%s", cfg["username"]),
		Password:  fmt.Sprintf("%s", cfg["password"]),
		CloudID:   fmt.Sprintf("%s", cfg["cloudID"]),
		APIKey:    fmt.Sprintf("%s", cfg["APIKey"]),
	}
	return opt
}

// ConfigParserES7 config parser for elastic search 7
func ConfigParserES7(conf interface{}) es7.Config {
	cfg := make(map[string]interface{})
	if fmt.Sprintf("%T", conf) != "map[string]interface {}" {
		cfg["addresses"] = []string{"http://localhost:9200"}
		cfg["username"] = ""
		cfg["password"] = ""
		cfg["cloudID"] = ""
		cfg["APIKey"] = ""
	} else {
		conf := conf.(map[string]interface{})
		cfg["addresses"] = conf["addresses"]
		cfg["username"] = conf["username"]
		cfg["password"] = conf["password"]
		cfg["cloudID"] = conf["cloudid"]
		cfg["APIKey"] = conf["apikey"]
	}
	var addresses []string
	for _, v := range cfg["addresses"].([]interface{}) {
		addresses = append(addresses, v.(string))
	}
	opt := es7.Config{
		Addresses: addresses,
		Username:  fmt.Sprintf("%s", cfg["username"]),
		Password:  fmt.Sprintf("%s", cfg["password"]),
		CloudID:   fmt.Sprintf("%s", cfg["cloudID"]),
		APIKey:    fmt.Sprintf("%s", cfg["APIKey"]),
	}
	return opt
}

// ConfigParserES8 config parser for elastic search 8
func ConfigParserES8(conf interface{}) es8.Config {
	cfg := make(map[string]interface{})
	if fmt.Sprintf("%T", conf) != "map[string]interface {}" {
		cfg["addresses"] = []string{"http://localhost:9200"}
		cfg["username"] = ""
		cfg["password"] = ""
		cfg["cloudID"] = ""
		cfg["APIKey"] = ""
	} else {
		conf := conf.(map[string]interface{})
		cfg["addresses"] = conf["addresses"]
		cfg["username"] = conf["username"]
		cfg["password"] = conf["password"]
		cfg["cloudID"] = conf["cloudid"]
		cfg["APIKey"] = conf["apikey"]
	}
	var addresses []string
	for _, v := range cfg["addresses"].([]interface{}) {
		addresses = append(addresses, v.(string))
	}
	opt := es8.Config{
		Addresses: addresses,
		Username:  fmt.Sprintf("%s", cfg["username"]),
		Password:  fmt.Sprintf("%s", cfg["password"]),
		CloudID:   fmt.Sprintf("%s", cfg["cloudID"]),
		APIKey:    fmt.Sprintf("%s", cfg["APIKey"]),
	}
	return opt
}
