package listener

import "github.com/stretchr/testify/mock"

type publisherMock struct {
	mock.Mock
}

func (p *publisherMock) Publish(subject string, msg []byte) error {
	args := p.Called(subject, msg)
	return args.Error(0)
}

func (p *publisherMock) Close() error {
	return p.Called().Error(0)
}
