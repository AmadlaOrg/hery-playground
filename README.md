<img src=".assets/playground.jpg" alt="Electronics photo" style="width: 400px;" align="right">

# Hery - Playground 🛝
🛝 Hery - Playground 🛝

Simple single page web application to try out EQL (Entity Query Language) against a HERY caching file or just a `.hery` 
file.

It is possible to drag & drop a [YAML](https://yaml.org/)/`.hery` file and/or a [SQLite3](https://www.sqlite.org/) db file (with the HERY caching structure).
So to be able to run EQL queries against.

There is also support for [JSON-Schema](https://json-schema.org/) to validate a `.hery` file.

## How it works?
It runs 100% on the frontend. And no it doesn't use any of the overly complicated JS frameworks like React and Angular 🤮.
To get a more deep understand of how it works: [.docs/dev.md](.docs/dev.md).

There are two side to the UX one is the query editor and the other the data editor where you find a code editor and 
through the tabs the DB explorer if a [SQLite3](https://www.sqlite.org/) db file was drop in.

It is possible to upload or drag & drop a [YAML](https://yaml.org/)/`.hery` file or/and a [SQLite3](https://www.sqlite.org/) db file into the playground
and everything will automatically load into the editor without a page reload.

## Why?
It is possible to try out EQL queries with HERY but the terminal environment is not as fun, intuitive and easy to use than
a playground. Same goes for IDEs.

## Production? Staging? Dev?
Yes it can be used in different environments where it can interact via an HTTP server with HERY. That said it would
be wise to make sure the public does not have access to it for security reasons.

## Requirements
These are main browser supports that are required: 
- [wasm](https://caniuse.com/wasm)
- [FileAPI](https://caniuse.com/fileapi)

There is no need for: [Web SQL Database](https://caniuse.com/sql-storage) since [sql.js](https://github.com/sql-js/sql.js/)
does not require it to work with [SQLite3](https://www.sqlite.org/) db files.

> Most modern browsers from 2017 and up should work with the HERY - Playground. Even mobile web browsers should also work.
> That said the UX might be awkward to use on a smaller screen. 

## ©️ Copyright
- "<a rel="noopener noreferrer" href="https://www.rawpixel.com/image/6512120/image-vintage-public-domain-kid">Boy playing toy drawing, vintage</a>" is marked with <a rel="noopener noreferrer" href="https://creativecommons.org/publicdomain/zero/1.0/?ref=openverse">CC0 1.0 <img src="https://mirrors.creativecommons.org/presskit/icons/cc.svg" style="height: 1em; margin-right: 0.125em; display: inline;" /><img src="https://mirrors.creativecommons.org/presskit/icons/zero.svg" style="height: 1em; margin-right: 0.125em; display: inline;" /></a>.

## :scroll: License

The license for the code and documentation can be found in the [LICENSE](./LICENSE) file.

---

Made in Québec 🏴󠁣󠁡󠁱󠁣󠁿, Canada 🇨🇦!