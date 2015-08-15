package main

import (
    "bufio"
    "os"
    "strings"
    "unicode"
    "regexp"
    "fmt"
    "crypto/rand"
    "math/big"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        sanitized, err := sanitize(scanner.Text())
        if err != nil {
            fmt.Println(err)
            continue
        }
        tld, err := randomTLD()
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Println(sanitized + "." + tld)
    }
}

func sanitize(str string) (string, error) {
    if matched, _ := regexp.MatchString("^[\\w\\s_-]+$", str); !matched {
        return "", fmt.Errorf("Included invalid characters")
    }
    lower := strings.ToLower(str)
    hyphened := []rune{}
    for _, r := range lower {
        if unicode.IsSpace(r) {
            r = '-'
        }
        hyphened = append(hyphened, r)
    }
    return string(hyphened), nil
}

func randomTLD() (string, error) {
    tlds := []string{"com", "org", "net"}
    big, err := rand.Int(rand.Reader, big.NewInt(int64(len(tlds))))
    if err != nil {
        return "", err
    }
    return tlds[big.Int64()], nil
}
