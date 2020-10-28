![Go version](https://img.shields.io/badge/Go-v1.15-blue)
![GitHub license](https://img.shields.io/badge/license-Apache%202-blueviolet)

<img src="https://raw.githubusercontent.com/YektaDev/RanGo/main/res/RanGo.jpg" alt="RanGo!" width="500">

# RanGo
A **time-independant** safe random number/string generator.

## Installation
First:
```
go get github.com/YektaDev/RanGo
```
Then, add the dependency to your code:
```go
import "github.com/YektaDev/RanGo"
```
## Simple Usage
Here's all you need to do in order to:

### Generate a random int:
**Use:**
```go
RnInt(startIncluded int, endNotIncluded int)
```
Example:
```go
r := RnInt(0,8)  //r ϵ {0,1,2,...,7,8}
```
### Generate a random string from a set of characters:
**Use:**
```go
RnStringFrom(length int, chars string)
```
Example:
```go
r := RnStringFrom(8,"abcd")  //r (example): "dadaadbb"
```
### Generate a random string:
**Use:**
```go
RnString(length int, containsLowercase bool, containsUppercase bool, containsNumber bool, containsSpecial bool)
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
