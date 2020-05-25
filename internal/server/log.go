package server

import (
	"fmt"
	"sync"
)

type Log struct {
	//mu é um Mutex
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}
func (c *Log) Append(record Record) (uint64, error) {
	//tranca o mutex no início da execução da função
	c.mu.Lock()
	//destranca o mutex no final da execução da função
	defer c.mu.Unlock()
	//salva o offset da gravação atual, que será o tamanho do slice de gravações anterior
	record.Offset = uint64(len(c.records))
	//adiciona a nova gravação no final do vetor de gravações
	c.records = append(c.records, record)
	return record.Offset, nil
}

//função para ler uma gravação
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//se o offset for maior que o tamanho do slice de gravações, retorne um erro
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrorOffsetNotFound
	}
	return c.records[offset], nil
}

//a estrutura record retorna um json que é uma variável que manipula um json
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrorOffsetNotFound = fmt.Errorf("offset not found")
