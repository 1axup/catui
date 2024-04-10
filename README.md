# catui

Chat TUI interface for Go Projects

## Current usage

To print something to the console

``` go
cat := Cat {
  Name: "Name" // Cat Name
  Icon: "cat1" // Icon Name
  IsBot: false // change to preferred mode
}
```

To get  user input

``` go
catui.Read() // returns a string
```