package main

import (
    "fmt"
    //"time"
    // "hotuananh.com/hourshelper"
    "github.com/hotuananh3010/go-hours-helper/v2"
)

func main() {
    // Get a greeting message and print it.
    times := hourshelper.Create("06:00", "07:00", 5, "15:04")
    fmt.Println(times)
}