# FileSeeker

[![Go Reference](https://pkg.go.dev/badge/github.com/singoesdeep/fileseeker)](https://pkg.go.dev/github.com/singoesdeep/fileseeker)

FileSeeker is a Go package that provides a powerful and flexible way to search for files within a given folder. With FileSeeker, you can easily retrieve a list of file paths based on specific criteria such as file patterns, extensions, and the option to include or exclude subdirectories.

FileSeeker is designed to be lightweight, efficient, and easy to use. Whether you need to scan a directory for specific file types or perform more advanced file filtering, FileSeeker has got you covered.

To get started, check out the documentation and examples in the repository. Contributions and feedback are welcome, so feel free to open issues or submit pull requests if you have any ideas or improvements.

Give FileSeeker a try and simplify your file searching tasks in Go!

## TODO

- [x] Implement regular expression pattern matching for file filtering
- [x] Add option to exclude subdirectories from seeking
- [ ] Add option to exclude specific files/subdirs
- [ ] Add more examples demonstrating different usage scenarios
- [ ] Provide benchmark tests and performance optimizations
- [ ] Add support for custom filters and sorting options
- [ ] Include more comprehensive documentation and usage guidelines
- [ ] Implement additional error handling and robustness checks
- [ ] Support parallel file seeking for improved performance
- [ ] Add support for file metadata extraction (e.g., file size, modification time)

Feel free to contribute by tackling any of the above tasks. If you have any other ideas or suggestions, please open an issue or reach out with your thoughts. Your contributions are greatly appreciated!



## Installation

Use the following `go get` command to install the package:

```bash
go get github.com/singoesdeep/fileseeker
```

## Usage

Here's an example demonstrating how to use the FileSeeker package:

```go
package main

import (
	"fmt"

	"github.com/singoesdeep/fileseeker"
)

func main() {
	// Create a FileSeekerBuilder
	builder := fileseeker.NewFileSeekerBuilder("/path/to/folder").
		Patterns([]string{"*.txt"}). //you can find files with regexp
		FileExtensions([]string{"jpg", "png"}). //or extension based and if you use both it will find both
		ExcludeSubdirs() //dont check subdirs, default is true so dont use it if you want check subdirs

	// Build the FileSeeker
	seeker := builder.Build()

	// Seek files
	files, err := seeker.SeekFiles()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the found files
	for _, file := range files {
		fmt.Println(file)
	}
}
```

## Features

- Seek files within a specified folder
- Filter files by patterns and extensions
- Include or exclude subdirectories
- Supports regular expression pattern matching

## Documentation

The package documentation can be found at: [pkg.go.dev/github.com/singoesdeep/fileseeker](https://pkg.go.dev/github.com/singoesdeep/fileseeker)

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).