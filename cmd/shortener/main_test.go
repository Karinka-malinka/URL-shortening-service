package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func Test_mainHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name  string
		metod string
		body  string
		want  want
	}{
		{
			name:  "POST positive test #1",
			metod: http.MethodPost,
			body:  "https://practicum.yandex.ru/",
			want: want{
				code:        201,
				response:    `OK`,
				contentType: "text/plain",
			},
		},
		{
			name:  "POST negative test #2",
			metod: http.MethodPost,
			body:  "",
			want: want{
				code:        400,
				response:    `BAD`,
				contentType: "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			bodyReader := strings.NewReader(tt.body)

			request := httptest.NewRequest(tt.metod, "/", bodyReader)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			mainHandler(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, tt.want.code, res.StatusCode)

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			//resBody, err := io.ReadAll(res.Body)

			//require.NoError(t, err)
			//assert.JSONEq(t, tt.want.response, string(resBody))
			//assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
