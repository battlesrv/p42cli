package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/battlesrv/p42/common"
	"github.com/battlesrv/p42/db"

	"github.com/urfave/cli"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	app := cli.NewApp()
	app.Author = "Konstantin Kruglov"
	app.Email = "kruglovk@gmail.com"
	app.Version = "0.0.1"
	app.HideHelp = true
	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"u"},
			Action:  userManager,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "get",
					Usage: "get information about user from token DB",
				},
				cli.BoolFlag{
					Name:  "set",
					Usage: "get information about user from token DB",
				},
				cli.StringFlag{
					Name:  "pk",
					Usage: "the unique ID of user (=email)",
				},
				cli.BoolFlag{
					Name:  "unlimited",
					Usage: "user use unlimited price",
				},
				cli.StringFlag{
					Name:  "dbhost",
					Usage: "address of DB tokens",
					Value: "127.0.0.1",
				},
				cli.IntFlag{
					Name:  "dbport",
					Usage: "port of DB tokens",
					Value: 3000,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func userManager(c *cli.Context) {
	if err := common.CheckFlags(c, "pk"); err != nil {
		log.Fatalln(err)
	}

	db.NewConn(c.String("dbhost"), c.Int("dbport"))

	if c.Bool("get") {
		var user db.User
		if err := db.Read(c.String("pk"), &user); err != nil {
			log.Fatalln(err)
		}

		fmt.Println(user)
		return
	}

	if c.Bool("set") {
		token := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano()))))
		user := db.User{TokenHash: common.Sha256Sum(token)}

		if c.Bool("unlimited") {
			user.NextAccess = 0
		} else {
			user.NextAccess = uint32(time.Now().Unix())
		}

		if err := db.Write(&user, c.String("pk")); err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("API Token of user %s: %s\n", c.String("pk"), token)
	}
}
