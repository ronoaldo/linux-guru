## linux-guru: seu assistente pessoal de IA para o Terminal 

Este projeto utiliza o poder do Gemini AI, um modelo de linguagem avançado do
Google, para fornecer funcionalidades inteligentes. Para usar o programa, siga
os passos abaixo:

### Instalação

1. **Pré-requisitos:**
   - Go (versão 1.18 ou superior) instalado em seu sistema.
   - Conta Google pessoal para acessar o Google AI Studio.

2. **Instalação do Programa:**
   - Abra seu terminal e execute o seguinte comando:
```bash
go install github.com/ronoaldo/linux-guru/cmd/linux-guru@latest
```
   - Isso irá baixar e instalar o programa `linux-guru` em seu sistema.

### Obtenção e Configuração da Chave de API do Gemini (Google AI Studio)

1. **Acesse o Google AI Studio:**
   - Faça login em sua conta Google pessoal.
   - Acesse o site do Google AI Studio:
     [https://aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey)

2. **Crie uma Nova Chave de API:**
   - Clique no botão "Get API Key" (Obter Chave de API).
   - Siga as instruções para criar uma nova chave de API.
   - Uma chave de API única será gerada para você.

3. **Configure a Variável de Ambiente:**
   - Copie a chave de API obtida.
   - Abra seu terminal e execute o seguinte comando, substituindo
     `<SUA_CHAVE_DE_API>` pela chave que você copiou:
```bash export
GEMINI_API_KEY=<SUA_CHAVE_DE_API>
```

### Uso do Programa

Agora que o programa está instalado e a variável de ambiente `GEMINI_API_KEY`
está configurada, você pode executar o programa `linux-guru` em seu terminal
para começar a utilizá-lo. Consulte a documentação específica do projeto para
saber como interagir com o Gemini AI e aproveitar seus recursos.

**Importante: na versão gratuita, os dados enviados são utilizados pelo Google
para aprimorar os modelos. Caso queira, ative o faturamento e faça upgrade
para uma conta paga onde seus dados são privados.**

