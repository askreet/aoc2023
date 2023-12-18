package advent

import (
	"bytes"
)

// Scanner function for scanning puzzles where input is large sections of text divided by a full blank line.
func ScanSections(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}

func ScanCommaSeperated(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if idx := bytes.IndexByte(data, ','); idx != -1 {
		return idx + 1, data[0:idx], nil
	} else if atEOF {
		// No more data to load, the rest is a single word.
		return len(data), data, nil
	} else {
		// We need more data to be sure.
		return 0, nil, nil
	}
}
