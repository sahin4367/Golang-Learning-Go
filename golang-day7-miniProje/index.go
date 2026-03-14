package main

import (
	"fmt"
	"os"
)

type Student struct {
	name  string
	grade float64
}

func (g Student) Bal(b float64) string {
	switch {
	case b >= 91:
		return "A"
	case b >= 81:
		return "B"
	case b >= 71:
		return "C"
	case b >= 61:
		return "D"
	case b >= 51:
		return "E"
	default:
		return "F - Kesildin!"
	}
}

func averageGrade(students []Student) float64 {
	total := 0.0
	for _, s := range students {
		total += s.grade
	}
	return total / float64(len(students))
}

func generateHTML(students []Student) string {
	rows := ""

	for _, s := range students {
		rows += fmt.Sprintf(`
		<tr>
			<td>%s</td>
			<td>%.1f</td>
			<td>%s</td>
		</tr>`, s.name, s.grade, s.Bal(s.grade))
	}

	avg := averageGrade(students)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Student Grade Tracker</title>
    <style>
        body { font-family: Arial; padding: 30px; background: #f0f2f5; }
        h1 { color: #2c3e50; }
        table { border-collapse: collapse; width: 100%%; background: white; border-radius: 8px; box-shadow: 0 2px 8px rgba(0,0,0,0.1); }
        th, td { padding: 12px 16px; border: 1px solid #ddd; text-align: left; }
        th { background: #4CAF50; color: white; }
        tr:hover { background: #f9f9f9; }
        .avg { margin-top: 20px; font-size: 18px; color: #2c3e50; }
    </style>
</head>
<body>
    <h1>🎓 Student Grade Tracker</h1>
    <table>
        <tr>
            <th>Ad</th>
            <th>Qiymət</th>
            <th>Dərəcə</th>
        </tr>
        %s
    </table>
    <p class="avg">📊 Orta Qiymət: <strong>%.1f</strong></p>
</body>
</html>`, rows, avg)

	return html
}

func main() {
	students := []Student{
		{name: "Shahin", grade: 74},
		{name: "Vusal", grade: 45},
		{name: "Ali", grade: 99},
	}

	// Terminal output
	for _, s := range students {
		fmt.Println(s.name, "-", s.Bal(s.grade))
	}
	fmt.Println("Orta qiymət:", averageGrade(students))

	// HTML yarat və fayla yaz
	html := generateHTML(students)
	err := os.WriteFile("index.html", []byte(html), 0644)
	if err != nil {
		fmt.Println("Xəta:", err)
	} else {
		fmt.Println("index.html yaradıldı! Brauzerda aç!")
	}
}
