package command

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var generateJWTKey = &cobra.Command{
	Use:   "generate-jwt-key",
	Short: "Generate jwt key",
	Long:  `Generate jwt key`,
	Run: func(cmd *cobra.Command, args []string) {
		generateKey()
	},
}

func init() {
	rootCmd.AddCommand(generateJWTKey)
}

func generateKey() {
	if jwt := os.Getenv("JWT_KEY"); jwt != "" {
		return
	}

	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		panic(err)
	}

	jwtKey := hex.EncodeToString(key)

	os.Setenv("JWT_KEY", jwtKey)

	envFileContent, err := os.ReadFile("../../.env")

	if err != nil {
		log.Fatalf("Error readling .env file: %v", err)
		return
	}

	content := string(envFileContent)
	newLine := fmt.Sprintf("JWT_KEY=%s", jwtKey)

	if strings.Contains(content, "JWT_KEY=") {
		content = strings.Replace(content, "JWT_KEY=", newLine, 1)
	} else {
		content = content + "\n" + newLine
	}

	if err = os.WriteFile("../../.env", []byte(content), 0664); err != nil {
		log.Fatalf("Error writing to .env file: %v", err)

		return
	}

	return
}
