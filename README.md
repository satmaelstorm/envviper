# EnvViper
[![Go Report](https://goreportcard.com/badge/github.com/satmaelstorm/envviper)](https://goreportcard.com/report/github.com/satmaelstorm/envviper) 
[![GoDoc](https://godoc.org/github.com/satmaelstorm/envviper?status.svg)](http://godoc.org/github.com/satmaelstorm/envviper)
[![Coverage Status](https://coveralls.io/repos/github/satmaelstorm/envviper/badge.svg?branch=master)](https://coveralls.io/github/satmaelstorm/envviper?branch=master) 
![Go](https://github.com/satmaelstorm/envviper/workflows/Go/badge.svg)

[Viper](https://github.com/spf13/viper) package doesn't consider environment variables while unmarshaling.
See: [188](https://github.com/spf13/viper/issues/188) and [761](https://github.com/spf13/viper/issues/761)

EnvViper is a wrapper for [viper](http://github.com/spf13/viper) with the same API, and resolve this issues.
