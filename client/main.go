package main

import (
	"context"
	"fmt"
	"github.com/youmie723/GoMall/kitex_gen/user/userservice"
	user "github.com/youmie723/GoMall/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver("192.168.126.134:8500")
	if err != nil {
		log.Fatal(err)
	}
	u, err := userservice.NewClient("greet.server", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	resp,err:=u.Login(context.TODO(),&user.LoginReq{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v",resp)
}

