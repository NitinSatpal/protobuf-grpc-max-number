package main

import (
	crypto_rand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math"
	"max/pb"
	"max/server/store"
	"strconv"

	"max/server/store/types"

	"golang.org/x/crypto/nacl/box"
)

type Session struct {
	grpcnode pb.Service_MessageServer
	sid      string
	uid      string

	// Outbound mesages, buffered.
	send chan *pb.ServerMessage
	stop chan interface{}

	state int

	serverPrivateKey *[32]byte
	serverPublicKey  *[32]byte
	clientPublicKey  *[32]byte
}

func NewSession(grpcnode pb.Service_MessageServer) *Session {
	return &Session{
		sid:              generateID(10),
		grpcnode:         grpcnode,
		send:             make(chan *pb.ServerMessage),
		stop:             make(chan interface{}),
		clientPublicKey:  &[32]byte{},
		serverPrivateKey: &[32]byte{},
		serverPublicKey:  &[32]byte{},
	}
}

func (sess *Session) dispatch(m *pb.ClientMessage) {
	switch {
	case m.GetHello() != nil:
		sess.hello(m.GetHello())
	case m.GetNumber() != nil:
		sess.number(m.GetNumber())
	case m.GetCurrentMaxNumber() != nil:
		sess.currentMaxNumber(m.GetCurrentMaxNumber())
	case m.GetResetCurrentMaxNumber() != nil:
		sess.resetCurrentMaxNumber(m.GetResetCurrentMaxNumber())

	}
}

func (sess *Session) cleanUp() int {
	count := globals.sessionStore.Delete(sess)
	return count
}

func (sess *Session) closeGrpc() {
	sess.grpcnode = nil
	// sess.grpcnode.
}

func (sess *Session) hello(n *pb.ClientHello) {
	temp := sha256.Sum256(n.PublicId[:])

	// every client is identified by the hash of its public key
	id := base64.StdEncoding.EncodeToString(temp[:])

	user, err := store.Users.Get(id)

	// if user does not exist, create it
	if err != nil {
		user = &types.User{}
		user.Identity = id
		user.Key = base64.StdEncoding.EncodeToString(n.PublicId[:])
		// user.CurrentMaxNumber = 0
		user.CurrentMaxNumber = math.MinInt64

		_, err = store.Users.Create(user)
		if err != nil {
			panic("could not save user")
		}

	}

	// attach the user to the session
	sess.uid = user.Identity

	// user public key is saved as base64 encoded string, so it needs to be decoded
	k, _ := base64.StdEncoding.DecodeString(user.Key)
	copy(sess.clientPublicKey[:], k)

	sess.state = user.CurrentMaxNumber
	sess.serverPublicKey, sess.serverPrivateKey, _ = box.GenerateKey(crypto_rand.Reader)

	// send ServerHi back to the client which contains server public ID of this session
	sess.send <- &pb.ServerMessage{
		Message: &pb.ServerMessage_Hi{
			Hi: &pb.ServerHi{
				Id:       n.Id,
				PublicId: sess.serverPublicKey[:],
			},
		},
	}

	fmt.Printf("client %s connected\n", id)
}

func (sess *Session) number(n *pb.ClientNumber) {
	num, err := sess.DecryptToNumber(n.Number)
	if err != nil {
		return
	}

	if num > sess.state {
		// update the session
		sess.state = num

		// update the DB
		m := make(map[string]interface{})
		m["currentMaxNumber"] = sess.state
		err := store.Users.Update(sess.uid, m)
		if err != nil {
			fmt.Println("err", err)
		}

		// send back to client
		sess.send <- &pb.ServerMessage{
			Message: &pb.ServerMessage_Number{
				Number: &pb.ServerNumber{
					Id:     n.Id,
					Number: sess.EncryptFromNumber(sess.state),
				},
			},
		}
	}
}

func (sess *Session) currentMaxNumber(n *pb.ClientCurrentMaxNumber) {
	sess.send <- &pb.ServerMessage{
		Message: &pb.ServerMessage_Number{
			Number: &pb.ServerNumber{
				Id: n.Id,
				// Number: sess.state,
				Number: sess.EncryptFromNumber(sess.state),
			},
		},
	}
}

func (sess *Session) resetCurrentMaxNumber(n *pb.ClientResetCurrentMaxNumber) {
	// update the session
	sess.state = 0

	// update the DB
	m := make(map[string]interface{})
	m["currentMaxNumber"] = sess.state
	err := store.Users.Update(sess.uid, m)
	if err != nil {
		fmt.Println("err", err)
	}

	// send back to client
	sess.send <- &pb.ServerMessage{
		Message: &pb.ServerMessage_Number{
			Number: &pb.ServerNumber{
				Id:     n.Id,
				Number: sess.EncryptFromNumber(sess.state),
			},
		},
	}
}

// Decrypt takes an encrypted slice of bytes and return it unencrypted
func (sess *Session) Decrypt(encrypted []byte) ([]byte, error) {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	decrypted, ok := box.Open(nil, encrypted[24:], &decryptNonce, sess.clientPublicKey, sess.serverPrivateKey)
	if !ok {
		return nil, errors.New("Error Decrypting the number")
	}

	return decrypted, nil
}

// DecryptToNumber takes a slice of bytes and decrypt it and return an int
func (sess *Session) DecryptToNumber(encrypted []byte) (int, error) {
	d, err := sess.Decrypt(encrypted)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(string(d))
	if err != nil {
		return 0, err
	}

	return i, nil
}

// Encrypt takes a slice of bytes
func (sess *Session) Encrypt(msg []byte) []byte {
	nonce := genNonce()

	encrypted := box.Seal(nonce[:], msg, nonce, sess.clientPublicKey, sess.serverPrivateKey)

	return encrypted
}

// EncryptFromString takes a string
func (sess *Session) EncryptFromString(msg string) []byte {
	return sess.Encrypt([]byte(msg))
}

// EncryptFromNumber takes an int
func (sess *Session) EncryptFromNumber(n int) []byte {
	return sess.EncryptFromString(fmt.Sprintf("%d", n))
}

func genNonce() *[24]byte {
	var nonce [24]byte
	if _, err := io.ReadFull(crypto_rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	return &nonce
}
