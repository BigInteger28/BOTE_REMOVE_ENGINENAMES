package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Open het invoerbestand
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Fout bij openen van bestand:", err)
        return
    }
    defer file.Close()

    // Maak een scanner om het bestand te lezen
    scanner := bufio.NewScanner(file)

    // Maak een slice om de codes op te slaan
    var codes []string

    // Lees het bestand regel voor regel
    for scanner.Scan() {
        line := scanner.Text()
        // Splits de regel op ':'
        parts := strings.Split(line, ":")
        // Neem het laatste deel (de code)
        if len(parts) > 2 {
            code := parts[len(parts)-1]
            codes = append(codes, code)
        }
    }

    // Controleer op fouten bij het lezen
    if err := scanner.Err(); err != nil {
        fmt.Println("Fout bij lezen van bestand:", err)
        return
    }

    // Schrijf de codes naar een outputbestand
    output, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Fout bij maken van outputbestand:", err)
        return
    }
    defer output.Close()

    for _, code := range codes {
        _, err := output.WriteString(code + "\n")
        if err != nil {
            fmt.Println("Fout bij schrijven naar outputbestand:", err)
            return
        }
    }

    fmt.Println("Codes succesvol geëxtraheerd en opgeslagen in output.txt")
}
