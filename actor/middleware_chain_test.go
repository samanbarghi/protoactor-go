package actor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func middleware(called *int) func(ReceiveFunc) ReceiveFunc {
	return func(next ReceiveFunc) ReceiveFunc {
		fn := func(context Context) {
			*called = context.Message().(int)

			next(context)
		}
		return fn
	}
}

func TestMakeReceiverMiddleware_CallsInCorrectOrder(t *testing.T) {
	var c [3]int

	r := []func(ReceiveFunc) ReceiveFunc{
		middleware(&c[0]),
		middleware(&c[1]),
		middleware(&c[2]),
	}

	mc := &mockContext{}
	mc.On("Message").Return(1).Once()
	mc.On("Message").Return(2).Once()
	mc.On("Message").Return(3).Once()

	chain := makeMiddlewareChain(r, func(_ Context) {})
	chain(mc)

	assert.Equal(t, 1, c[0])
	assert.Equal(t, 2, c[1])
	assert.Equal(t, 3, c[2])
	mock.AssertExpectationsForObjects(t, mc)
}

func TestMakeReceiverMiddleware_ReturnsNil(t *testing.T) {
	assert.Nil(t, makeMiddlewareChain([]func(ReceiveFunc) ReceiveFunc{}, func(_ Context) {}))
}
