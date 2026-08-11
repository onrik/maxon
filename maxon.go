package maxon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
Maxon
https://dev.max.ru/docs-api
*/
type Maxon struct {
	apiURL string
	token  string
	http   *http.Client
}

func New(token string, opts ...Option) *Maxon {
	options := Options{
		apiURL:     "https://platform-api.max.ru",
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Maxon{
		apiURL: options.apiURL,
		token:  token,
		http:   options.httpClient,
	}
}

func (m *Maxon) do(request *http.Request, target any) error {
	request.Header.Set("Authorization", m.token)

	response, err := m.http.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= 500 {
		return fmt.Errorf("http status: %d", response.StatusCode)
	}

	if response.StatusCode >= 400 {
		maxErr := Error{}
		if err := json.Unmarshal(body, &maxErr); err != nil {
			return err
		}

		return maxErr
	}

	if target != nil {
		return json.Unmarshal(body, target)
	}

	return nil
}

func (m *Maxon) Get(ctx context.Context, path string, target any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, m.apiURL+path, nil)
	if err != nil {
		return err
	}

	return m.do(request, target)
}

func (m *Maxon) Delete(ctx context.Context, path string, target any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, m.apiURL+path, nil)
	if err != nil {
		return err
	}

	return m.do(request, target)
}

func (m *Maxon) Post(ctx context.Context, path string, data, target any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, m.apiURL+path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	return m.do(request, target)
}

func (m *Maxon) Patch(ctx context.Context, path string, data, target any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPatch, m.apiURL+path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	return m.do(request, target)
}

func (m *Maxon) Put(ctx context.Context, path string, data, target any) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPut, m.apiURL+path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	return m.do(request, target)
}

func (m *Maxon) Me(ctx context.Context) (User, error) {
	user := User{}
	err := m.Get(ctx, "/me", &user)
	return user, err
}

func (m *Maxon) Updates(ctx context.Context) ([]Update, error) {
	response := updatesResponse{}
	err := m.Get(ctx, "/updates", &response)
	return response.Updates, err
}

func (m *Maxon) MessageSend(ctx context.Context, chatID int64, text string, options *MessageSendOptions) (Message, error) {
	data := map[string]any{
		"text": text,
	}
	if options != nil {
		data["disable_link_preview"] = options.DisableLinkPreview
		data["notify"] = !options.NotifyDisable
		if options.Format != "" {
			data["format"] = options.Format
		}
	}

	response := messageResponse{}
	err := m.Post(ctx, fmt.Sprintf("/messages?chat_id=%d", chatID), data, &response)

	return response.Message, err
}

func (m *Maxon) MessageEdit(ctx context.Context, messageID, text string, options *MessageEditOptions) error {
	data := map[string]any{
		"text": text,
	}
	if options != nil {
		data["disable_link_preview"] = options.DisableLinkPreview
		data["notify"] = !options.NotifyDisable
		if options.Format != "" {
			data["format"] = options.Format
		}
	}

	err := m.Put(ctx, fmt.Sprintf("/messages?message_id=%s", messageID), data, nil)
	return err
}

func (m *Maxon) MessageDelete(ctx context.Context, messageID string) error {
	err := m.Delete(ctx, fmt.Sprintf("/messages?message_id=%s", messageID), nil)
	return err
}
