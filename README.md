# PKG mailing - chesgomailing

### INSTALL
```shell
go get github.com/Leonardo-Antonio/chesgomailings
```

### USE

[Example](main.go)

```go
package main

import "github.com/Leonardo-Antonio/chesgomailings/chesgomailing"

func main() {
	config := chesgomailing.Config{
        From: "example@test.com", Password: "abcdefgbigk", 
        Host: "smtp.gmail.com", Port: "587",
    }
	chesgomailing.Send(&config, chesgomailing.DataToSend{
		Format:  chesgomailing.TYPE_HTML,
		Subject: "HOLA",
		Body:    "<h1>HOLA</h1>",
		To:      []string{"test01@test.com", "test02.nl8@test.com"},
	})
}

```