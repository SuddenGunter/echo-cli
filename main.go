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
package main

import (
	"github.com/SuddenGunter/echo-cli/cmd"
	"github.com/SuddenGunter/echo-cli/cmd/config"
	"github.com/SuddenGunter/echo-cli/pkg/echo"
	"github.com/SuddenGunter/echo-cli/pkg/tokenstorage"
)

func main() {
	tokenCfg := tokenstorage.DefaultTempFileTokenStorageConfig
	storage := tokenstorage.NewTempFileTokenStorage(tokenCfg)

	client := echo.NewClient("localhost:8080", "/ws")

	state := config.NewState(storage, cmd.UnauthorizedErrorHandler, client)
	cmd.Execute(state)

}
