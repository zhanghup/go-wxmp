package wxmp

import (
	"fmt"
	"testing"
)

func TestMenu_Create(t *testing.T) {
	err := c.Menu().Create([]Button{
		{Type: "click", Name: "今日金曲1", Key: "V1001_TODAY_MUSIC"},
		{Type: "click", Name: "今日金曲2", Key: "V1001_TODAY_MUSIC"},
	})
	if err != nil {
		panic(err)
	}
}

func TestMenu_Delete(t *testing.T) {
	err := c.Menu().Delete()
	if err != nil {
		panic(err)
	}
}

func TestMenu_Get(t *testing.T) {
	buttons, err := c.Menu().Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(buttons)
}
