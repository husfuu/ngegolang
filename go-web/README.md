# Golang Web

## TABLE OF CONTENTS

- [Handling Request](#1-handling-request)
  - [HTTP library](#11-http-library)
  - [Serving Go](#12-serving-go)
    - [The Go Web Server](#121-the-go-web-server)
    - [Serving through HTTPS](#122-serving-through-https)
  - [Handler and Handler Function](#13-handlers-and-handler-functions)
  - [Handling Request](#131-handling-request)
  - [More Handler](#132-more-handlers)
  - [Handler Function](#133-handler-functions)
- [Summary](#14-summary)

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

## 1.4 Summary

- Golang punya library standar untuk bikin web application, yaitu new/http dan html/template
- Meskipun lebih gampang dan cepat kalo make web frameworks di luar golang, kek gin, gorilla dsb.
- handlers di Go bisa struct apa aja yang memiliki method bernama `ServeHTTP` dengan 2 parameter, yaitu `HTTPResponseWriter` interface dan pointer ke sebuah `Request` struct.
- handler function adalah function yang berprilaku seperti handler. Fungsi handler memiliki signature yang sama dengan method `ServeHTTP` dan digunakan untuk memproses request.
- Multiplexer adalah handler juga. `ServerMux` adalah sebuah HTTP request multiplexer. Ini menerima HTTP request dan redirect ke handler yg sesuai dengan URL request. `DefaultServeMux` adalah instance ServeMux yang tersedia untuk umum yang digunakan sebagai multiplexer default.

## 2. Processing Request

### 2.1 Request and Response

Ada 2 tipe HTTP message, yaitu HTTP request dan HTTP response.
Keduanya punya strutur yang sama:

1. Request or response line
2. Zero or more headers
3. An empty line, followed by ...
4. ... an optional message body

Berikut adalah contoh dari GET request:

```
GET /Protocols/rfc2616/rfc2616.html HTTP/1.1
Host: www.w3.org
User-Agent: Mozilla/5.0
(empty line)
```

Library http menyediakan structur untuk HTTP message ini, selanjutnya kita akan bahas mengenai itu.

#### 2.1.1 Request

`Request` struct merepresentasikan sebuah HTTP request message yang dikirimkan dari client. Berikut adalah bagian penting dari `Request`

- URL
- Header
- Body
- Form, PostForm dan Multipartform

Kita juga bisa mendapatkan akses kepada cookies di request dan URL dan user agent dari method di `Request`.

#### 2.1.2 Request URL

URL field dari sebuah request merupakan representasi dari URL yang dikirim sebagai bagian dari request line (line pertama dari HTTP request). URL field adalah sebuah pointer ke `url.URL` type, yang merupakan struct dengan sejumlah field, seperti yang ditunjukkan di sini.

```golang
type URL struct{
	Scheme 		string
	Opaque 		string
	User 		*Userinfo
	Host 		string
	Path 		string
	RawQuery 	string
	Fragment 	string
}

```

Bentuk umumnya kek gini:
`scheme://[userinfo@]host/path[?query][#fragment]`

URL yang ngak dimulai dengan slash setelah scheme diintepretasikan sebagai

`scheme:opaque[?query][#fragment]`

#### 2.1.3 Request Header

Request dan response header dideskripsikan di Header type, yang merupakan map yang merepresentasikan key-value pair dalam HTTP header.

[Gambar 4.1]

### 2.2.2 Response Writer

`ResponseWriter` adalah interface yang handler gunakan untuk bikin suatu HTTP response.

Interface `ResponseWriter` punya 3 method:

- `Write`
- `WriteHeader`
- `Header`

#### 2.2.3 Writing to the ResponseWriter

`Write` method mengambil array of bytes dan ditulis ke dalam body dari HTTP response.

Code 4.8 Write to send responses to the client

- `WriteHeader` method sangat guna ketika pengen return status/error code. Kalo ngak specify status code make method ini, maka entar return-nya 200 terus.

- Writing headers to redirect the client

  ```golang
  func headerExample(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Location", "http://google.com")
  	w.WriteHeader(302)
  }
  ```

  test code: `curl -i 127.0.0.1:8080/redirect`

- Writing JSON input

  ```golang
  func jsonExample(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	post := &Post{
  		User:
  		"Sau Sheong",
  		Threads: []string{"first", "second", "third"},
  	}
  	json, _ := json.Marshal(post)
  	w.Write(json)
  }
  ```

  test code: `curl -i 127.0.0.1:8080/json`

## Cookies

Cookie adalah small information that stored at the client, originally dikirim dari server lewat HTTP response message. Setiap kali client ngirim request HTTP request ke server, cookie dikirim bersamaan.

There are a number of types of cookies, including interestingly named ones like
super cookies, third-party cookies, and zombie cookies. But generally there are only
two classes of cookies: session cookies and persistent cookies. Most other types of
cookies are variants of the persistent cookies.

### Cookie with Go

`Cookie` struct, adalah repsentasi dari cookie di golang.

```golang
type Cookie struct {
	Name string
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string
	MaxAge int
	Secure bool
	HttpOnly bool
	Raw string
	Unparsed []string
}
```

### Sending cookie ke browser

[code 4.13] Sending cookie to the browser

### Summary

- Go provides a representation of the HTTP requests through various structs, which can be used to extract data from the requests.
- The Go Request struct has three fields, Form, PostForm, and MultipartForm, that allow easy extraction of different data from a request. To get data from these fields, call ParseForm or ParseMultipartForm to parse the request and then access the Form, PostForm, or MultipartForm field accordingly.
- Form is used for URL -encoded data from the URL and HTML form, PostForm is used for URL -encoded data from the HTML form only, and MultipartForm is used for multi-part data from the URL and HTML form.
- To send data back to the client, write header and body data to ResponseWriter.
- To persist data at the client, send cookies in the ResponseWriter.
- Cookies can be used for implementing flash messages.
