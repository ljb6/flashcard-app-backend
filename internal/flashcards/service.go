package flashcards

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var KEY_SECRET string

func GetKeys() {
	KEY_SECRET = os.Getenv("KEY")
}

type FlashcardService struct {
	repository *FlashcardRepository
}

func NewFlashcardService(repository *FlashcardRepository) *FlashcardService {
	return &FlashcardService{repository: repository}
}

func (s *FlashcardService) CreateFlashcard(front, back string) error {
	if len(front) > 250 || len(back) > 250 {
		return errors.New("flashcard content exceeds maximum length of 250 characters")
	}

	err := s.repository.CreateFlashcard(front, back)
	if err != nil {
		return err
	}

	return nil
}

func (s *FlashcardService) GetAllFlashcards() ([]byte, error) {
	flashcards, err := s.repository.GetAllFlashcards()
	if err != nil {
		return nil, err
	}

	jsonFlashcards, err := json.Marshal(flashcards)
	if err != nil {
		return nil, errors.New("error in marshal")
	}

	return jsonFlashcards, nil
}

func (s *FlashcardService) GetXFlashcards(req GetFlashcardsReq) ([]byte, error) {
	flashcards, err := s.repository.GetXFlashcards(req)
	if err != nil {
		return nil, err
	}

	jsonFlashcards, err := json.Marshal(flashcards)
	if err != nil {
		return nil, errors.New("error in marshal")
	}

	return jsonFlashcards, nil
}

func (s *FlashcardService) GetXFlashcardsByError(req GetFlashcardsReq) ([]byte, error) {
	flashcards, err := s.repository.GetXFlashcardsByError(req)
	if err != nil {
		return nil, err
	}

	jsonFlashcards, err := json.Marshal(flashcards)
	if err != nil {
		return nil, errors.New("error in marshal")
	}

	return jsonFlashcards, nil
}

func (s *FlashcardService) DeleteFlashcardByID(id int) error {
	err := s.repository.DeleteFlashcardByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *FlashcardService) EditFlashcardByID(id int, front, back string) error {
	if len(front) > 250 || len(back) > 250 {
		return errors.New("flashcard content exceeds maximum length of 250 characters")
	}

	err := s.repository.EditFlashcardByID(id, front, back)
	if err != nil {
		return err
	}

	return nil
}

func (s *FlashcardService) DeleteAllFlashcards() error {
	err := s.repository.DeleteAllFlashcards()
	if err != nil {
		return err
	}
	return nil
}

func (s *FlashcardService) UpdateFlashcardFields(id int, correct bool) error {
	card, err := s.repository.GetFlashcardByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	now := time.Now()
	card.LastReview = &now

	if correct {
		if card.ReviewStage < len(ReviewIntervals)-1 {
			card.ReviewStage++
		}
		card.CorrectAnswers++
	} else {
		card.IncorrectAnswers++
		card.ReviewStage = 0
	}

	err = s.repository.UpdateFlashcardFields(card)
	if err != nil {
		return err
	}
	return nil
}

func (s *FlashcardService) GetDueFlashcards() ([]byte, error) {
	dueCards, err := s.repository.GetDueFlashcards()
	if err != nil {
		return nil, err
	}

	jsonFlashcards, err := json.Marshal(dueCards)
	if err != nil {
		return nil, errors.New("error in marshal")
	}

	return jsonFlashcards, nil
}

func (s *FlashcardService) CreateFlashcardsWithAI(theme string) error {

	prompt := fmt.Sprintf("Crie 5 flashcards básicos em pt-br sobre %s no formato pergunta e resposta. Mande separado por linhas (pergunta, resposta). Não escreva nada a mais. Cada linha pode ter no máximo 200 caracteres. SEMPRE SIGA ESSE FORMATO. SEMPRE 5 FLASHCARDS. NUNCA ESCREVA pergunta ou resposta nas frases.", theme)
	auth := fmt.Sprintf("Beares %s", KEY_SECRET)

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "meta-llama/llama-3.3-8b-instruct:free",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	})
	if err != nil {
		fmt.Println("Error creating request body:", err)
		return err
	}

	// Create request
	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// Set headers
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	content := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	//fmt.Println("Content: ", content)

	cards := strings.Split(content, "\n\n")
	for _, part := range cards {

		//fmt.Printf("Parte %d:\n%s\n\n", i+1, part)
		sides := strings.Split(part, "\n")
		fmt.Println(sides[0], " : ", sides[1])

		s.CreateFlashcard(sides[0], sides[1])
	}



	return nil
}
