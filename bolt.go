package main

import (
	"bytes"
	"fmt"
	"github.com/boltdb/bolt"
	"strings"
)

func boltGet(key string) string {

	var result []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(key))
		result = v
		return nil
	})

	return string(result)
}

func boltSet(key, value string) error {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(key), []byte(value))
		return err
	})

	return err
}

func boltDelete(key string) error {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(key))
		return err
	})

	return err
}

func boltScanExt() {
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		c := tx.Bucket([]byte(bucketName)).Cursor()

		prefix := []byte("ext:")
		fmt.Println("Extension	|	Folder")
		for k, v := c.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = c.Next() {

			ext := strings.Split(string(k), ":")

			fmt.Printf("%s		|	%s\n", ext[1], v)
		}

		return nil
	})
}
