package main

import (
        "io/ioutil"
        "log"
        "net/http"
        "regexp"
        "time"
)

func ping(w http.ResponseWriter, r *http.Request) {
        ping := []byte("ping")
        w.Write(ping)
}

func blaze(w http.ResponseWriter, r *http.Request) {
        tests, ok := r.URL.Query()["test"]

        if !ok || len(tests[0]) < 1 {
                log.Println("Url Param 'test' is missing")
                return
        }

        // Query()["test"] will return an array of items,
        // we only want the single item.
        key := tests[0]
        urlmap := map[string]string{
                "yacbacktest1": "http://yacback.redf4rth.net:8077/static/img/student.html",
        }
        expectedmap := map[string]string{
                "yacback1": "<title>Student</title>",
        }

        matchb := []byte("unknown")

        _, found := urlmap[key]
        if !found {
        } else {

                timeout := time.Duration(5 * time.Second)
                client := http.Client{
                        Timeout: timeout,
                }
                res, err := client.Get(urlmap[key])

                if err != nil {
                        matchb = []byte("timeout")
                } else {
                        robots, err := ioutil.ReadAll(res.Body)
                        res.Body.Close()
                        if err != nil {
                                log.Fatal(err)
                        }

                        match, _ := regexp.MatchString(expectedmap[key], string(robots))

                        if match {
                                matchb = []byte("yes")
                        }
                }
        }

        w.Header().Set("Content-Type", "application/json")
        // Access-Control-Allow-Origin yes is necessary.
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Write(matchb)
}
