package main

import (
    "errors"
    "os"
    "io/ioutil"
)

type Command struct{
    Args []string
    user User
    cmds map[string]func(user User) error
    help_message string
}

func NewCommand(args []string) *Command {
    p := new(Command)
    p.Args = args 
    p.cmds = map[string]func(user User) error{
                                            "--yo":send_yo,
                                            "--yoall":send_yoall,
                                            "--me":get_me,
                                        }
    p.help_message = `help:
    yo init "api_key"   : Write "api_key" .yo file
    yo --yo "username"  : Send yo "username"
    yo --yoall          : Send yo all
    yo --me             : Get me object`
    return p
}

func (command Command) exec() error {
    if len(command.Args) < 2 {
        return errors.New(command.help_message)
    }

    if command.Args[1] == "init" {
        if len(os.Args) == 3 {
            ioutil.WriteFile(".yo",[]byte(os.Args[2]),os.ModePerm)
            return nil
        } else {
            return errors.New(command.help_message)
        }
    }else if function, ok := command.cmds[command.Args[1]]; ok {
        if isExist(".yo") {
            file, err := ioutil.ReadFile(".yo")

            command.user = User{Api_token: string(file)}

            if err != nil {
                return errors.New(".yo file does not exist")
            }
        } else {
            return errors.New(".yo file does not exist")
        }

        err := function(command.user)
        if err != nil {
            return errors.New(command.help_message)
        }
    } else {
        return errors.New(command.help_message)
    }
    return nil
}

func send_yo(user User) error{
    if len(os.Args) == 3 {
        user.yo(os.Args[2])
        return nil
    } else {
        return errors.New("")
    }
}

func send_yoall(user User) error{
    return user.yoall()
}

func get_me(user User) error{
    return user.me()
}

func isExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}
