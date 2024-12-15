package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jonathanhecl/chunker"
	"github.com/jonathanhecl/gollama"
)

func main() {
	fmt.Println("RAG Example")

	filename := "text.txt"
	embedding_model := "nomic-embed-text"
	chat_model := "llama3.2"

	// Create a Gollama instances
	e := gollama.New(embedding_model)
	e.PullIfMissing()

	c := gollama.New(chat_model)
	c.PullIfMissing()

	fmt.Println("Embedding model:", embedding_model)
	fmt.Println("Chat model:", chat_model)

	// Read the text file
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(f)

	fmt.Println("File ", filename, "has", len(text), "bytes...")

	// Chunk the text
	chunks := chunker.ChunkSentences(text)

	fmt.Println("Total chunks:", len(chunks))

	// Embed the chunks
	embeds := make([][]float64, 0)
	for _, chunk := range chunks {
		embedding, err := e.Embedding(chunk)
		if err != nil {
			fmt.Println(err)
			return
		}
		embeds = append(embeds, embedding)
	}

	// Save into a struct
	type tEmbedding struct {
		Chunk string
		Embed []float64
	}
	embeddings := make([]tEmbedding, 0)
	for i, embedding := range embeds {
		embeddings = append(embeddings, tEmbedding{Chunk: chunks[i], Embed: embedding})
	}

	fmt.Println("Total embeddings:", len(embeddings))

	// Run the chat loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a question ('q' to quit):")
		var question string
		if !scanner.Scan() {
			break
		}
		question = scanner.Text()

		if question == "q" {
			return
		}

		// Get the question embedding
		question_emb, _ := e.Embedding(question)

		// Search contexts
		contexts := make([]string, 0)
		for _, embedding := range embeddings {
			similarity := gollama.CosenoSimilarity(question_emb, embedding.Embed)
			if similarity > 0.65 {
				fmt.Println("> Context:", embedding.Chunk+" (Similarity: "+fmt.Sprintf("%.2f", similarity)+")")
				contexts = append(contexts, embedding.Chunk)
			}
		}

		if len(contexts) == 0 {
			fmt.Println("> No context found")
			continue
		}

		// Create the prompt
		prompt := "Respond to the following question using the provided context, don't add anything else:\n\n" +
			"Context:\n" + strings.Join(contexts, "\n") + "\n\nQuestion:\n" + question

		fmt.Println("Prompt:", prompt)

		// Get the answer
		answer, err := c.Chat(prompt)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		fmt.Println("> Answer:", answer.Content)
		fmt.Println()
	}

}
