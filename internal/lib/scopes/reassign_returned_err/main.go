package main

import (
	"errors"
	"fmt"
)

func main() {
	if err := foo(); err != nil {
		fmt.Print(err)
	}
}

func foo() (err error) {
	c := &Client{}

	defer func() {
		err = errors.Join(err, c.Close())
	}()

	err = c.Exec()

	return
}

type Client struct {
}

func (c *Client) Exec() error {
	return nil
}

func (c *Client) Close() error {
	return errors.New("1")
}
