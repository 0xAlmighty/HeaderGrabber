package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
)

const serverToken = "Almighty"

var (
	port string
	ip   string
)

func displayBanner() {
	banner := `
	__  __     ______     ______     _____     ______     ______     ______     ______     ______     ______     ______     ______     ______   
	/\ \_\ \   /\  ___\   /\  __ \   /\  __-.  /\  ___\   /\  == \   /\  ___\   /\  == \   /\  __ \   /\  == \   /\  == \   /\  ___\   /\  == \  
	\ \  __ \  \ \  __\   \ \  __ \  \ \ \/\ \ \ \  __\   \ \  __<   \ \ \__ \  \ \  __<   \ \  __ \  \ \  __<   \ \  __<   \ \  __\   \ \  __<  
	 \ \_\ \_\  \ \_____\  \ \_\ \_\  \ \____-  \ \_____\  \ \_\ \_\  \ \_____\  \ \_\ \_\  \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\  \ \_\ \_\
	  \/_/\/_/   \/_____/   \/_/\/_/   \/____/   \/_____/   \/_/ /_/   \/_____/   \/_/ /_/   \/_/\/_/   \/_____/   \/_____/   \/_____/   \/_/ /_/
	   

 	HeaderGrabber - Simple Server to analyze requests
 	Version: 1.0.0
	`
	color.Magenta(banner)
}

func displayHelp() {
	color.Yellow("Default value if running without options: 127.0.0.1:8080\n")
	color.Yellow("Usage: headergrabber [OPTIONS]\n")
	color.Yellow("  -ip <ip>   	Specify the address to listen on")
	color.Yellow("  -port <port>    Specify the port the server will listen on")
	color.Yellow("  -h, --help      Display this help message and exit")
}

func init() {
	flag.StringVar(&port, "port", "8080", "Specify the port the server will listen on")
	flag.StringVar(&ip, "ip", "127.0.0.1", "Specify the IP address the server will listen on")
}

func printCookies(r *http.Request) {
	for _, cookie := range r.Cookies() {
		color.New(color.FgYellow).Printf("[*] Cookies:\n%s=%s\n", cookie.Name, cookie.Value)
		if cookie.Name == "session_cookie_name" {
			color.New(color.FgGreen).Printf("[*] Found sessioncookie: %s. Attempting actions...\n", cookie.Name)
		}
	}
	fmt.Println()
}

func printHeaders(headers http.Header) {
	color.New(color.FgCyan).Println("\nHeaders:")
	for name, values := range headers {
		for _, value := range values {
			color.New(color.FgYellow).Printf("  %s: %s\n", name, value)
		}
	}
	fmt.Println()
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", serverToken)
	w.Header().Set("Content-Type", "text/html")

	color.New(color.FgHiMagenta).Printf("\n[*] %s Request:\n", r.Method)
	color.New(color.FgGreen).Printf("\n[*] Full URL: %s\n", r.URL.String())

	printHeaders(r.Header)
	printCookies(r)

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			color.New(color.FgRed).Printf("ParseForm() err: %v\n", err)
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		color.New(color.FgYellow).Printf("Params:\n%v\n", r.PostForm)
	}
}

func main() {
	displayBanner()

	helpFlag := flag.Bool("help", false, "Display this help message and exit")
	flag.Parse()

	if *helpFlag {
		displayHelp()
		return
	}

	address := fmt.Sprintf("%s:%s", ip, port)

	server := &http.Server{
		Addr:    address,
		Handler: nil,
	}

	http.HandleFunc("/", handleRequests)

	color.New(color.FgHiGreen).Printf("[*] Starting server on %s...\n", address)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	color.New(color.FgRed, color.Bold).Println("\n\n[*] Server stopping...")
}
