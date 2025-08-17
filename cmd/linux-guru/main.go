package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	flag.Parse()
	ctx := context.Background()

	client := newClient(ctx)
	defer client.Close()

	model := setupModel(client)

	userPrompt := strings.Join(flag.Args(), " ")
	userPrompt = strings.TrimSpace(userPrompt)
	if userPrompt == "" {
		userPrompt = "Em que você pode me ajudar?"
	}

	resp, err := model.GenerateContent(ctx, genai.Text(userPrompt))
	if err != nil {
		log.Fatalf("Error sending message: %v\n", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}

func newClient(ctx context.Context) *genai.Client {
	apiKey, found := os.LookupEnv("GEMINI_API_KEY")
	if !found {
		log.Fatal("linux-guru: environment variable GEMINI_API_KEY not set. " +
			"Setup a free KEY at https://aistudio.google.com/app/apikey")
	}
	option := option.WithAPIKey(apiKey)

	client, err := genai.NewClient(ctx, option)
	if err != nil {
		log.Fatalf("linux-guru: error creating client: %v", err)
	}
	return client
}

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

func setupModel(client *genai.Client) *genai.GenerativeModel {
	model := client.GenerativeModel("gemini-2.5-flash")

	model.SetTemperature(0.2)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemPrompt)},
	}

	return model
}
