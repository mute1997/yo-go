package main

import (
    "net/http"
    "net/url"
    "fmt"
    "io/ioutil"
)

func get(uri string,values url.Values) error{
    resp, err := http.Get(uri + "?" + values.Encode())

    if err != nil {
        return err
    }

    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byteArray))

    defer resp.Body.Close()

    return err
}

func post(uri string,values url.Values) error{
    resp, err := http.PostForm(uri, values)

    if err != nil {
        return err
    }

    byteArray, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(byteArray))

    defer resp.Body.Close()

    return err
}
