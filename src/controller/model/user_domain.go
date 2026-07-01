package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/rest_err"
)

// NewUserDomain é uma "factory function": recebe os dados básicos de um usuário
// e retorna um ponteiro para UserDomain, mas tipado como a INTERFACE
// UserDomainInterface (não como o struct concreto).
//
// Isso é um padrão comum em Go para desacoplamento: quem chama essa função
// só enxerga os métodos definidos na interface (CreateUser, UpdateUser,
// FindUser, DeleteUser), e não os detalhes internos do struct.
func NewUserDomain(
	name, password, email string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		name, password, email, age,
	}
}

// UserDomain representa o usuário na camada de domínio (regra de negócio),
// diferente das structs de Request/Response que servem só para comunicação HTTP.
// Não tem tags `json:"..."` porque não é (e não deveria ser) serializada
// diretamente numa resposta de API sem cuidado (a senha, por exemplo,
// acabaria vazando — ver createUser.go).
type UserDomain struct {
	Name     string
	Password string
	Email    string
	Age      int8
}

// EncryptPassword transforma a senha em texto puro em um hash MD5.
// Isso evita que a senha original fique salva/armazenada em texto puro.
//
// ATENÇÃO (aviso técnico): MD5 é considerado inseguro para hash de senhas
// hoje em dia (é rápido demais e vulnerável a ataques de força bruta e
// rainbow tables). Para um projeto real, o recomendado seria usar bcrypt,
// scrypt ou argon2, que são feitos especificamente para senhas.
func (ud *UserDomain) EncryptPassword() {
	// Cria um novo "hasher" MD5.
	hash := md5.New()
	// defer hash.Reset() -> zera o estado interno do hash ao final da função
	// (mais por boa prática/limpeza do que por necessidade real aqui).
	defer hash.Reset()
	// Escreve os bytes da senha original dentro do hash.
	hash.Write([]byte(ud.Password))
	// Sum(nil) calcula o hash final (em bytes) e hex.EncodeToString
	// converte esses bytes para uma string hexadecimal legível.
	// Substitui a senha original pela senha "criptografada" (hash).
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

// UserDomainInterface define o "contrato" das operações de CRUD que
// qualquer implementação de domínio de usuário precisa ter.
// Isso permite, por exemplo, trocar UserDomain por um mock em testes,
// já que o resto do código depende da interface, não do struct concreto.
type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

//A struct userDomain representa o usuário dentro da camada de domínio da aplicação. Diferente das structs utilizadas para comunicação com a API, ela não possui tags JSON porque não é utilizada diretamente para serialização ou desserialização de requisições HTTP. O método EncryptPassword() é responsável por transformar a senha em um hash MD5 antes do armazenamento no banco de dados, evitando que a senha original seja salva em texto puro. Já a interface UserDomain define um contrato contendo as operações básicas do CRUD, como criar, atualizar, buscar e remover usuários, permitindo desacoplamento entre as regras de negócio e as implementações dessas operações.
