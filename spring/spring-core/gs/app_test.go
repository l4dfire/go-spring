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

package gs_test

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/go-spring/spring-core/assert"
	"github.com/go-spring/spring-core/conf"
	"github.com/go-spring/spring-core/environ"
	"github.com/go-spring/spring-core/gs"
)

func startApplication(cfgLocation string) (*gs.App, gs.Pandora) {

	app := gs.NewApp()
	gs.Setenv("SPRING_BANNER_VISIBLE", true)
	gs.Setenv("SPRING_CONFIG_LOCATION", cfgLocation)
	app.Property(environ.EnablePandora, true)

	var p gs.Pandora
	type PandoraAware struct{}
	app.Provide(func(b gs.Pandora) PandoraAware {
		p = b
		return PandoraAware{}
	})

	go app.Run()
	time.Sleep(100 * time.Millisecond)
	return app, p
}

func TestConfig(t *testing.T) {

	t.Run("config via env", func(t *testing.T) {
		os.Clearenv()
		gs.Setenv("SPRING_ACTIVE_PROFILE", "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(environ.SpringActiveProfile), "dev")
	})

	t.Run("config via env 2", func(t *testing.T) {
		os.Clearenv()
		gs.Setenv("SPRING_ACTIVE_PROFILE", "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(environ.SpringActiveProfile), "dev")
	})

	t.Run("profile via env&config 2", func(t *testing.T) {

		os.Clearenv()
		gs.Setenv("SPRING_ACTIVE_PROFILE", "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(environ.SpringActiveProfile), "dev")

		var m map[string]string
		_ = p.Bind(&m, conf.Key(conf.RootKey))
		for _, k := range sortedKeys(m) {
			fmt.Println(k, "=", p.Prop(k))
		}
	})
}

func sortedKeys(m map[string]string) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return
}
