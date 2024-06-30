# Projeto de Demonstração de PC Fractal

## Introdução
Este projeto em Go demonstra o uso de um PC fractal para resolver dois problemas distintos de maneira eficiente: estimativa de Pi usando o método de Monte Carlo e cálculo da probabilidade de obter pelo menos um número par ao lançar um dado seis vezes. O PC fractal integra cálculos discretos e simulações analógicas em um ambiente paralelo.

## PC Fractal: Visão Geral
Um PC fractal, ou Computador Fractal, combina capacidades de processamento paralelo com estruturas recursivas ou auto-similares. Ele é projetado para maximizar a eficiência e o desempenho ao lidar com problemas complexos através de múltiplos núcleos de processamento e estruturas que se repetem em diferentes escalas. Principais características incluem:

- **Paralelismo Explícito:** Utiliza múltiplos processadores para realizar tarefas simultâneas e independentes.
- **Estruturas Recursivas:** Incorpora elementos auto-similares que permitem otimizações em algoritmos complexos.
- **Flexibilidade e Escalabilidade:** Pode escalar para lidar com problemas de diferentes tamanhos e complexidades.

### Aplicações Comuns de um PC Fractal:
- **Simulações Complexas:** Como estimativa de Pi usando Monte Carlo.
- **Processamento de Dados Massivos:** Análise de big data e aprendizado de máquina.
- **Modelagem e Renderização:** Computação gráfica e visualização científica.

### Benefícios:
- **Desempenho Aumentado:** Melhora significativa de desempenho comparado a sistemas tradicionais.
- **Eficiência Energética:** Uso eficiente de energia através de distribuição de tarefas.
- **Adaptabilidade:** Flexível para diferentes tipos de problemas e requisitos de computação.

## Estrutura do Projeto
- **main.go:** Contém a implementação principal do programa, incluindo a definição das funções `analogProcessing` e `discreteProcessing`, além da função `main` que coordena a execução do programa.
- **README.md:** Este arquivo, fornecendo uma visão geral do projeto, sua finalidade e explicação do conceito de PC fractal.

## Como Executar
Para executar o projeto, certifique-se de ter o Go instalado e configurado em seu ambiente. Em seguida, execute o seguinte comando na linha de comando:

```bash
go run main.go
