/*
 * @Author: Y
 * @Date: 2022-07-22
 * @Description:
 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if context, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", context)
	}
}
