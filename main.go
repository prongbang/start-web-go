package main

import ("fmt" ; "net/http" ; "io")

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html",)
	io.WriteString(res,`
		<doctype html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>Hello World!</body>
		</html>`,
	)
}

func wellcome(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html",)
	io.WriteString(res,`
		<doctype html>
		<html>
			<head>
				<title>Wellcome</title>
			</head>
			<body>Wellcome</body>
		</html>`,
	)
}

func main() {

	fmt.Println("Started port : 9000")

	http.HandleFunc("/", wellcome)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
	
	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer( http.Dir("assets") ),),
	)
	
}