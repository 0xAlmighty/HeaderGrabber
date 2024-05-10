# 🛡️ HeaderGrabber

HeaderGrabber is a powerful, simple, Go-based tool designed to provide real-time analysis of HTTP requests. It's built to capture requests information. 🚀

## Features 🌟

- **Real-time HTTP Request Analysis:** Instantly captures and displays HTTP headers, cookies, and HTTP method data.
- **IP and Port Customization:** Tailor the server listening address and port.
- **Session Cookie Detection:** Highlights session cookies for further exploration and testing, very useful for grabbing those cookies from XSS.
- **Ease of Use:** Simple setup and user-friendly.

## Getting Started 🚀

### Prerequisites

Just need Go installed on your system.

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

Want to contribute? Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b)
3. Commit your Changes (`git commit -m 'adding new'`)
4. Push to the Branch (`git push origin new`)
5. Open a Pull Request

## License 📜

Distributed under the MIT License. See `LICENSE` for more information.

Made with ❤️ and ☕ by [almighysec](https://almightysec.com/). 2024.
