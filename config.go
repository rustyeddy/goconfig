package config

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Configuration struct {
	Addr		string	`json:"addr"`
	Broker		string	`json:"broker"`
	Verbose		bool	`json:"verbose"`
	Filename	string	`json:"-"`			// Do NOT include in confg JSON
}

func (c Configuration) Read(fname string) error {

	// read the json formatted file
	content, err := ioutil.ReadFile(fname)
    if err != nil {
        return fmt.Errorf("Failed to readfile", fname, err)
    }

	err = json.Unmarshal(content, &c)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal ", fname, err)
	}
	return nil
}

func (c Configuration) Write(fname string) error {

	jbuf, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to encode file", fname, err)
	}

	err = ioutil.WriteFile(fname, jbuf, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write file", fname, err)		
	}
	return nil
}

func (c Configuration) Dump() error {

	jbuf, err := json.MarshalIndent(&c, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to encode file", err)
	}
	fmt.Println(string(jbuf))
	return nil
}

func (c Configuration) ServeHTTP(w *http.ResponseWriter, r http.Request) {

	if r.Method != "GET" {
		// return error
	}

	// 

}
