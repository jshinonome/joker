/**
 * Copyright 2022 Jo Shinonome
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"time"

	"github.com/jshinonome/geek"
)

func main() {
	// connect to a q process @ 1800
	q := geek.QProcess{Port: 1800}
	q.Dial()
	sym := "a"
	f := struct {
		Api string
		Sym string
	}{
		"getTrade", sym,
	}
	r := make([]trade, 0)
	err := q.Sync(&r, f)
	if err != nil {
		fmt.Println(err)
	}
	for _, t := range r {
		fmt.Printf("%+v\n", t)
	}
}

type trade struct {
	Time  time.Time `k:"time"`
	Sym   string    `k:"sym"`
	Price float64   `k:"price"`
	Qty   int64     `k:"qty"`
}
