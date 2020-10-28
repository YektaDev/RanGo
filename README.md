[![PkgGoDev](https://pkg.go.dev/badge/github.com/yektadev/rango)](https://pkg.go.dev/github.com/yektadev/rango)
[![Go Report Card](https://goreportcard.com/badge/github.com/yektadev/rango)](https://goreportcard.com/report/github.com/yektadev/rango)
[![Go version](https://img.shields.io/badge/Go-v1.15-blue)](#)
[![Go coverage](https://img.shields.io/badge/Go%20Coverage-85.9%25-brightgreen)](#)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blueviolet)](#)
---
[![RanGo Logo](https://raw.githubusercontent.com/YektaDev/RanGo/main/res/RanGo_small.jpg "RanGo!")](#)

# 🦎RanGo
A **time-independent** random _number_/_string_ generator.

## 📦Installation
**First:**
```
go get github.com/yektadev/rango
```
**Then, add the dependency to your code:**
```go
import "github.com/yektadev/rango"
```

## 📝Simple Usage
Here's all you need to do in order to:

### 🔺Generate a random int:
**Use:**
```go
RanGo.RnInt(startIncluded int, endNotIncluded int)
```
Example:
```go
r := RanGo.RnInt(0,8)  //r ϵ {0,1,2,...,6,7}
```

### 🔺Generate a random string from a set of characters:
**Use:**
```go
RanGo.RnStringFrom(length int, chars string)
```
Example:
```go
r := RanGo.RnStringFrom(8,"abcd")  //r (example): "dadaadbb"
```

### 🔺Generate a random string:
**Use:**
```go
RanGo.RnString(length int, containsLowercase bool, containsUppercase bool, containsNumber bool, containsSpecial bool)
```
Examples:
```go
r := RanGo.RnString(18,true,true,true,true)  //r (example): "}WCg*(?w4P$<HS\jOb"
```
```go
r := RanGo.RnString(18,true,false,false,false)  //r (example): "jzoqagpchhsyhotvrj"
```
```go
r := RanGo.RnString(18,false,false,true,false)  //r (example): "325803510203358683"
```

## 🔧More Options
If the above functions face an error while generating a _time-independent_ seed, the seed will be automatically generated using ```time.Now().UnixNano()```. If you need to know whether the seed is generated _time-dependent_ or _time-independent_, then use the following functions:
* **```RandomInt()```** instead of ```RnInt()```.
* **```RandomStringFrom()```** instead of ```RnStringFrom()```.
* **```RandomString()```** instead of ```RnString()```.

In case of using these functions, you'll have **```isSeedTimeDependent```** boolean as the second returned output.

---

###### Special thanks to _John Leidegren_ because of [this](https://stackoverflow.com/a/54491783/6155559) answer on _Stack Overflow_.
[![](https://forthebadge.com/images/badges/it-works-why.svg)](#)
