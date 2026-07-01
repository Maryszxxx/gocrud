package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetPassword() string
	GetAge() int8

	EncryptPassword()
}

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
	return &userDomain{
		name, password, email, age,
	}
}

// UserDomain representa o usuário na camada de domínio (regra de negócio),
// diferente das structs de Request/Response que servem só para comunicação HTTP.
// Não tem tags `json:"..."` porque não é (e não deveria ser) serializada
// diretamente numa resposta de API sem cuidado (a senha, por exemplo,
// acabaria vazando — ver createUser.go).
type userDomain struct {
	name     string
	password string
	email    string
	age      int8
}

// EncryptPassword transforma a senha em texto puro em um hash MD5.
// Isso evita que a senha original fique salva/armazenada em texto puro.
//
// ATENÇÃO (aviso técnico): MD5 é considerado inseguro para hash de senhas
// hoje em dia (é rápido demais e vulnerável a ataques de força bruta e
// rainbow tables). Para um projeto real, o recomendado seria usar bcrypt,
// scrypt ou argon2, que são feitos especificamente para senhas.
func (ud *userDomain) EncryptPassword() {
	// Cria um novo "hasher" MD5.
	hash := md5.New()
	// defer hash.Reset() -> zera o estado interno do hash ao final da função
	// (mais por boa prática/limpeza do que por necessidade real aqui).
	defer hash.Reset()
	// Escreve os bytes da senha original dentro do hash.
	hash.Write([]byte(ud.password))
	// Sum(nil) calcula o hash final (em bytes) e hex.EncodeToString
	// converte esses bytes para uma string hexadecimal legível.
	// Substitui a senha original pela senha "criptografada" (hash).
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}
