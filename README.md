![Go version](https://img.shields.io/badge/Go-v1.15-blue)
![GitHub license](https://img.shields.io/badge/license-Apache%202-blueviolet)

<img src="https://raw.githubusercontent.com/YektaDev/RanGo/main/res/RanGo.jpg" alt="RanGo!" width="500">

# RanGo
A time-independant safe random number/string generator.

## Installation
First:
```
go get github.com/YektaDev/RanGo
```
Then, add the dependency to your code:
```
import "github.com/YektaDev/RanGo"
```
## Simple Usage
Here's all you need to do in order to:

### Generate a random int:
Use:
```go
RnInt(startIncluded int, endNotIncluded int)
```
Example:
```go
RnInt(0,8)  //output ϵ {0,1,2,...,7,8}
```
### Generate a random string from a set of characters:
Use:
```go
RnStringFrom(length int, chars string)
```
Example:
```go
RnStringFrom(8,"abcd")  //output (example): "dadaadbb"
```
