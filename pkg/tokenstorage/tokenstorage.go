/*
Copyright © 2019 ARTEM KOLOMYTSEV kolomytsev1996@gmail.com

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
package tokenstorage

import "github.com/pkg/errors"

// TokenStorage defines interface for auth token persistent storage
type TokenStorage interface {
	// Delete old tokens and save new
	Save(token string) error
	// Read and return token from storage
	// Returns ErrTokenNotFound if token not found
	Read() (string, error)
}

var ErrTokenNotFound = errors.New("Auth token not found")
