# Tema: Idosos e o isolamento 
Durante a quarentena, a população idosa ficou impedida de ver familiares e de realizar as atividades diárias que possuíam. Com isso, a saúde mental de muitos se viu fragilizada, acometendo-os de medo, estresse e solidão. Muitos não possuíam a habilidade básica de comunicação pelas redes sociais, dificultando, ainda mais, a sua interação com a vida fora de casa. Que ações podem ser tomadas para que estes efeitos sejam minimizados na vida do idoso?

# Idéia:
Sistema formato de um app para idosos (preciosos), que acompanha individualmente cada idoso através de um login especial (líder discipulador) via web.

[Apresentação detalhada do projeto](https://docs.google.com/presentation/d/1BXm2HNW1eh9b9wU_JP_d342fmw4iCKRggohnpPkIkwcd)

# Objetivos:
**Objetivo principal**
- Fornecer uma solução para o discipulado de pessoas da terceira idade (preciosos), com acompanhamento individual
 
**Objetivos secundários**
- Orientar na saúde mental do precioso
- Oferecer conteúdo bíblico
- Gerar engajamento e interação do precioso

# Como executar:
O projeto foi desenvolvido utilizando a linguagem de programação [GO](https://golang.org). Uma vez instalado, basta executar:

```go
go run main.go
```

Para visualizar o sistema, acesse o seguinte URL:
```
http://127.0.0.1:9090
```

Devido ao tempo limitado, apenas as funcionalidades básicas foram implementas através de chamadas REST api.

- Criar discipulado:
```
curl -X POST http://127.0.0.1:9090/discipulados -H 'Content-Type: application/json' -d '{"nome":"Joao da Silva","aniversario": "01/01/1968", "contato": { "endereco":"Rua Mestre Pastinha, n. 147, Edf. Beija-Flor, Apt. 248", "cidade": "Salvador, Bahia", "telefone": "7155556666", "email":"joao@ig.com.br"}}'
```

- Listar discipulado:
```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://127.0.0.1:9090/discipulados/3f9cead049d4e688231cb146086b3dea999db6b7af00fc756d68ebc3b83dc65f
```

- Remover discipulado:
```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X DELETE http://127.0.0.1:9090/discipulados/3f9cead049d4e688231cb146086b3dea999db6b7af00fc756d68ebc3b83dc65f
```
