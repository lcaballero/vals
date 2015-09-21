# Introduction

`vals` is a simple library that provides an API over a generic
interface in `Go`. The idea is that an instance of `vals.Value` is
created with some kind of backing data.  Typically the backing data is
a `map[string]interface{}`.

## Usage

Below is the typical usage for dealing with maps.

```
m := make(map[string]interface{})
m["city"] = "Boulder"
m["firstName"] = "Bruce"
m["lastName"] = "Wayne"
m["tags"] = []string{ "super", "hero" }

v := vals.New(m)
v.At("city").AsString()
v.At("tags").In(0).AsString()

```

## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
