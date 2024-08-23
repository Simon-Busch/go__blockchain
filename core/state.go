package core

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Simon-Busch/go__blockchain/types"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrInsufficientBalance = errors.New("insufficient balance")
)

type Account struct {
	Balance 		uint64
	Address 		types.Address
}

func (a *Account) String() string {
	return fmt.Sprintf("Account{Address: %s, Balance: %d}", a.Address, a.Balance)
}

type AccountState struct {
	mu    				sync.RWMutex
	accounts 			map[types.Address]*Account
}


func NewAccountState() *AccountState {
	return &AccountState{
		accounts: make(map[types.Address]*Account),
	}
}

func (s *AccountState) CreateAccount(address types.Address) *Account {
	s.mu.Lock()
	defer s.mu.Unlock()

	acc := &Account{
		Address: 		address,
	}
	s.accounts[address] = acc
	return acc
}

func (s *AccountState) GetAccount(address types.Address) (*Account, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.getAccountWithoutLock(address)
}

func (s *AccountState) getAccountWithoutLock(address types.Address) (*Account, error) {
	account, ok := s.accounts[address]
	if !ok {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (s *AccountState) GetBalance(address types.Address) (uint64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account, err := s.getAccountWithoutLock(address)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (s *AccountState) Transfer(from, to types.Address, amount uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fromAccount, err := s.getAccountWithoutLock(from)

	if err != nil {
		return err
	}

	// If it's the coinbase
	if fromAccount.Address.String() != "996fb92427ae41e4649b934ca495991b7852b855" {
		if fromAccount.Balance < amount  {
			return ErrInsufficientBalance
		}
	}

	fromAccount.Balance -= amount

	if s.accounts[to] == nil {
		s.accounts[to] = &Account{
			Address: 		to,
		}
	}

	s.accounts[to].Balance += amount

	return nil
}


// func (s *AccountState) Transfer(from, to types.Address, amount uint64) error {
// 	if err := s.SubBalance(from, amount); err != nil {
// 		return err
// 	}

// 	return s.AddBalance(to, amount)
// }

// func (s *AccountState) GetBalance(to types.Address) (uint64, error) {
// 	balance, ok := s.state[to]
// 	if !ok {
// 		return 0.0, fmt.Errorf("address (%s) unkown", to)
// 	}

// 	return balance, nil
// }

// func (s *AccountState) SubBalance(to types.Address, amount uint64) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	balance, ok := s.state[to]
// 	if !ok {
// 		return fmt.Errorf("address (%s) unknown", to)
// 	}

// 	if balance < amount {
// 		return fmt.Errorf("insuccient account balance (%d) for amount (%d)", balance, amount)
// 	}

// 	s.state[to] -= amount

// 	return nil
// }

// func (s *AccountState) AddBalance(to types.Address, amount uint64) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.state[to] += amount

// 	return nil
// }

type State struct {
	data 					map[string][]byte
}

func NewState() *State {
	return &State{
		data: make(map[string][]byte),
	}
}

func (s *State) Put(k, v []byte) error {
	s.data[string(k)] = v

	return nil
}

func (s *State) Delete(k []byte) error {
	delete(s.data, string(k))

	return nil
}

func (s *State) Get(k []byte) ([]byte, error) {
	key := string(k)
	value, ok := s.data[key]
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return value, nil
}
