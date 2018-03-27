// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/hidevopsio/hi/boot/pkg/application"
	"github.com/hidevopsio/hi/cicd/pkg/web/controllers"
)

func init() {
	// iris app
	app := application.Instance()

	user := controllers.UserController{}

	app.Post("/user/login", user.Login)

	application.ApplyJwt()

	// cicd
	cicd := controllers.CicdController{}
	cicdRouters := app.Party("/cicd", cicd.Before)
	{
		// Method POST: http://localhost:8080/deployment/deploy
		cicdRouters.Post("/run", cicd.Run)
	}
}

func Run()  {
	application.Run()
}


