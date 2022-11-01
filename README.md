<h1><p align="center">Godog with Cucumber - API Testing 🌭🐕</p></h1>

<p align="center"><img src="logo.png" alt="Godog logo" style="width:200px;" /></p>


This project is an API Testing example using Go, [Godog](https://github.com/cucumber/godog) and [Cucumber](https://cucumber.io/).

## Features 🧪

-   Go
-   Gherkin - Cucumber

## Requirements 📚

-   [go1.17.2 darwin/amd64](https://www.npackd.org/p/org.golang.Go64/1.17.2).
-   Link Instalação (oficial) `https://go.dev/doc/install`.
-   Github dos estudos `https://github.com/cucumber/godog`

## Getting Started

Linux e macOS:

```
1. Open the package file you downloaded and follow the prompts to install Go.
The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.

2. Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version

3.Confirm that the command prints the installed version of Go.
```

Windows:

```
1. Open the MSI file you downloaded and follow the prompts to install Go.
By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

2. Verify that you've installed Go.
  In Windows, click the Start menu.
  In the menu's search box, type cmd, then press the Enter key.
  In the Command Prompt window that appears, type the following command:
  $ go version

3. Confirm that the command prints the installed version of Go.
```

After installing the Go version, you must:

Install Godog
```bash
go install github.com/cucumber/godog/cmd/godog@latest
```

<b>Run tests:</b>
with this command you will be able to run all your test suites based on cucumber.

```bash
godog run
```


## Contributing 👨‍💻🤝

Feel free to complement/report something in pull requests. The project is ours! 🤝
