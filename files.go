package GoEmbed

import "embed"

// FileText don't use space in front go:embed
//
//go:embed files/version.txt
var FileText string

//go:embed files/haerin.png
var FileImage []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var FileMultiple embed.FS // embed.FS is a type of io/fs.FS

//go:embed files/*.txt
var FilePath embed.FS
