package core_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"pantori/internal/domains/notifiers/core"
	"pantori/internal/domains/notifiers/mocks"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/stretchr/testify/assert"
)

type CreateReportCase struct {
	Description    string
	InputGoods     []core.Good
	ExpectedLog    string
	ExpectedReport core.Report
}

func TestCreateReport(t *testing.T) {
	testCases := []CreateReportCase{
		{
			Description: "successfull report",
			InputGoods: []core.Good{
				{
					Name:   "carrot",
					Expire: "01/01/2000",
				},
				{
					Name:   "grape",
					Expire: time.Now().Format("02/01/2006"),
				},
				{
					Name:   "eggs",
					Expire: time.Now().Add(24 * time.Hour).Format("02/01/2006"),
				},
			},
			ExpectedLog: "",
			ExpectedReport: core.Report{
				Expired: []core.Good{
					{
						Name:   "carrot",
						Expire: "01/01/2000",
					},
				},
				ExpiresToday: []core.Good{
					{
						Name:   "grape",
						Expire: time.Now().Format("02/01/2006"),
					},
				},
				ExpiresSoon: []core.Good{
					{
						Name:   "eggs",
						Expire: time.Now().Add(24 * time.Hour).Format("02/01/2006"),
					},
				},
			},
		},
		{
			Description: "skip empty goods",
			InputGoods: []core.Good{
				{
					Name:     "fish",
					Expire:   "01/01/2000",
					Quantity: "Empty",
				},
			},
			ExpectedLog:    "",
			ExpectedReport: core.Report{},
		},
		{
			Description: "invalid date",
			InputGoods: []core.Good{
				{
					Name:   "fish",
					Expire: "invalid",
				},
			},
			ExpectedLog:    "Given date have incorrect format: invalid",
			ExpectedReport: core.Report{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

			svc := core.NewService(
				&mocks.GoodsMocks{},
				&mocks.UsersMock{},
				&mocks.EmailMock{},
				1,
			)

			report := svc.CreateReport(testCase.InputGoods)

			logOutput := buf.String()

			assert.Equal(testCase.ExpectedReport, report)
			assert.Contains(logOutput, testCase.ExpectedLog)
		})
	}
}

type NotifyCase struct {
	Description           string
	WhenHaveGoodsToExpire bool
	WhenListUserErr       error
	WhenGetGoodsErr       error
	WhenSendEmailErr      error
	ExpectedLog           string
	ExpectedInvocation    string
}

func TestNotify(t *testing.T) {
	testCases := []NotifyCase{
		{
			Description:           "successfull case with report",
			WhenHaveGoodsToExpire: true,
			WhenListUserErr:       nil,
			WhenGetGoodsErr:       nil,
			WhenSendEmailErr:      nil,
			ExpectedLog:           "",
			ExpectedInvocation:    "-ListUsers-GetGoods-SendEmail-GetGoods-SendEmail",
		},
		{
			Description:           "successfull case without report",
			WhenHaveGoodsToExpire: false,
			WhenListUserErr:       nil,
			WhenGetGoodsErr:       nil,
			WhenSendEmailErr:      nil,
			ExpectedLog:           "",
			ExpectedInvocation:    "-ListUsers-GetGoods-GetGoods",
		},
		{
			Description:           "list users err",
			WhenHaveGoodsToExpire: false,
			WhenListUserErr:       errors.New("op failed"),
			WhenGetGoodsErr:       nil,
			WhenSendEmailErr:      nil,
			ExpectedLog:           "Something wrong while listing users: op failed",
			ExpectedInvocation:    "-ListUsers",
		},
		{
			Description:           "get goods err",
			WhenHaveGoodsToExpire: false,
			WhenListUserErr:       nil,
			WhenGetGoodsErr:       errors.New("op failed"),
			WhenSendEmailErr:      nil,
			ExpectedLog:           "Something wrong while getting goods: op failed",
			ExpectedInvocation:    "-ListUsers-GetGoods-GetGoods",
		},
		{
			Description:           "send email err",
			WhenHaveGoodsToExpire: true,
			WhenListUserErr:       nil,
			WhenGetGoodsErr:       nil,
			WhenSendEmailErr:      errors.New("op failed"),
			ExpectedLog:           "Something wrong while sending email: op failed",
			ExpectedInvocation:    "-ListUsers-GetGoods-SendEmail-GetGoods-SendEmail",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

			svc := core.NewService(
				&mocks.GoodsMocks{
					ErrGet:     testCase.WhenGetGoodsErr,
					HaveGoods:  testCase.WhenHaveGoodsToExpire,
					Invocation: &invocationTrail,
				},
				&mocks.UsersMock{
					ErrList:    testCase.WhenListUserErr,
					Invocation: &invocationTrail,
				},
				&mocks.EmailMock{
					ErrSend:    testCase.WhenSendEmailErr,
					Invocation: &invocationTrail,
				},
				1,
			)

			svc.NotifyExpiredGoods()

			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.Contains(logOutput, testCase.ExpectedLog)
		})
	}
}
