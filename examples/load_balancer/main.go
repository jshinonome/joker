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

import "github.com/jshinonome/geek"

func main() {
	// connect to a q process @ 1800
	q1 := geek.QProcess{Port: 1800}
	q1.Dial()
	// connect to a q process @ 1801
	q2 := geek.QProcess{Port: 1801}
	q2.Dial()
	qConnPool := geek.NewConnPool()
	qConnPool.Put(&q1)
	qConnPool.Put(&q2)
	qConnPool.Serving()

	qEngine := geek.DefaultEngine(qConnPool)
	qEngine.Run()
}
