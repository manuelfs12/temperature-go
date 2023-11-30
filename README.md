# Temperature GO
### Video Demo:  <URL HERE>
### Description:
This is a temperature conversion CLI tool written in Go.

### What is the project?
This is a temperature conversion tool written in Go for the final project of CS50x.

[//]: # (gif)

The software takes a temperature input, be it in Celsius, Fahrenheit or Kelvin and outputs the converted values as a formatted string.

When program is launched, it opens a terminal window where you can write a temperature, then it takes that input, validates it, and converts the temperature into the other scales.

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


* **LICENSE**: The license file of the repository, in this case is MIT license.


* **main.go**: This is the body of the application, where all the coding logic is in. I import some packages from the standard library, 
declare the `main()` function and all the other functions that makes the application work.

### Debated design choices? Using a framework vs standard library, Go vs other languages, etc.

#### Go vs Another language
Since before starting CS50x, I was toying with the idea of learning Go, but got scared when I read that it was
a combination of Python and C.
After going through CS50x, and dipping my toes in C, it motivated me to work with Go, since I like static type languages.

I have previous experience working with languages like Python, Javascript and C#, But for this project I wanted to challenge
myself and create something in a new language.

#### Framework vs Standard Library
When I decided to use Go as my language for this project, I considered using a framework called [Cobra](https://github.com/spf13/cobra).
This is the recommended framework on Go website for creating modern and powerful CLI applications, and
it's the framework that empowers some everyday CLI tools like Kubernetes, GitHub CLI and others.

Since this was going to be my first real project in Go, I preferred going with the standard library to get more familiar with the Go syntax and how things work within the language. 

#### Handling the input
One of the questions I had before starting the project was "How am I going to handle the user input?"

| Value | Value + System |
|:-----:|:--------------:|
|  50   |   50c or 50f   |

The first idea was two separate inputs, one for the value and the other for the temperature system.
I didn't want to prompt the user twice, instead, I used some of the string methods, alongside conditional logic
in Go to work with a single input, parse it and output the result to the terminal.


#### Working with functions
Early during the development of the application, I wrote functions with the different conversion formulas.
This helped me organize my code a bit better and have `main()` more clean.