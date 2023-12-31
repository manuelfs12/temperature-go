# Temperature GO
### Video Demo:  [Link](https://youtu.be/MRzZix2o_o4)
### What is the project?
This is a temperature conversion tool written in Go for the final project of CS50x.

![Gif of a terminal window showcasing the software](/temperature-go.gif)

The software takes a temperature value `50`, be it in Celsius `c`,
Fahrenheit `f` or Kelvin `k` and outputs the converted values as a formatted string.


### What each file on the project contains and does?
The project structure is as follows:
```
▾ temperature-go/
    .gitignore
    go.mod
    LICENSE
    main.go
    README.md
```
* **.gitignore**: Excludes unwanted files from the version control.


* **go.mod**: Describes the module properties, including its own dependencies, every modern go project hast it.


* **LICENSE**: The license file of the repository, in this case, is MIT license.


* **main.go**: This is the body of the application, where all the coding logic is in. I import some packages from the standard library, 
declare the `main()` function and all the other functions that makes the application work.

### Debated design choices: Using a framework vs standard library, Go vs other languages, etc.

#### Go vs Another language
Since before starting CS50x, I was toying with the idea of learning Go, but got scared when I read that it was
a combination of Python and C.
After going through CS50x, and dipping my toes in C,
this encouraged me to work with Go because I like static type languages.

I have previous experience working with languages like Python, Javascript and C#, But for this project I wanted to challenge
myself and create something in a new language.

#### Framework vs Standard Library
When I decided to use Go as my language for this project, I considered using a framework called [Cobra](https://github.com/spf13/cobra).
This is the recommended framework on Go website for creating modern and powerful CLI applications, and
it's the framework that empowers some everyday CLI tools like Kubernetes, GitHub CLI and others.

Considering that it was going to be my first real project in Go, I preferred going with the standard library to 
get more familiar with the Go syntax and how things work within the language. 

#### Handling the input
One of the questions I had before starting the project was "How am I going to handle the user input?"

| Value |  Value + System   |
|:-----:|:-----------------:|
|  50   | 50c \| 50f \| 50k |

The first idea was two separate inputs, one for the value and the other for the temperature system.
I didn't want to prompt the user twice, instead, I used some of the string methods, alongside conditional logic
in Go to work with a single input, parse it and output the result to the terminal.


#### Working with functions
Early during the development of the application, I wrote functions with the different conversion formulas.
This helped me organize my code a bit better and have `main()` more clean.

#### Writing unit tests
Testing is a topic I always wanted to dig in, and after reading about it, 
I learned that writing unit tests in Go is relatively straightforward.
1. The file needs to be named like `xxx_test.go`.
2. The test function must start with the word `Test`.
3. The test function takes one argument only `t *testing.T`.
4. To use the `t *testing.T` we need to import `"testing"` like we did with `"fmt"`.