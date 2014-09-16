package mime

import (
	"database/sql/driver"
	"fmt"
	"mime"
	"strings"
)

type Type string

const (
	TypeJPEG Type = "image/jpeg"
	TypePNG  Type = "image/png"
	TypeGIF  Type = "image/gif"
	TypeBMP  Type = "image/x-ms-bmp"

	// raw photos
	TypeRawCanonCR2   = "image/x-canon-cr2"
	TypeRawNikonNEF   = "image/x-nikon-nef"
	TypeRawRW2        = "image/x-raw"
	TypeRawOlympusORF = "image/x-olympus-orf"
	TypeRawSonyRaw    = "image/x-sony-raw"

	// video
	Type3GPP              Type = "video/3gpp"
	TypeAVI               Type = "video/avi"
	TypeFlashVideo        Type = "video/x-flv"
	TypeMatroska          Type = "video/x-matroska"
	TypeMP4               Type = "video/mp4"
	TypeMPEG              Type = "video/mpeg"
	TypeMPEG2TS           Type = "video/mp2t"
	TypeOGG               Type = "video/ogg"
	TypeQuickTime         Type = "video/quicktime"
	TypeWebM              Type = "video/webm"
	TypeWindowsMediaVideo Type = "video/x-ms-wmv"
)

var (
	ext2mime map[string]Type = make(map[string]Type)
	mime2ext map[Type]string = make(map[Type]string)
)

func registerMimetypeExt(mimeType Type, ext string) {
	ext2mime[ext] = mimeType
	mime2ext[mimeType] = ext
	mime.AddExtensionType("."+ext, string(mimeType))
}

func init() {
	registerMimetypeExt(TypeJPEG, "jpeg")
	registerMimetypeExt(TypeJPEG, "jpg")
	registerMimetypeExt(TypePNG, "png")
	registerMimetypeExt(TypeGIF, "gif")
	registerMimetypeExt(TypeBMP, "bmp")

	registerMimetypeExt(TypeRawCanonCR2, "cr2")
	registerMimetypeExt(TypeRawNikonNEF, "nef")
	registerMimetypeExt(TypeRawRW2, "rw2")
	registerMimetypeExt(TypeRawOlympusORF, "orf")
	registerMimetypeExt(TypeRawSonyRaw, "arf")

	registerMimetypeExt(Type3GPP, "3gp")
	registerMimetypeExt(TypeAVI, "avi")
	registerMimetypeExt(TypeFlashVideo, "flv")
	registerMimetypeExt(TypeMatroska, "mkv")
	registerMimetypeExt(TypeMP4, "m4v")
	registerMimetypeExt(TypeMP4, "mp4")
	registerMimetypeExt(TypeMPEG, "mpeg")
	registerMimetypeExt(TypeMPEG, "mpg")
	registerMimetypeExt(TypeMPEG2TS, "m2ts")
	registerMimetypeExt(TypeMPEG2TS, "mts")
	registerMimetypeExt(TypeOGG, "ogv")
	registerMimetypeExt(TypeQuickTime, "mov")
	registerMimetypeExt(TypeWebM, "webm")
	registerMimetypeExt(TypeWindowsMediaVideo, "wmv")
}

// Implements the sql.Scanner interface
func (mimeType *Type) Scan(value interface{}) error {
	// github.com/mattn/sqlite3 used to send strings, now it sends []byte
	switch value.(type) {
	case string:
		*mimeType = Type(value.(string))
	case []byte:
		*mimeType = Type(value.([]byte))
	default:
		return fmt.Errorf("Unable to convert %+v to Type", value)
	}

	return nil
}

// Implements the driver.Valuer interface
func (mimeType Type) Value() (driver.Value, error) {
	return string(mimeType), nil
}

// Returns an extension for a mimetype, without a leading "."
func (mimeType Type) DefaultExtension() string {
	return mime2ext[mimeType]
}

// Returns the mimetype for an extension, which must start with "."
func TypeByExtension(ext string) Type {
	return Type(mime.TypeByExtension(ext))
}

// Returns true if ext is one of the hardcoded extensions in this module.
// Does not look at the system's mime db.  Extension can optionally start with
// "." and are not case-sensitive.
func IsKnownExtension(ext string) bool {
	if strings.HasPrefix(ext, ".") {
		ext = ext[1:]
	}
	_, ok := ext2mime[strings.ToLower(ext)]
	return ok
}
