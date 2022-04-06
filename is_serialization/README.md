## WTF is Serialization?
Converting an in-memory data structure to a value that can be stored or transferred.
When a "LINK:runtime" builds a data structure, engine store each field of the data strucutre in memory locations and provide references
This data sturcture is easy to ingest in the same runtime context, but how would another machine process this after the process ends?
Conversion of the in-memory data structe into series of bytes that record the data strucre in a recoverable format
Deserialization is the other side of the coin, we have a series of bytes that then needs to be recovered to an in memory data structure.
At one runtime context the golang struct is known in the process memory and then in the browser context it can know the same.

### Simple Binary Encoding:
- codec aimed at addressing the efficiency inssues in low-latency trading
- SBE OSI 6 presentation for encoding decoding message in bianry format

### CapnP:
- mmapable zero-copy serialization formats


JavaScript: JSON strings and JSON parse, ascii encodings(slow)
Protobufs: 
CapnP:
