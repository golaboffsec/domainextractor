package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func printBanner() {
	banner := `
┌┬┐┌─┐┌┬┐┌─┐┬┌┐┌            
 │││ ││││├─┤││││            
─┴┘└─┘┴ ┴┴ ┴┴┘└┘            
┌─┐─┐ ┬┌┬┐┬─┐┌─┐┌─┐┌┬┐┌─┐┬─┐
├┤ ┌┴┬┘ │ ├┬┘├─┤│   │ │ │├┬┘
└─┘┴ └─ ┴ ┴└─┴ ┴└─┘ ┴ └─┘┴└─
		By: lupedsagaces 
                                                                          
`
	fmt.Println(banner)

}
func main() {
	printBanner()
	// Solicita ao usuário o nome do arquivo
	var fileName string
	fmt.Print("Informe o nome do arquivo com as URLs (exemplo: dominios.txt): ")
	fmt.Scanln(&fileName)

	// Lê o conteúdo do arquivo
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Cria um slice para armazenar as URLs filtradas
	var filteredUrls []string
	scanner := bufio.NewScanner(file)

	// Solicita ao usuário o domínio para filtrar
	var domain string
	fmt.Print("Informe o domínio que deseja extrair (exemplo: bol): ")
	fmt.Scanln(&domain)

	// Filtra as URLs que contêm o domínio fornecido
	for scanner.Scan() {
		url := scanner.Text()
		if strings.Contains(url, domain) {
			filteredUrls = append(filteredUrls, url)
		}
	}

	// Verifica se houve erro na leitura do arquivo
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Reinicia a leitura do arquivo
	file.Seek(0, 0) // Reseta o ponteiro do arquivo para o início

	// Cria um mapa para armazenar os domínios únicos
	domainSet := make(map[string]struct{})
	scanner2 := bufio.NewScanner(file)

	// Expressão regular para verificar se o domínio está no formato correto
	regex := regexp.MustCompile(fmt.Sprintf(`(^|\.)[a-zA-Z0-9-]*%s\.[a-zA-Z]{2,}`, domain))

	// Filtra as URLs que atendem ao critério
	for scanner2.Scan() {
		url := scanner2.Text()
		// Extrai apenas o domínio da URL
		parts := strings.Split(url, "/")
		if len(parts) > 0 {
			domain := parts[0]
			// Verifica se o domínio atende ao critério
			if regex.MatchString(domain) {
				domainSet[domain] = struct{}{} // Adiciona ao mapa para evitar duplicados
			}
		}
	}

	// Verifica se houve erro na leitura do arquivo
	if err := scanner2.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Escreve os domínios únicos em um novo arquivo
	outputFileName := "dominios_unicos.txt"
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer outputFile.Close()

	// Escreve os domínios únicos no arquivo
	for domain := range domainSet {
		outputFile.WriteString(domain + "\n")
	}

	fmt.Printf("Arquivo '%s' criado com os domínios únicos.\n", outputFileName)
}
