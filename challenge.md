### Desafio 1: Simplificação de Expressões Booleanas
**Descrição:**
Dado um array de expressões booleanas, você deve simplificá-las eliminando redundâncias. Por exemplo, se a expressão contém `(A && !A)`, ela deve ser simplificada para `false`.

**Exemplo:**
```javascript
const expressions = [
  "A && !A",
  "A && B || !A && !B",
  "(A || !A) && B"
];

function simplifyExpressions(expressions) {
  // Implemente a simplificação das expressões aqui
}

console.log(simplifyExpressions(expressions));
// Saída esperada: [false, "B || !B", "B"]
```


### Desafio 2: Parênteses Balanceados
**Descrição:**
Escreva uma função que verifica se os parênteses em uma string estão balanceados. Uma string está balanceada se todo parêntese de abertura tiver um parêntese de fechamento correspondente e se os parênteses forem corretamente aninhados.

**Exemplo:**
```javascript
const str1 = "((A && B) || C)";
const str2 = "((A && B) || (C || D))";
const str3 = "((A && B) || C";

function areParenthesesBalanced(str) {
  // Implemente a verificação de balanceamento aqui
}

console.log(areParenthesesBalanced(str1)); // true
console.log(areParenthesesBalanced(str2)); // true
console.log(areParenthesesBalanced(str3)); // false
```


### Desafio 3: Combinação de Números para Alvo
**Descrição:**
Dado um array de números inteiros e um número alvo, escreva uma função que retorna todas as combinações possíveis de números que somam o valor do alvo. Você pode usar cada número uma vez em cada combinação.

**Exemplo:**
```javascript
const nums = [2, 3, 6, 7];
const target = 7;

function combinationSum(nums, target) {
  // Implemente a lógica de combinação aqui
}

console.log(combinationSum(nums, target));
// Saída esperada: [[7], [2, 2, 3]]
```
