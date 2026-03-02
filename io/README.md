

# Notes

IO in Go is built around two core interfaces:

type Reader interface {
Read(p []byte) (n int, err error)
}

type Writer interface {
Write(p []byte) (n int, err error)
}

Everything else builds on top of these.


Yes. Here is a short, README-friendly version.

---

## IO in Go – When Do You Use It?

IO (Input/Output) is used whenever your program moves data across boundaries:

* HTTP requests and responses (`r.Body`, `http.ResponseWriter`)
* Calling external services (`resp.Body`)
* Reading and writing files (`os.Open`, `os.Create`)
* Streaming large data (`io.Copy`)
* Logging to multiple outputs (`io.MultiWriter`)
* In-memory buffers for testing (`bytes.Buffer`)

Core idea:

```text
Reader → processing → Writer
```

Everything in Go that transfers data (network, files, compression, proxies, ETL pipelines) is built on `io.Reader` and `io.Writer`.

If your program talks to disk, network, or another service, you are using IO.
