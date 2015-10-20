package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

import (
	"github.com/couchbase/gocb"
)

// bucket reference - reuse as bucket reference in the application
var bucket *gocb.Bucket

// uuid generate unique couchbase document key
func uuid() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return string(out), nil
}

func main() {
	// Connect to Cluster
	cluster, err := gocb.Connect("couchbase://127.0.0.1")
	if err != nil {
		fmt.Println("ERRROR CONNECTING TO CLUSTER")
		panic(err)
	}

	// Open Bucket
	bucket, err = cluster.OpenBucket("default", "")
	if err != nil {
		fmt.Println("ERRROR OPENING BUCKET")
		panic(err)
	}

	fmt.Println("==> Enter a JSON document:")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		document := scanner.Text()

		key, _ := uuid()

		cas, err := bucket.Upsert(key, &document, 0)

		if err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Printf("Inserted document CAS is `%08x`\n", cas)
		}

		fmt.Println("==> Enter a JSON document:")
	}
}
