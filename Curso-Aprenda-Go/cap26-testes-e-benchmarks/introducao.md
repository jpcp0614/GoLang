Testes devem:
    - ficar num arquivo cuja terminação seja _test.go
    - ficar na mesma package que o código a ser testado
    - ficar em funções com nome "func TestNome(*testing.T)"
- Para rodar os testes:
    - go test
    - go test -v
- Para falhas, utilizamos t.Error(), onde a maneira idiomática é algo do tipo "expected: x. got: y."