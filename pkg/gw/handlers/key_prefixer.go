/*
Copyright 2019-2020 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gwhandlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/auth"
)

func getKeyPrefix(req *http.Request) ([]byte, error) {
	if len(req.URL.Query()["multi-tenant"]) <= 0 {
		return nil, nil
	}
	token := strings.TrimPrefix(req.Header.Get("Authorization"), "Bearer ")
	jsonToken, err := auth.ParsePublicTokenPayload(token)
	if err != nil {
		return nil, fmt.Errorf("error parsing public payload of auth token: %v", err)
	}
	prefix := append([]byte(jsonToken.Username), ':')
	return prefix, nil
}

func prefixKey(req *http.Request, key []byte) ([]byte, []byte, error) {
	prefix, err := getKeyPrefix(req)
	if err != nil {
		return nil, nil, err
	}
	if prefix == nil {
		return key, nil, nil
	}
	prefixedKey := bytes.Join([][]byte{prefix, key}, nil)
	return prefixedKey, prefix, nil
}

func prefixKeys(req *http.Request, keys []*schema.Key) ([][]byte, []byte, error) {
	prefix, err := getKeyPrefix(req)
	if err != nil {
		return nil, nil, err
	}
	keysBytes := make([][]byte, len(keys))
	if prefix == nil {
		for i, key := range keys {
			keysBytes[i] = key.Key
		}
		return keysBytes, nil, nil
	}
	for i, key := range keys {
		keysBytes[i] = bytes.Join([][]byte{prefix, key.Key}, nil)
	}
	return keysBytes, prefix, nil
}
