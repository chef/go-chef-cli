/*
Copyright © 2022 Progress Software Corporation and/or its subsidiaries or affiliates. All Rights Reserved.
Author: Marc A. Paradise <marc.paradise@gmail.com>

Licensed under the Apache License, version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
LIMITATIONS UNDER THE LICENSE.
*/

package main

import (
	"fmt"
	"os"

	"github.com/chef/go-chef-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
