# godesk - A quick [Desk](https://github.com/jamesob/desk) workspace launcher with fuzzy filtering

![demo](https://github.com/hamin/godesk/blob/master/demo.gif) 

`godesk` is a simple terminal app to launch a [Desk](https://github.com/jamesob/desk) workspace. It offers a fuzzy filtering menu to select your [Desk](https://github.com/jamesob/desk) file.

`godesk` heavily uses the UI+Fuzzy filtering provided by [go-selecta](https://github.com/thomasboyt/go-selecta) for only one task, launching [Desk](https://github.com/jamesob/desk) workspaces.


## Requirements

Install [Desk](https://github.com/jamesob/desk) for your environment

## Installation Options

#### Homebrew

```shell
brew tap hamin/godesk https://github.com/hamin/godesk
brew install godesk
```    

#### `go get`
```shell
  go get -u github.com/hamin/godesk
```

## Usage

```shell
godesk
```

## License

`godesk` is licensed under the MIT License.

## Libraries

* [go-selecta](https://github.com/thomasboyt/go-selecta)
* [termbox-go](github.com/nsf/termbox-go)
* [cli](https://github.com/urfave/cli)
