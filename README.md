# Slinky

[![Go Reference](https://pkg.go.dev/badge/github.com/FurqanSoftware/slinky.svg)](https://pkg.go.dev/github.com/FurqanSoftware/slinky) 

Parse social media URLs in Go.

## Usage

``` go
slinky.Parse("https://github.com/hjr265")
// Output:
// 	&URL{
// 		Service: slinky.GitHub,
// 		Type:    "User",
// 		ID:      "hjr265",
// 		Data:    map[string]string{,
// 			"username": "hjr265",
// 		},
//	}
```

## Contributing

Contributions are welcome.

## License

Slinky is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).
