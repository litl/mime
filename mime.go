package mime

import (
	"database/sql/driver"
	"fmt"
	"mime"
)

func init() {
	mime.AddExtensionType(".jpeg", string(TypeJPEG))
	mime.AddExtensionType(".jpg", string(TypeJPEG))
	mime.AddExtensionType(".png", string(TypePNG))

	mime.AddExtensionType(".cr2", string(TypeRawCanonCR2))
	mime.AddExtensionType(".nef", string(TypeRawNikonNEF))
	mime.AddExtensionType(".rw2", string(TypeRawRW2))
	mime.AddExtensionType(".orf", string(TypeRawOlympusORF))
	mime.AddExtensionType(".arf", string(TypeRawSonyRaw))

	mime.AddExtensionType(".3gp", string(Type3GPP))
	mime.AddExtensionType(".avi", string(TypeAVI))
	mime.AddExtensionType(".flv", string(TypeFlashVideo))
	mime.AddExtensionType(".mkv", string(TypeMatroska))
	mime.AddExtensionType(".m4v", string(TypeMP4))
	mime.AddExtensionType(".mp4", string(TypeMP4))
	mime.AddExtensionType(".mpeg", string(TypeMPEG))
	mime.AddExtensionType(".mpg", string(TypeMPEG))
	mime.AddExtensionType(".m2ts", string(TypeMPEG2TS))
	mime.AddExtensionType(".mts", string(TypeMPEG2TS))
	mime.AddExtensionType(".ogv", string(TypeOGG))
	mime.AddExtensionType(".mov", string(TypeQuickTime))
	mime.AddExtensionType(".webm", string(TypeWebM))
	mime.AddExtensionType(".wmv", string(TypeWindowsMediaVideo))
}

type Type string

const (
	TypeJPEG Type = "image/jpeg"
	TypePNG  Type = "image/png"

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

// Implements the sql.Scanner interface
func (mimeType *Type) Scan(value interface{}) error {
	// github.com/mattn/sqlite3 used to send strings, now it sends []byte
	switch value.(type) {
	case string:
		*mimeType = Type(value.(string))
	case []byte:
		*mimeType = Type(value.([]byte))
	default:
		return fmt.Errorf("Unable to convert %+v to ServerMediaId", value)
	}

	return nil
}

// Implements the driver.Valuer interface
func (mimeType Type) Value() (driver.Value, error) {
	return string(mimeType), nil
}

func (mimeType Type) DefaultExtension() string {
	switch mimeType {
	case TypeJPEG:
		return "jpeg"
	case TypePNG:
		return "png"

	case TypeRawCanonCR2:
		return "cr2"
	case TypeRawNikonNEF:
		return "nef"
	case TypeRawRW2:
		return "rw2"
	case TypeRawOlympusORF:
		return "orf"
	case TypeRawSonyRaw:
		return "raw"

	case Type3GPP:
		return "3gp"
	case TypeAVI:
		return "avi"
	case TypeFlashVideo:
		return "flv"
	case TypeMatroska:
		return "mkv"
	case TypeMP4:
		return "mp4"
	case TypeMPEG:
		return "mpg"
	case TypeMPEG2TS:
		return "mts"
	case TypeOGG:
		return "ogv"
	case TypeQuickTime:
		return "mov"
	case TypeWebM:
		return "webm"
	case TypeWindowsMediaVideo:
		return "wmv"
	default:
		return ""
	}
}

func TypeByExtension(ext string) Type {
	return Type(mime.TypeByExtension(ext))
}
