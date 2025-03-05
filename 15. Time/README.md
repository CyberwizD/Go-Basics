# Working with Time in Go

## Overview
In Go, `Time` is a type that describes a moment in time. The date and time information can be accessed, compared, and manipulated using its methods. Additionally, the `time` package provides functions to work with `Time` values, such as retrieving the current time using `time.Now()`.

The `time.Parse` function is used to parse strings into values of type `Time`. Go requires a special layout format for parsing dates, which is based on the reference timestamp:

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

This timestamp is used to define the layout for `time.Parse` and `time.Format`.

## Example: Parsing Time

```go
import "time"

func parseTime() time.Time {
    date := "Tue, 09/22/1995, 13:00"
    layout := "Mon, 01/02/2006, 15:04"

    t, err := time.Parse(layout, date) // time.Time, error

    if err != nil {
        panic(err)
    }

    return t
}
```

## In this example:

We are parsing a date string Tue, 09/22/1995, 13:00.
The layout "Mon, 01/02/2006, 15:04" represents the expected format of the date string.
The parsed result is of type time.Time.

### Output:
```go
1995-09-22 13:00:00 +0000 UTC
```

## Example: Formatting Time
The Time.Format() method is used to return a string representation of a Time object. The format is specified using a layout based on the special timestamp.

```go
import (
    "fmt"
    "time"
)

func main() {
    t := time.Date(1995, time.September, 22, 13, 0, 0, 0, time.UTC)
    formattedTime := t.Format("Mon, 01/02/2006, 15:04")
    fmt.Println(formattedTime)
}
```

### Output:
```go
Fri, 09/22/1995, 13:00
```

## Layout Options
You can use different combinations of layout options to customize the format. The following table shows some common layout options:

| **Time Component** | **Layout Options**                |
|--------------------|-----------------------------------|
| **Year**           | `2006`, `06`                      |
| **Month**          | `Jan`, `January`, `01`, `1`       |
| **Day**            | `02`, `2`, `_2` (for preceding 0) |
| **Weekday**        | `Mon`, `Monday`                   |
| **Hour**           | `15` (24-hour), `3`, `03` (AM/PM) |
| **Minute**         | `04`, `4`                         |
| **Second**         | `05`, `5`                         |
| **AM/PM Mark**     | `PM`                              |
| **Day of Year**    | `002`, `__2`                      |

## Accessing Specific Time Components
The time.Time type provides several methods to access particular components of time. Some of these methods include:

* **Hour()** - Returns the hour.
* **Month()** - Returns the month.
* **Day()** - Returns the day of the month.

## Additional Features of time Package
The time package in Go also includes:

### Duration: A type representing elapsed time.
Time zones and locations.
Timers and other related functionality.
You can explore more in the official documentation.
