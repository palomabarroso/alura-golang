# Alura - Go: a linguagem do Google

Exercícios do curso Go: a linguagem do Google da plataforma de estudos Alura

# Particularidades
- Linguagem compilada, leve
- Fortemente tipada
- Não existe while no GO, o equivalente é um laço for sem condição
- Funções, geralmente retornam dois valores. Um é a resposta e o outro a variável com o valor de um possível erro: Exemplo: 

```golang
resp, _ := http.Get(site)
```
# Variáveis
Ao declarar uma variável, o GO consegue inferir o seu tipo de acordo com sua inicialização.

# Estruturas de dados
## Array
Tipados e com quantidade pré determinada. 

### Funções
- len - Tamanho do array
- range - Retorna o item e o index do array
- cap - Retorna a capacidade do array

## Slices
Abstrações de array onde a própria linguagem gerencia entre outras coisas, o tamanho do array. 
Quando estouramos o tamanho do slice, ele dobra o tamanho. 

Exemplo: 
```golang
valores: = [ ]string{"Paloma", "Arthur"}
fmt.Println(len(valores)) -> retorna tamanho 2
fmt.Println(cap(valores)) -> retorna capacidade 2
```
```golang
valores := append(valores, "Vera")
fmt.Println(len(valores)) -> retorna tamanho 3
fmt.Println(cap(valores)) -> retorna capacidade 4
```

# Pacotes
## fmt 
Equivalente ao console.log e suas variações do js. Serve para imprimir dados em tela
## reflect 
Extrai informações em tempo de execução.
## os 
O Pacote OS fornece uma interface independente de plataforma para a funcionalidade do sistema operacional.
## net 
Pacote net fornece uma interface portátil para E/S de rede, incluindo TCP/IP, UDP, resolução de nome de domínio e soquetes de domínio Unix.

## time
Pacote fornece funcionalidade para medir e exibir o tempo.