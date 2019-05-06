package utils

import (
	"fmt"
	"time"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  07-May-2019
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func NewClientId() string {
	return fmt.Sprintf("mqtt-sh-%s", time.Now().Format(time.RFC3339))
}
