package main

import (
	"testing"
)

// func TestGenXid(t *testing.T) {

// 	genXid()
// 	t.Fail()

// }

func BenchmarkGenXid(b *testing.B) {

	//b.Run("github.com/rs/xid", func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genXid()
	}
	//})
}

func BenchmarkGenBetterGUID(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genBetterGUID()
	}
}

func BenchmarkGenSid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genSid()
	}
}

func BenchmarkGenUUID(b *testing.B) {

	//b.Run("github.com/google/uuid", func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genUUID()
	}
	//})

}

func BenchmarkGenUUIDv4(b *testing.B) {

	//b.Run("github.com/satori/go.uuid", func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genUUIDv4()
	}
	//})
}

func BenchmarkGenKsuid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genKsuid()
	}
}

func BenchmarkGenUlid(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genUlid()
	}
}

func BenchmarkGenShortUUID(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genShortUUID()
	}
}

func BenchmarkGenSonyflake(b *testing.B) {

	for i := 0; i < b.N; i++ {
		genSonyflake()
	}
}
