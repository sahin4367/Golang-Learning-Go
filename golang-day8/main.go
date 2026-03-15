package main

import (
	"errors"
	"fmt"
)

// ─── 1. CUSTOM ERROR ───────────────────────────────────────────

type DBError struct {
	Code    int
	Message string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("DB xətası [%d]: %s", e.Code, e.Message)
}

// ─── 2. SENTINEL ERROR ─────────────────────────────────────────

var ErrBağlantı = errors.New("bağlantı xətası")

// ─── 3. AŞAĞI MƏRTƏBƏ — xəta burada doğur ─────────────────────

func dbdenOxu(id int) error {
	// Fərz et database cavab vermədi
	return &DBError{
		Code:    500,
		Message: "bağlantı kəsildi",
	}
}

// ─── 4. ORTA MƏRTƏBƏ — wrap edir ───────────────────────────────

func istifadəçiAl(id int) error {
	err := dbdenOxu(id)
	if err != nil {
		return fmt.Errorf("istifadəçiAl(%d): %w", id, err)
	}
	return nil
}

// ─── 5. MAIN — hər şeyi yoxlayır ───────────────────────────────

func main() {

	err := istifadəçiAl(5)

	if err == nil {
		fmt.Println("Uğurlu!")
		return
	}

	// Xətanın tam yolunu gör
	fmt.Println("Xəta:", err)
	fmt.Println("─────────────────────────")

	// errors.Is — sentinel error varmı?
	if errors.Is(err, ErrBağlantı) {
		fmt.Println("IS: Bağlantı problemidir — yenidən cəhd et")
	}

	// errors.As — DBError tipindən varmı?
	var dbErr *DBError
	if errors.As(err, &dbErr) {
		fmt.Println("AS: Tip tapıldı:")
		fmt.Println("    Kod    →", dbErr.Code)
		fmt.Println("    Mesaj  →", dbErr.Message)
	}
}

/*
Xəta: istifadəçiAl(5): DB xətası [500]: bağlantı kəsildi
─────────────────────────
IS: Bağlantı problemidir — yenidən cəhd et
AS: Tip tapıldı:
    Kod    → 500
    Mesaj  → bağlantı kəsildi
*/
