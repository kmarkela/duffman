# DuffMan: Diagnostic Utility for Fuzzing and Fault Management of API Nodes

<p align="center">
  <img src="./assets/duffman.png" alt="DuffMan"/>
</p>

DuffMan is a tool written in Go that allows users to parse Postman collections and perform fuzz testing on all the endpoints defined within. This tool is designed to help developers and security analysts discover potential vulnerabilities and ensure robust error handling in their APIs.

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
  -f, --collection string   path to collection file
  -e, --enviroment string   path to enviroment file
  -h, --help                help for DuffMan

Use "DuffMan [command] --help" for more information about a command.
```

### Parse

Parses Postman Collection and Enviroment files and print Requests/Variables/etc defined within.

```sh
parse only collection file

Usage:
  DuffMan parse [flags]

Flags:
  -h, --help            help for parse
      --output string   output type. Possible values: brief, req, full (default "req")

Global Flags:
  -f, --collection string   path to collection file
  -e, --enviroment string   path to enviroment file
```

#### Example

```sh
duffman parse -e test/testing_environment.json -f test/testing_collection.json

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

[*] Envoriment:
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

### License 

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

### Disclamer 

The Postman Collection Fuzzer is intended for security research and testing purposes only. This tool should only be used on systems that you own or are explicitly authorized to test. Ethical conduct is required from all users.

The author(s) of this tool take no responsibility for any misuse of the software. It is the end user's responsibility to comply with all applicable local, state, federal, and international laws. By using this tool, you agree that you hold responsibility for any consequences that arise from its use.

### Contributing
