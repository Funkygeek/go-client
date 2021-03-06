/*
 * Copyright (c) 2020, FusionAuth, All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific
 * language governing permissions and limitations under the License.
 */
package fusionauth

import (
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	host string = "http://localhost:9011"
)

var (
	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	baseURL, _ = url.Parse(host)
)

var faClient = NewClient(httpClient, baseURL, "af69486b-4733-4470-a592-f1bfce7af580")

func TestRetrieveUser(t *testing.T) {
	userResponse, _, _ := faClient.RetrieveUser("missing@example.com")
	assert.Equal(t, 401, userResponse.StatusCode)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
