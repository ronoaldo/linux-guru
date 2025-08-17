package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

var model = flag.String("model", "gemini-2.5-flash", "The model to use when calling")

var systemPrompt = `Você é um assistente de usuários de terminais de comandos do Linux.
Especializado em várias distribuições de Linux e GNU/Linux.
Você só responde a perguntas sobre Linux, como instalação de pacotes,
origem e criação do Linux, programação utilizando Shell Script.
Suas respostas devem ser simples.
Evite respostas muito prolixas.
Para qualquer outra pergunta você responde:
"Não conheço sobre este assunto! Tente uma pergunta sobre Linux".
Se a resposta tiver algum código fonte ou comando, incluir no final:
"Atenção: sempre revise o código e comandos antes de executá-los!"`

var config = &genai.GenerateContentConfig{
	Temperature:       genai.Ptr[float32](0.2),
	MaxOutputTokens:   8192,
	ResponseMIMEType:  "text/plain",
	SystemInstruction: genai.Text(systemPrompt)[0],
}

func main() {
	flag.Parse()
	ctx := context.Background()

	client := newClient(ctx)

	userPrompt := strings.Join(flag.Args(), " ")
	userPrompt = strings.TrimSpace(userPrompt)
	if userPrompt == "" {
		userPrompt = "Em que você pode me ajudar?"
	}

	resp, err := client.Models.GenerateContent(ctx, *model, genai.Text(userPrompt), config)
	if err != nil {
		log.Fatalf("Error sending message: %v\n", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part.Text)
	}
}

func newClient(ctx context.Context) *genai.Client {
	apiKey, found := os.LookupEnv("GEMINI_API_KEY")
	if !found {
		log.Fatal("linux-guru: environment variable GEMINI_API_KEY not set. " +
			"Setup a free KEY at https://aistudio.google.com/app/apikey")
	}
	config := &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	}

	client, err := genai.NewClient(ctx, config)
	if err != nil {
		log.Fatalf("linux-guru: error creating client: %v", err)
	}
	return client
}
