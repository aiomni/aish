# AISH

AISH is a command line tool powered by ChatGPT. You can input descriptive text
about what you want to do, and AISH will help you find the corresponding
command.

## install

use cURL install AISH:

```sh
curl -o- https://raw.githubusercontent.com/aiomni/aish/dev/install.sh | sh
```

or use wget:

```sh
wget -qO- https://raw.githubusercontent.com/aiomni/aish/dev/install.sh | sh
```

## Usage

Available Commands:

* ask         Ask AI what you want to do
* config      Set and Get your custom config
* completion  Generate the autocompletion script for the specified shell
* help        Help about any command
* version     Print the version number of AISH

### config

You can use `aish config` to set OpenAI API key、ORGANIZATION ID for use chatGPT.

Also provide PROXY DOMAIN config for Chinese user, for more detail: [link](https://foreverz.cn/chatgpt-guide#heading-openai-open-api)

### ask

You can use the `aish ask` command for interactive questioning：

```sh
 / aish ask
Please Input What you want todo: Get the current system time and output it in milliseconds format.
Found the command for you: date +%!s(MISSING)%!N(MISSING)
This command uses the `date` utility to get the current system time in seconds since the Unix epoch, and then multiplies it by 1000 to convert it to milliseconds. The `%3N` option is used to output the milliseconds with leading zeros if necessary.
Use the arrow keys to navigate: ↓ ↑ → ←
? Select What you want to do:
  ▸ Execute
    Copy
    Abort
```

Then you can choose execute the command directly or copy it.

## TODO

* [x] Check Config When ASK.
* [x] View the config settings.
* [x] Reset the config settings.
* [ ] Support for both interactive commands and flags.
* [ ] Custom Shell, eg:  `cat /etc/shells`.
* [ ] Proxy Service.
* [ ] i18n support.
* [ ] Find other command.
* [ ] When execute command error, ask chatgpt how to resolve.
* [ ] When the command is not installed, ask if it needs to be installed.
