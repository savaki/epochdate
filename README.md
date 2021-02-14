epochdate 
--------------------------------------

`epochdate` calculates the number of days since the start of unix epoch, Jan 1, 1970.  

## Motivation

Calculating the number of days between two points can often be a tricky operation because of leap years
and leap seconds.  Because of this, it can often be simpler to pre-calculate the date of an event
as an epoch date.  This is especially true if we don't want to worry about the challenges of incorporating
time zone calculations into date.

## Usage

```go
t := time.Now()
d := epochdate.From(t) // d contains the number of days since Jan 1, 1970
```