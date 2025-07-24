# learn-lang-go
Learn | Language Programming | Go.


# Problems

🧮 Modulo Fibonacci Sequence Server
• Designed to generate Fibonacci numbers modulo 10910^9109.
• Implemented a rate limiter to ensure responses are sent no earlier than 10 milliseconds apart.
• Handled concurrent requests using Go channels to synchronize data flow.

🛜 TCP Server for Reversing Strings
• Built a TCP server that listens on a specified address.
• Employed goroutines to handle multiple simultaneous connections efficiently.
• Implemented string reversal by converting strings to rune slices to properly handle Unicode characters.

💾 HTTP Server for Managing Lake Data
• Developed RESTful endpoints to add, retrieve, and delete lake data.
• Utilized Go's net/http package for handling HTTP requests and responses.
• Ensured robust JSON encoding/decoding and appropriate HTTP status codes for various scenarios.
