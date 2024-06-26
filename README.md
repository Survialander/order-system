# Order System

## Como executar

- Primeiro clone o repositório:
```
git clone git@github.com:Survialander/order-system.git
```

- Na pasta raiz do projeto, rode os seguintes comandos:
```
// Inicializa o banco de dados e o rabbitmq

docker-compose up -d
```

```
// Execute a aplicação

cd cmd/ordersysterm
go run main.go wire_gen.go
```

## Executando testes

### Api REST

utilize os arquivos .http encontrados na pasta `api` do projeto

### Api GRAPHQL

Acesse o playground, ele se encontra na seguinte url: [GraphQL](http://localhost:8080)

utilize as seguintes queries:
```
// Cria ordem
mutation createOrder {
	createOrder(input: { id: "idteste", Price: 50, Tax: 2 }) {
		id
		Price
		Tax
		FinalPrice
	}
}

// Lista ordens
query listOrders {
	listOrders {
		id
		Price
		Tax
		FinalPrice
	}
}
```

### API gRPC

Utilize o arquivo orders.proto, ele se encontra no seguinte caminho `internal/infra/grpc/proto/orders.proto`, você pode utilizar o postman. Caso não saiba utilizar siga os passos no [tutorial](https://learning.postman.com/docs/sending-requests/grpc/grpc-request-interface/)

após importar o arquivo no postman você pode fazer as requests usando o seguinte endpoint `grpc://localhost:50051`
