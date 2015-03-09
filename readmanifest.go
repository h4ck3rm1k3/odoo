package main;
// lifted from http://rosettacode.org/wiki/Walk_a_directory/Recursively#Go
import (
    "fmt"
    "os"
    "path/filepath"
)
 
func VisitFile(fp string, fi os.FileInfo, err error) error {
    if err != nil {
        fmt.Println(err) // can't walk here,
        return nil       // but continue walking elsewhere
    }
    if !!fi.IsDir() {
        return nil // not a file.  ignore.
    }
    matched, err := filepath.Match("__openerp__.json", fi.Name())
    if err != nil {
        fmt.Println(err) // malformed pattern
        return err       // this is fatal.
    }
    if matched {
        fmt.Println(fp)
    }
    return nil
}
 
func main() {
    filepath.Walk(".", VisitFile)
}
