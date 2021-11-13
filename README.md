Terraform Provider for ☁️Zscaler Internet Access (ZIA)☁️
=========================================================================

⚠️  **Attention:** This provider is not affiliated with, nor supported by Zscaler in any way.

- Website: <https://www.terraform.io>
- Documentation: <https://help.zscaler.com/zia>
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)
<img src="https://github.com/hashicorp/terraform-website/blob/master/content/source/assets/images/logo-terraform-main.svg" width="600px">

Requirements
------------

- Install [Terraform](https://www.terraform.io/downloads.html) 0.12.x/0.13.x/0.14.x/0.15.x (0.11.x or lower is incompatible)
- Install [Go](https://golang.org/doc/install) 1.16+ (This will be used to build the provider plugin.)
- Create a directory, go, follow this [doc](https://github.com/golang/go/wiki/SettingGOPATH) to edit ~/.bash_profile to setup the GOPATH environment variable)

Building The Provider (Terraform v0.12+)
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-zia`

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone https://github.com/terraform-providers/terraform-provider-zia.git
```

To clone on windows

```sh
mkdir %GOPATH%\src\github.com\terraform-providers
cd %GOPATH%\src\github.com\terraform-providers
git clone https://github.com/terraform-providers/terraform-provider-zia.git
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-zia
make fmt
make build
```

To build on Windows

```sh
cd %GOPATH%\src\github.com\terraform-providers\terraform-provider-zia
go fmt
go install
```

Building The Provider (Terraform v0.13+)
-----------------------

### MacOS / Linux

Run the following command:

```sh
make build13
```

### Windows

Run the following commands for cmd:

```sh
cd %GOPATH%\src\github.com\terraform-providers\terraform-provider-zia
go fmt
go install
xcopy "%GOPATH%\bin\terraform-provider-zia.exe" "%APPDATA%\terraform.d\plugins\zscaler.com\zia\zia\1.0.0\windows_amd64\" /Y
```

Run the following commands if using powershell:

```sh
cd "$env:GOPATH\src\github.com\willguibr\terraform-provider-zia"
go fmt
go install
xcopy "$env:GOPATH\bin\terraform-provider-zia.exe" "$env:APPDATA\terraform.d\plugins\zscaler.com\zia\zia\1.0.0\windows_amd64\" /Y
```

**Note**: For contributions created from forks, the repository should still be cloned under the `$GOPATH/src/github.com/terraform-providers/terraform-provider-zia` directory to allow the provided `make` commands to properly run, build, and test this project.

Using Zscaler Internet Access Provider (Terraform v0.12+)
-----------------------

Activate the provider by adding the following to `~/.terraformrc` on Linux/Unix.

```sh
providers {
  "zia" = "$GOPATH/bin/terraform-provider-zia"
}
```

For Windows, the file should be at '%APPDATA%\terraform.rc'. Do not change $GOPATH to %GOPATH%.

In Windows, for terraform 0.11.8 and lower use the above text.

In Windows, for terraform 0.11.9 and higher use the following at '%APPDATA%\terraform.rc'

```sh
providers {
  "zia" = "$GOPATH/bin/terraform-provider-zia.exe"
}
```

If the rc file is not present, it should be created

Using Zscaler Internet Access Provider (Terraform v0.13+)
-----------------------

For Terraform v0.13+, to use a locally built version of a provider you must add the following snippet to every module
that you want to use the provider in.

```hcl
terraform {
  required_providers {
    zia = {
      source  = "zscaler.com/zia/zia"
      version = "1.0.0"
    }
  }
}
```

Examples
--------

Visit [here](https://github.com/willguibr/terraform-provider-zia/tree/master/website/docs/) for the complete documentation for all resources on github.

Issues
=========

Please feel free to open an issue using [Github Issues](https://github.com/willguibr/terraform-provider-zia/issues) if you run into any problems using this zia Terraform provider.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.16+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-zia
...
```

In order to test the provider, you can simply run `make test`.

```sh
make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
make testacc
```

License
=========

MIT License

Copyright (c) 2021 [William Guilherme](https://github.com/willguibr)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
