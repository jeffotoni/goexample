package main

import (
    "fmt"
)

var routes map[string]func() string

func main() {
    routes = map[string]func() string{
        "GET /":      homePage,
        "GET /about": aboutPage,
        "POST /user": userPage,
    }

    fmt.Println("GET /", pageContent("GET /"))
    fmt.Println("GET /about", pageContent("GET /about"))
    fmt.Println("GET /unknown", pageContent("GET /unknown"))
    fmt.Println("POST /user", pageContent("POST /user"))
    // Output:
    // GET / Home page
    // GET /about About page
    // GET /unknown 404: Page Not Found
    // POST /user Page User
}

func pageContent(route string) string {
    page, ok := routes[route]
    if ok {
        return page()
    } else {
        return notFoundPage()
    }
}

func homePage() string {
    return "Home page"
}

func aboutPage() string {
    return "About page"
}

func userPage() string {
    return "Page User"
}

func notFoundPage() string {
    return "404: Page Not Found"
}
