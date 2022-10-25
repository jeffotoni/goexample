package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	//"time"
)

//Function to Calculate a Current_spend

func Current_spend(current_spend []uint8) uint8 {
	var c_s uint8
	for u := 0; u < len(current_spend); u++ {
		c_s = c_s + current_spend[u]
	}

	//UpdateCurrentSpend(c_s,q)
	return c_s
}

func UpdateCurrentSpend(csi uint8, l int) {
	var ur []uint8
	ur = append(ur, csi)
	s := make([]string, 20, 20)
	s = []string{"abcdefgh01", "abcdefgh02", "abcdefgh03", "abcdefgh04", "abcdefgh05", "abcdefgh06", "abcdefgh07", "abcdefgh08", "abcdefgh09", "abcdefgh10", "abcdefgh11", "abcdefgh12", "abcdefgh13", "abcdefgh14", "abcdefgh15", "abcdefgh16", "abcdefgh17", "abcdefgh18", "abcdefgh19", "abcdefgh20"}

	var tr string
	tr = s[l]
	db, err := bolt.Open("t22.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("ADSERVER"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ADSERVER"))
		err := b.Put([]byte(tr), []byte(ur[0:1]))
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ADSERVER"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%d\n", k, v)
		}

		return nil
	})
}

func main() {
	s := make([]string, 20, 20)
	s = []string{"abcdefgh01", "abcdefgh02", "abcdefgh03", "abcdefgh04", "abcdefgh05", "abcdefgh06", "abcdefgh07", "abcdefgh08", "abcdefgh09", "abcdefgh10", "abcdefgh11", "abcdefgh12", "abcdefgh13", "abcdefgh14", "abcdefgh15", "abcdefgh16", "abcdefgh17", "abcdefgh18", "abcdefgh19", "abcdefgh20"}

	//uint slice of values i.e.bids
	bids := make([]uint8, 20, 20)
	bids = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	//uint slice of keys i.e. timestamp
	time_stamp := make([]uint8, 20, 20)
	time_stamp = []uint8{20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	//Database Access Codes

	db, err := bolt.Open("t22.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	x := 0
	for x < 20 {
		db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(s[x]))
			if err != nil {
				return err
			}
			return nil
		})
		x++
	}
	z := 0
	for z < 20 {
		y := 1
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(s[z]))
			for y < 21 {
				b.Put(bids[y-1:y], time_stamp[y-1:y])
				y++
			}
			return err
		})
		z++
	}
	v := 0
	for v < 20 {

		cs := make([]uint8, 20)
		db.View(func(tx *bolt.Tx) error {
			j := tx.Bucket([]byte(s[v]))
			c := j.Cursor()
			for k, val := c.First(); k != nil; k, val = c.Next() {
				fmt.Println(k, val)

				cs = append(cs, val[0])
			}
			return nil
		})
		var g_cs uint8
		g_cs = Current_spend(cs)
		UpdateCurrentSpend(g_cs, v)
		v++
	}
}
