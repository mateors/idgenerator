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

func genShortUUID() {

	shortuuid.New()
	//id := shortuuid.New()
	//fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
}

func genUUID() {
	id := guuid.New()
	id.String()
	//fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func genXid() {
	id := xid.New()
	id.String()
	//fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	id.String()
	//fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
}

func genBetterGUID() {
	betterguid.New()
	//id := betterguid.New()
	//fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	id.String()
	//fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	_, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	//fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
}

func genSid() {
	//id := sid.Id()
	//fmt.Printf("github.com/chilts/sid:          %s\n", id)
	sid.Id()
}

func genUUIDv4() {
	id := uuid.NewV4()
	id.String()
	//fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func main() {

	genXid()   //
	genKsuid() //
	genBetterGUID()
	genUlid()      //
	genSonyflake() //
	genSid()       //
	genShortUUID() //
	genUUIDv4()    //
	genUUID()      //ok

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
