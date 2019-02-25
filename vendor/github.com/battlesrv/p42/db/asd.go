package db

import (
	"log"
	"time"

	as "github.com/aerospike/aerospike-client-go"
)

var (
	// ASClient ..
	ASClient       *as.Client
	nameSpace      = "p42"
	expirationTime = uint32(604800)  // in seconds, 7 days
	writeTimeout   = time.Second * 2 // 2 seconds
)

// User ..
type User struct {
	TokenHash string `as:"token_hash"`
	// TokenHash  string `as:"token_hash"`
	NextAccess uint32 `as:"next_access"`
}

// NewConn ..
func NewConn(host string, port int) {
	var err error
	ASClient, err = as.NewClient(host, port)
	if err != nil {
		log.Fatalln(err)
	}
}

// Read ..
func Read(pk string, out *User) error {
	key, err := as.NewKey(nameSpace, "users", pk)
	if err != nil {
		return err
	}

	if err = ASClient.GetObject(nil, key, out); err != nil {
		return err
	}
	//
	return nil
}

// Write ..
func Write(user *User, pk string) error {
	key, err := as.NewKey(nameSpace, "users", pk)
	if err != nil {
		return err
	}

	newPolicy := as.NewWritePolicy(0, 0)
	newPolicy.Expiration = expirationTime
	newPolicy.Timeout = writeTimeout

	if err = ASClient.PutObject(newPolicy, key, user); err != nil {
		return err
	}
	//
	return nil
}
