# Golang Web

## 1. Handling Request

### 1.1 Http library

- notis first
  - lebih fokus kapabilitas http library menangani server, instead of client.

Library bawaan golang, http library di bagi menjadi 2 bagian

- _Client_ -- Client, Response, Header, Request, Cookie
- _Server_ -- Server, ServeMux, Handler/HandlerFunc, ResponseWriter, Header, Request, Cookie.

[Gambar_3.1]

### 1.2 Serving Go

http library punya kemampuan untuk starting up HTTP server yang ngehandle request dan ngirim response ke request tersebut. Juga nyedian interface untuk multiplexer dan default multiflexer (gatau itu apaan :v).

[Gambar_3.2]

### 1.2.1 The Go Web Server

Bikin server itu trivial di golang, cukup make function `ListenAndServe`.

[Code 3.1] The simplest web server

```golang
	type Server struct {
		Addr           string
		Handler        Handler
		ReadTimeOut    time.Duration
		WriteTimeOut   time.Duration
		MaxHeaderBytes int
		TLSConfig      *tls.Config
		TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
		ConnState      func(net.Conn, ConnState)
		ErrorLog       *log.Logger
	}
```

### 1.2.2 Serving through HTTPS

HTTPS ngak lebih dari layering HTTP di atas SSL (Transport Security Layer [ TLS ]). Untuk ngeserve web aplikasi kita lewat HTTPS, kita akan gunakan function`ListenAndServe`.

[code 3.4] Serving througst HTTP

## 1.3. Handlers and handler functions

Nyalain server kek tadi, ngak bikin web aplikasi kita guna. Default multiplexer akan digunakan jika paramater handler adalah nil ngak dapat nemuin handler apapun akan ngeresponse 404.

### 1.3.1 Handling Request

Handler adalah sebuah interface yang memiliki method bernama `ServeHTTP` dengan 2 parameter, yaitu:

- Interface `HTTPResponseWriter`
- Pointer ke Struct Request

```golang
ServeHTTP(http.ResponseWriter, *http.Request)
```

[] code 3.6

handler yang ada di code di atas, ngehandle semua request akibatnya response selalu "Hello World!"

### 1.3.2 More Handlers

Karena kita ngak pengen 1 handler ngehandle semua request, instead kita pengen di setiap handler berbeda untuk URL yang berbeda.

Untuk melakukan itu, kita ngak specify Handler field di `Server` struct (yg mana ini akan menggunakan `DefaultServeMux` sebagai handler); kita akan make `http.Handle` function untuk attach sebuah handler ke `DefaultServeMux`.

Beriku kita bikin 2 handlers dan kemudian attach handler ke URL masing-masing. Sekarang jika pergi ke http://localhost:8080/hello maka dapet "Hello!", jika pergi ke http://localhost:8080/world, dapet "World!"

[code 3.7]

### 1.3.3 Handler Functions

Sebelumnya kita ngak make handler functions, ketika bikin handler. Jadi handler function adalah function yang berprilaku kek handler.

[code 3.8]

source code untuk handler functions

```go
// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}
```

### 1.3.4 Chaining handlers and handler functions

[gambar 3.3]
Misal

## 1.3 Summary

- Golang punya library standar untuk bikin web application, yaitu new/http dan html/template
- Meskipun lebih gampang dan cepat kalo make web frameworks di luar golang, kek gin, gorilla dsb.
- handlers di Go bisa struct apa aja yang memiliki method bernama `ServeHTTP` dengan 2 parameter, yaitu `HTTPResponseWriter` interface dan pointer ke sebuah `Request` struct.
- handler function adalah function yang berprilaku seperti handler. Fungsi handler memiliki signature yang sama dengan method `ServeHTTP` dan digunakan untuk memproses request.
- Multiplexer adalah handler juga. `ServerMux` adalah sebuah HTTP request multiplexer. Ini menerima HTTP request dan redirect ke handler yg sesuai dengan URL request. `DefaultServeMux` adalah instance ServeMux yang tersedia untuk umum yang digunakan sebagai multiplexer default.
