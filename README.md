# Action Statistics (from a JumpCloud "Assignment")
A library to assist in keeping and reporting statistics for some set of timed actions.

# Specification
See [here](./JumpCloud%20-%20Backend%20Take%20Home%20Assignment%20(v2)%20-%204_17_2019.pdf).

# Installation
We assume that only Go applications will be using this library.

Use
```bash
$ go get github.com/cargregter/jc_assignment
```

# Testing
Unit and integration tests are provided. A couple of mechanisms
for running the tests are suggested.

## `go test`
From directory `.../jc_assignment/action`, use
```bash
$ go test --test.v
```

## GoConvey
[GoConvey](http://goconvey.co/) provides a nice framework within
which to define test cases along with lots of matchers and
a convenient web interface.

# Usage
## API
This library exposes two functions.

- `Add(action string) error`: Used to add the details of an action,
the `action` parameter is a JSON string like
```bash
`{"action":"run", "time":75}`
```
Units of the `time` value are your choice. An error is returned
if there is difficulty interpreting the input.
- `GetStats() string`: Used to obtain a report of the statistics
having been added. The report is in formatted JSON form like
```bash
[
    {
        "action": "jump",
        "avg": 200
    },
    {
        "action": "run",
        "avg": 75
    }
]
```

## Example Code
In your code, add lines like the following as appropriate:
```go
...

import "github.com/cargregter/jc_assignment/action"

    ...
    errAdd := action.Add(`{"action":"jump", "time":100}`)
    if errAdd != nil {
    	return errAdd
    }
    
    ...
    statsReport := action.GetStats()
    if statsReport == "" {
    	return error.New("error generating statistics report")
    }
    
    ...
```

# Notes
## Parallelism
Both functions `Add()` and `GetStats()` are [thread safe](https://en.wikipedia.org/wiki/Thread_safety).
However, keep in mind that if you invoke `GetStats()` while
`Add()` is being invoked (or could be invoked) in parallel,
the returned statistics report may or may not include the latest
added actions.

## Storage
Added actions are kept in memory. There are no provisions
made to recover stored actions in the event that the application
using this library should fail.

# Improvements
## Persistence
There may be value in writing actions to a persistent store
or at least offering the option.

## More Statistics
Of course, more statistics could be added as thought useful.
It would likely mean capturing more raw data. But even with
what we currently capture, one could see the usefulness of
including the total number of observations for an action, and
perhaps the mode and median of the provided times.

## Editing Actions
Currently actions can only be added. There is no contemplation
of modifying or removing them, except that each added action
is kept separately.

## Dumping Action Detail
While all action detail is kept, there is currently no
mechanism for dumping them for inspection.
