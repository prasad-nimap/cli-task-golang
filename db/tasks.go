package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

/* Read-write transactions(using bolt) */
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, nil
	}
	return id, nil
}

/* Read-only transaction(using bolt) */
func AllTask() ([]Task, error) {
	var task []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket) // Grab the bucket
		/* Iterating over keys */
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			/* For every key value pair add to task */
			task = append(task, Task{
				Key:   btio(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return task, nil
}

/* Delete task (Delete is Read-write transaction) */
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

/* Convert integer to binary */
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

/* Convert binary to integer */
func btio(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
