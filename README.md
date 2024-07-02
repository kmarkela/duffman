# DuffMan: Diagnostic Utility for Fuzzing and Fault Management of API Nodes

[![Go Report Card](https://goreportcard.com/badge/github.com/kmarkela/duffman)](https://goreportcard.com/report/github.com/kmarkela/duffman)

<p align="center">
  <img src="./assets/duffman.png" alt="DuffMan"/>
</p>

DuffMan is a tool written in Go that allows users to parse Postman collections and perform fuzz testing on all the endpoints defined within. This tool is designed to help developers and security analysts discover potential vulnerabilities and ensure robust error handling in their APIs.

## Install 

There are three ways to install the DuffMan: building from source, using `go get`, or downloading pre-compiled binaries.

### Building from Source

```sh
git clone https://github.com/kmarkela/duffman
cd duffman
go build -o duffman
```

### Using `go get`

```sh
go get github.com/kmarkela/duffman
```

### Binaries 

If you prefer to download a pre-compiled binary for your platform, follow these steps:

1. Navigate to the [Releases page](https://github.com/kmarkela/duffman/releases) of the DuffMan repository.
2. Download the appropriate binary for your operating system and architecture.
3. Unzip the downloaded file and move the binary to your desired location.


## Usage

```sh
Diagnostic Utility for Fuzzing and Fault Management of API Nodes

Usage:
  DuffMan [flags]
  DuffMan [command]

Available Commands:
  fuzz        fuzz all endpoint from Postman Collection
  help        Help about any command
  parse       parse only collection file
  version     Print Version

Flags:
  -f, --collection  string   path to collection file
  -e, --environment string   path to environment file
  -h, --help                 help for DuffMan

Use "DuffMan [command] --help" for more information about a command.
```

### Parse

Parses Postman Collection and Environment files and print Requests/Variables/etc defined within.

```sh
parse only collection file

Usage:
  DuffMan parse [flags]

Flags:
  -h, --help            help for parse
      --output string   output type. Possible values: brief, req, full (default "req")

Global Flags:
  -f, --collection  string   path to collection file
  -e, --environment string   path to environment file
```

#### Example

```sh
duffman parse -e test/testing_environment.json -f test/testing_collection.json

```

```sh

 ####                                       ###
 ######                                   #######
 ########                       ######   #########
 ##########                    ########  ###   ##
 ####  #####                   ###  ###  ###
 ####   #####                  ###       ###
 ####    #####                 ###       ###
 ####     #####                ###       ###
 ####     #####                ###       ###
 ####      #####               ###       ###
 ####      #####               ###       ###       ###
 ####      #####  ###   ####   ###       ###    ######
 ####      #####  ###   ####   ###       ### #######
 ####      #####  ###   ####   ###       ########
 ####       ####  ###   ####   ###      ######
 ####       ####  ###   ####   ###   ########
 ####       ####  #### #####   ##############
 ####      #####  #########    #######   ####
 ####      #####   ########  ######       ###
 ####      ####      ###  ########        ###
 ####     #####        ####### ###        ###
 ####    #####       ######    ###   ###  ###
 ####   #####      #####    #   ##  #### ####
 #### ######      ###     ####  ##  ########
 #########                ########   ######
 #######                   #######
 #####                      ####

[*] Environment:
  - env1: 9999
  - env2: 8888
  - env3: 7777
[*] Variables:
  - testing: 123456
[*] Req amount: 9
[*] Requests:
  - URL: http://foo.bar/3-sub/post/raw-json
  - URL: http://foo.bar/2-sub/post/raw-text
  - URL: http://foo.bar/2-sub/post/raw_params
  - URL: http://foo.bar/2-sub/post/form_params
  - URL: http://foo.bar/2-sub/post/urlen_params_header
  - URL: http://foo.bar/1-sub/get/var/1111/2222
  - URL: http://foo.bar/get/var/1111/2222
  - URL: http://foo.bar/get/variable/1111/2222
  - URL: http://foo.bar/env
```

### Fuzz

```sh
It allows to fuzz muptiple parameters over multiple endpoints

Usage:
  DuffMan fuzz [flags]

Flags:
      --headers strings               replace header if exists, add if it wasn't in original request
  -h, --help                          help for fuzz
  -r, --hide-response-time int        hide requests with response time less than X in milliseconds
  -m, --maxReq int                    max amount of requests per second
  -p, --proxy string                  proxy
  -b, --status-codes-blacklist ints   hide responses with specified status codes
      --variables strings             replace variables value
  -l, --wordlist string               wordlits to fuzz
  -w, --workers int                   amount of workers (default 10)

Global Flags:
  -f, --collection string    path to collection file
  -e, --environment string   path to environment file
```

#### Example

```sh
go run main.go fuzz -f test/testing_collection.json -e test/testing_environment.json --headers "User-Agent: duffman" --headers "X-Fuzz: test" -m 40 -w 100 -p http://127.0.0.1:8080 -l ~/1.lst -b 404,401
```

```sh

 ####                                       ###
 ######                                   #######
 ########                       ######   #########
 ##########                    ########  ###   ##
 ####  #####                   ###  ###  ###
 ####   #####                  ###       ###
 ####    #####                 ###       ###
 ####     #####                ###       ###
 ####     #####                ###       ###
 ####      #####               ###       ###
 ####      #####               ###       ###       ###
 ####      #####  ###   ####   ###       ###    ######
 ####      #####  ###   ####   ###       ### #######
 ####      #####  ###   ####   ###       ########
 ####       ####  ###   ####   ###      ######
 ####       ####  ###   ####   ###   ########
 ####       ####  #### #####   ##############
 ####      #####  #########    #######   ####
 ####      #####   ########  ######       ###
 ####      ####      ###  ########        ###
 ####     #####        ####### ###        ###
 ####    #####       ######    ###   ###  ###
 ####   #####      #####    #   ##  #### ####
 #### ######      ###     ####  ##  ########
 #########                ########   ######
 #######                   #######
 #####                      ####

########################################################
#                        DuffMan                       #
# [*] Wordlist count: 3                                #
# [*] Amount of request: 9                             #
# [*] Amount of parameters: 16                         #
# [*] Total to fuzz: 48                                #
# [*] Status Code Blacklist: 404,401                   #
########################################################

+-----------------------------------------------+--------+-------------------+------+------+--------+------+
| ENPOINT                                       | METHOD | PARAMETER         | FUZZ | CODE | LENGTH | TIME |
+-----------------------------------------------+--------+-------------------+------+------+--------+------+
| http://foo.bar/3-sub/post/raw-json            | POST   | test1             | pwn2 |  501 |    357 | 44ms |
| http://foo.bar/3-sub/post/raw-json            | POST   | test1             | pwn1 |  501 |    357 | 69ms |
| http://foo.bar/3-sub/post/raw-json            | POST   | test1             |      |  501 |    357 | 24ms |
| http://foo.bar/3-sub/post/raw-json            | POST   | test2.test3.test4 | pwn1 |  501 |    357 | 14ms |
| http://foo.bar/3-sub/post/raw-json            | POST   | test2.test3.test4 | pwn2 |  501 |    357 | 15ms |
| http://foo.bar/3-sub/post/raw-json            | POST   | test2.test3.test4 |      |  501 |    357 | 29ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | testing-param     | pwn1 |  501 |    357 | 11ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | testing-param     | pwn2 |  501 |    357 | 12ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | testing-param     |      |  501 |    357 | 18ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | test              | pwn1 |  501 |    357 | 92ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | test              | pwn2 |  501 |    357 | 89ms |
| http://foo.bar/2-sub/post/form_params         | POST   | testing-param     | pwn1 |  501 |    357 | 49ms |
| http://foo.bar/2-sub/post/raw_params          | POST   | test              |      |  501 |    357 | 78ms |
| http://foo.bar/2-sub/post/form_params         | POST   | testing-param     | pwn2 |  501 |    357 | 38ms |
| http://foo.bar/2-sub/post/form_params         | POST   | testing-param     |      |  501 |    357 | 18ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | testing-param     | pwn1 |  501 |    357 | 12ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | testing-param     | pwn2 |  501 |    357 | 21ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | testing-param     |      |  501 |    357 | 14ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | params            | pwn1 |  501 |    357 | 16ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | params            | pwn2 |  501 |    357 | 11ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | params            |      |  501 |    357 | 18ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | param2            | pwn1 |  501 |    357 | 13ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | param2            | pwn2 |  501 |    357 | 15ms |
| http://foo.bar/2-sub/post/urlen_params_header | POST   | param2            |      |  501 |    357 | 17ms |
+-----------------------------------------------+--------+-------------------+------+------+--------+------+

[-] 3 Errors occur during Fuzz:
  - Endpoint http://foo.bar/2-sub/post/form_params:
    * Param: post
    * Error: no encoder for: multipart/form-data; boundary=------border
  - Endpoint http://foo.bar/2-sub/post/form_params:
    * Param: post
    * Error: no encoder for: multipart/form-data; boundary=------border
  - Endpoint http://foo.bar/2-sub/post/form_params:
    * Param: post
    * Error: no encoder for: multipart/form-data; boundary=------border
```

### License 

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

### Disclamer 

The DuffMan is intended for security research and testing purposes only. This tool should only be used on systems that you own or are explicitly authorized to test. Ethical conduct is required from all users.

The author(s) of this tool take no responsibility for any misuse of the software. It is the end user's responsibility to comply with all applicable local, state, federal, and international laws. By using this tool, you agree that you hold responsibility for any consequences that arise from its use.

### Contributing

We welcome contributions! Please read our [CONTRIBUTING](./CONTRIBUTING.md) for details on how to submit pull requests, the process for submitting bugs, feature requests, and how to propose changes.

### Authors

- Kanstantsin Markelau 
- Sam Cosentino 