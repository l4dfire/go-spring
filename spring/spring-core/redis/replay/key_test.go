/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package replay

import (
	"testing"

	"github.com/go-spring/spring-core/redis/testcases"
	"github.com/go-spring/spring-core/redis/testdata"
)

func TestDel(t *testing.T) {
	RunCase(t, testcases.Del, testdata.Del)
}

func TestDump(t *testing.T) {
	RunCase(t, testcases.Dump, testdata.Dump)
}

func TestExists(t *testing.T) {
	RunCase(t, testcases.Exists, testdata.Exists)
}

func TestExpire(t *testing.T) {
	RunCase(t, testcases.Expire, testdata.Expire)
}

func TestExpireAt(t *testing.T) {
	RunCase(t, testcases.ExpireAt, testdata.ExpireAt)
}

func TestKeys(t *testing.T) {
	RunCase(t, testcases.Keys, testdata.Keys)
}

func TestPersist(t *testing.T) {
	RunCase(t, testcases.Persist, testdata.Persist)
}

func TestPExpire(t *testing.T) {
	RunCase(t, testcases.PExpire, testdata.PExpire)
}

func TestPExpireAt(t *testing.T) {
	RunCase(t, testcases.PExpireAt, testdata.PExpireAt)
}

func TestPTTL(t *testing.T) {
	RunCase(t, testcases.PTTL, testdata.PTTL)
}

func TestRename(t *testing.T) {
	RunCase(t, testcases.Rename, testdata.Rename)
}

func TestRenameNX(t *testing.T) {
	RunCase(t, testcases.RenameNX, testdata.RenameNX)
}

func TestTouch(t *testing.T) {
	RunCase(t, testcases.Touch, testdata.Touch)
}

func TestTTL(t *testing.T) {
	RunCase(t, testcases.TTL, testdata.TTL)
}

func TestType(t *testing.T) {
	RunCase(t, testcases.Type, testdata.Type)
}
