package lists

// Stop is a list of reserved words, data types and Go library names.
var Stop []string

func init() {
	Stop = []string{
		"bool",
		"break",
		"byte",
		"case",
		"chan",
		"complex128",
		"complex64",
		"const",
		"continue",
		"default",
		"defer",
		"else",
		"fallthrough",
		"float32",
		"float64",
		"for",
		"func",
		"go",
		"goto",
		"if",
		"import",
		"int",
		"int16",
		"int32",
		"int64",
		"int8",
		"interface",
		"map",
		"package",
		"range",
		"return",
		"rune",
		"select",
		"string",
		"struct",
		"switch",
		"type",
		"uint",
		"uint16",
		"uint32",
		"uint64",
		"uint8",
		"uintptr",
		"var",
		"archive",
		"tar",
		"zip",
		"bufio",
		"builtin",
		"bytes",
		"compress",
		"bzip2",
		"flate",
		"gzip",
		"lzw",
		"zlib",
		"container",
		"heap",
		"list",
		"ring",
		"context",
		"crypto",
		"aes",
		"cipher",
		"des",
		"dsa",
		"ecdsa",
		"elliptic",
		"hmac",
		"md5",
		"rand",
		"rc4",
		"rsa",
		"sha1",
		"sha256",
		"sha512",
		"subtle",
		"tls",
		"x509",
		"pkix",
		"database",
		"sql",
		"driver",
		"debug",
		"dwarf",
		"elf",
		"gosym",
		"macho",
		"pe",
		"plan9obj",
		"encoding",
		"ascii85",
		"asn1",
		"base32",
		"base64",
		"binary",
		"csv",
		"gob",
		"hex",
		"json",
		"pem",
		"xml",
		"errors",
		"expvar",
		"flag",
		"fmt",
		"go",
		"ast",
		"build",
		"constant",
		"doc",
		"format",
		"importer",
		"parser",
		"printer",
		"scanner",
		"token",
		"types",
		"hash",
		"adler32",
		"crc32",
		"crc64",
		"fnv",
		"html",
		"template",
		"image",
		"color",
		"palette",
		"draw",
		"gif",
		"jpeg",
		"png",
		"index",
		"suffixarray",
		"io",
		"ioutil",
		"log",
		"syslog",
		"math",
		"big",
		"bits",
		"cmplx",
		"rand",
		"mime",
		"multipart",
		"quotedprintable",
		"net",
		"http",
		"cgi",
		"cookiejar",
		"fcgi",
		"httptest",
		"httptrace",
		"httputil",
		"pprof",
		"mail",
		"rpc",
		"jsonrpc",
		"smtp",
		"textproto",
		"url",
		"os",
		"exec",
		"signal",
		"user",
		"path",
		"filepath",
		"plugin",
		"reflect",
		"regexp",
		"syntax",
		"runtime",
		"cgo",
		"debug",
		"msan",
		"pprof",
		"race",
		"trace",
		"sort",
		"strconv",
		"strings",
		"sync",
		"atomic",
		"syscall",
		"js",
		"testing",
		"iotest",
		"quick",
		"text",
		"scanner",
		"tabwriter",
		"template",
		"parse",
		"time",
		"unicode",
		"utf16",
		"utf8",
		"unsafe"}
}
