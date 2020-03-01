package gohome

// An attempt to build a minimal framework for building home automation systems with go.
//
// Normally home automation is configurable with minimal scripting, this works for the
// non software engineer kind of home automator but it is also very limiting for anyone
// who can write code.
//
// Assumptions:
// - This library is used to write control functions and loops using go
// - Channels and goroutines are part of the design
// - Common interfaces for well known real life elemts exits
// - RESTful interfaces are exposed 'by default' for simple external integration
// - Monitoring with prometheus should be 'default'
// - Broken sensor, bulbs or actuators must not prevent system startup
