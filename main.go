package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func genXid() { //1
	//id.String()
	fmt.Printf("github.com/rs/xid:              %s\n", xid.New().String())
}

func genBetterGUID() { //2
	//betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:      %s\n", betterguid.New())
}

func genSid() { //3
	fmt.Printf("github.com/chilts/sid:          %s\n", sid.Id())
	//sid.Id()
}

func genUUIDv4() { //4
	id := uuid.NewV4()
	//id.String()
	fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func genUUID() { //5
	//id.String()
	fmt.Printf("github.com/google/uuid:         %s\n", guuid.New().String())
}

func genKsuid() { //6
	//id.String()
	fmt.Printf("github.com/segmentio/ksuid:     %s\n", ksuid.New().String())
}

func genUlid() { //7
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
	//id.String()
}

func genShortUUID() { //8
	//shortuuid.New()
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", shortuuid.New())
}

func genSonyflake() { //9
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
}

func main() {

	genXid()        //1
	genBetterGUID() //2
	genSid()        //3
	genUUIDv4()     //4
	genUUID()       //5
	genKsuid()      //6
	genUlid()       //7
	genShortUUID()  //8
	genSonyflake()  //9

	fmt.Println()
	guid := xid.New()
	fmt.Println(guid.Machine())
	fmt.Println(guid.Pid())
	fmt.Println(guid.Time())
	fmt.Println(guid.Counter())
	fmt.Println("xid0:", guid.String())
	fmt.Println("xid1:", xid.New().String())
	fmt.Println("xid2:", xid.New().String())

}
