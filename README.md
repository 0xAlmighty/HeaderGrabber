# 🛡️ HeaderGrabber

HeaderGrabber is a powerful, yet simple, Go-based tool designed to provide real-time analysis of HTTP requests. With its vibrant CLI interface and comprehensive request details, it's an indispensable tool for cybersecurity professionals and enthusiasts alike. 🚀

## Features 🌟

- **Real-time HTTP Request Analysis:** Instantly captures and displays HTTP headers, cookies, and POST data.
- **IP and Port Customization:** Tailor the server listening address to meet your specific testing needs.
- **Vibrant CLI Visualization:** Utilizes the Fatih Color library for an engaging and colorful display of request details.
- **Session Cookie Detection:** Highlights session cookies for further exploration and testing, very useful grabbing those cookies from XSS.
- **Ease of Use:** Simple setup and user-friendly commands make it accessible for all skill levels.

## Getting Started 🚀

### Prerequisites

Ensure you have Go installed on your system. HeaderGrabber is built with Go, so it's necessary for running the tool.

### Installation

Grab the release version from the release page

**Or**

Clone the repository to your local machine:

```
git clone https://github.com/0xAlmighty/HeaderGrabber.git
```

Navigate into the project directory:
```
cd HeaderGrabber
```

Build the tool:
```
go build
```

Copy to `PATH` for easier access :)

### Usage

To start HeaderGrabber with default settings (listening on `127.0.0.1:8080`), simply run:

```
./headergrabber
```

To customize the IP address and port, use the `-ip` and `-port` flags:

```
./headergrabber -ip IP_ADDRESS -port PORT
```

For help and more commands:
```
./headergrabber -help
```

## Contributing 🤝

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b)
3. Commit your Changes (`git commit -m 'adding new'`)
4. Push to the Branch (`git push origin new`)
5. Open a Pull Request

## License 📜

Distributed under the MIT License. See `LICENSE` for more information.

Made with ❤️ and ☕ by [almighysec](https://almightysec.com/). 2024.
