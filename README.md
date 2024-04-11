# catui

Chat TUI interface for Go Projects with cats

## Current usage

To print something to the console

``` go
cat := Cat {
  Name: "Name" // Cat Name
  Icon: "cat1" // Icon Name
  IsBot: false // change to preferred mode
}

cat.Say("Meow")
```

To get  user input

``` go
important_words := catui.Read() // returns a string
```