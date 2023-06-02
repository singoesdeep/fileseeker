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
- [ ] Support parallel file seeking for improved performance
- [x] Add support for file metadata extraction (e.g., file size, modification time)

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
		Patterns([]string{"^*.txt", "^*.jpg"}). //you can find files with regexp
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

### Will Return
```
{file.go file.go go 1912 2023-06-02 18:53:02.7379161 +0300 +03 -rw-rw-rw-}
{fileseeker.go fileseeker.go go 2084 2023-06-02 17:53:54.9805177 +0300 +03 -rw-rw-rw-}
{fileseeker_builder.go fileseeker_builder.go go 1896 2023-06-02 18:38:45.5054916 +0300 +03 -rw-rw-rw-}
{fileseeker_builder_test.go fileseeker_builder_test.go go 2459 2023-06-02 15:49:08.48243 +0300 +03 -rw-rw-rw-}
{fileseeker_config.go fileseeker_config.go go 244 2023-06-02 17:50:07.9253496 +0300 +03 -rw-rw-rw-}
{fileseeker_test.go fileseeker_test.go go 2190 2023-06-02 15:48:26.9262423 +0300 +03 -rw-rw-rw-}
```
---------------------------------------
---------------------------------------
```go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/singoesdeep/fileseeker"
)

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("#", "Path", "Name", "Ext", "Size", "ModDate", "Perms")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Create a FileSeekerBuilder
	builder := fileseeker.NewFileSeekerBuilder("./").
		ExcludeSubdirs().                                       //pass subdirs
		Patterns([]string{"^*.go", "^go.mod"}).                 //regexp pattern matching
		SizeRangeFilter([2]int64{1, fileseeker.MAX_FILE_SIZE}). //min and max bytes for filtering
		ModificationDateRangeFilter([2]time.Time{fileseeker.MIN_DATE_TIME, fileseeker.MAX_DATE_TIME})

	// Build the FileSeeker
	seeker := builder.Build()

	// Seek files
	files, err := seeker.SeekFiles()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the found files
	for i, file := range files {
		tbl.AddRow(
			strconv.Itoa(i+1),
			file.String()[0],
			file.String()[1],
			file.String()[2],
			file.String()[3],
			file.String()[4],
			file.String()[5],
		)
		fmt.Println()
	}

	tbl.Print()
}
```
### Will Return
| # | Path                        | Name                        | Ext | Size | ModDate              | Perms     |
|---|-----------------------------|-----------------------------|-----|------|----------------------|-----------|
| 1 | file.go                     | file.go                     | go  | 1912 | 2023-06-02 18:53:02  | rw-rw-rw- |
| 2 | fileseeker.go               | fileseeker.go               | go  | 2084 | 2023-06-02 17:53:54  | rw-rw-rw- |
| 3 | fileseeker_builder.go       | fileseeker_builder.go       | go  | 1896 | 2023-06-02 18:38:45  | rw-rw-rw- |
| 4 | fileseeker_builder_test.go  | fileseeker_builder_test.go  | go  | 2459 | 2023-06-02 15:49:08  | rw-rw-rw- |
| 5 | fileseeker_config.go        | fileseeker_config.go        | go  | 244  | 2023-06-02 17:50:07  | rw-rw-rw- |
| 6 | fileseeker_test.go          | fileseeker_test.go          | go  | 2190 | 2023-06-02 15:48:26  | rw-rw-rw- |
| 7 | go.mod                      | go.mod                      | mod | 288  | 2023-06-02 18:44:12  | rw-rw-rw- |


## Features

- Seek files within a specified folder
- Include or exclude subdirectories
- Supports regular expression pattern matching
- Supports file size range filtering
- Supports modification date range filtering

## Documentation

The package documentation can be found at: [pkg.go.dev/github.com/singoesdeep/fileseeker](https://pkg.go.dev/github.com/singoesdeep/fileseeker)

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE).
