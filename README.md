O que este programa faz:

Envia duas requisições simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000

http://viacep.com.br/ws/01153000/json/

Requisitos:

- Acatar a API que entrega a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
