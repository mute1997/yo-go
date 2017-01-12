package main

import (
    "net/url"
)

type User struct {
    Api_token string
}

//get user object
func (user User) me() error {
    uri := "https://api.justyo.co/me/"

    values := url.Values{}
    values.Set("api_token",user.Api_token)

    return get(uri,values)
}

//send yo
func (user User) yo(username string) error{
    uri := "https://api.justyo.co/yo/"

    values := url.Values{}
    values.Set("username",username)
    values.Set("api_token",user.Api_token)

    return post(uri,values)
}

//send yo all
func (user User) yoall() error {
    uri := "https://api.justyo.co/yoall/"

    values := url.Values{}
    values.Set("api_token",user.Api_token)

    return post(uri,values)
}
