//
// Copyright 2021-2023 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package github

import (
	"os"
	"reflect"
	"testing"
)

func TestParseGitHubTokens(t *testing.T) {
	tests := []struct {
		name               string
		githubTokenEnv     string
		githubTokenListEnv string
		want               []string
		wantErr            bool
	}{
		{
			name:    "No tokens set",
			wantErr: true,
		},
		{
			name:           "Only one token, stored in GITHUB_AUTH_TOKEN",
			githubTokenEnv: "some_token",
			want:           []string{"some_token"},
		},
		{
			name:               "Only one token, stored in GITHUB_TOKEN_LIST",
			githubTokenListEnv: "list_token",
			want:               []string{"list_token"},
		},
		{
			name:               "Two tokens, one each stored in GITHUB_AUTH_TOKEN and GITHUB_TOKEN_LIST",
			githubTokenEnv:     "some_token",
			githubTokenListEnv: "list_token",
			want:               []string{"some_token", "list_token"},
		},
		{
			name:               "Multiple tokens",
			githubTokenEnv:     "some_token",
			githubTokenListEnv: "list_token,another_token,third_token",
			want:               []string{"some_token", "list_token", "another_token", "third_token"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Unsetenv("GITHUB_AUTH_TOKEN")
			os.Unsetenv("GITHUB_TOKEN_LIST")
			if tt.githubTokenEnv != "" {
				os.Setenv("GITHUB_AUTH_TOKEN", tt.githubTokenEnv)
			}
			if tt.githubTokenListEnv != "" {
				os.Setenv("GITHUB_TOKEN_LIST", tt.githubTokenListEnv)
			}

			err := ParseGitHubTokens()
			if tt.wantErr != (err != nil) {
				t.Errorf("TestParseGitHubTokens() error: unexpected error value %v", err)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(Tokens, tt.want) {
					t.Errorf("TestParseGitHubTokens() error: expected %v got %v", tt.want, Tokens)
				}
			}

		})
	}
}

func TestGetNewGitHubClient(t *testing.T) {
	ghTokenClient := GitHubTokenClient{}
	tests := []struct {
		name               string
		client             GitHubToken
		githubTokenEnv     string
		githubTokenListEnv string
		wantErr            bool
	}{
		{
			name:    "No tokens initialized, error should be returned",
			client:  ghTokenClient,
			wantErr: true,
		},
		{
			name:           "One token set, should return client",
			client:         ghTokenClient,
			githubTokenEnv: "some_token",
			wantErr:        false,
		},
		{
			name:               "Multiple tokens, should return client",
			client:             ghTokenClient,
			githubTokenEnv:     "some_token",
			githubTokenListEnv: "another_token,third_token",
			wantErr:            false,
		},
		{
			name:               "Multiple tokens, should return client",
			client:             ghTokenClient,
			githubTokenEnv:     "some_token",
			githubTokenListEnv: "another_token,third_token",
			wantErr:            false,
		},
		{
			name:    "Mock client",
			client:  MockGitHubTokenClient{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Tokens = nil
			os.Unsetenv("GITHUB_AUTH_TOKEN")
			os.Unsetenv("GITHUB_TOKEN_LIST")
			if tt.githubTokenEnv != "" {
				os.Setenv("GITHUB_AUTH_TOKEN", tt.githubTokenEnv)
			}
			if tt.githubTokenListEnv != "" {
				os.Setenv("GITHUB_TOKEN_LIST", tt.githubTokenListEnv)
			}

			if !tt.wantErr {
				_ = ParseGitHubTokens()
			}
			_, err := tt.client.GetNewGitHubClient()
			if tt.wantErr != (err != nil) {
				t.Errorf("TestGetNewGitHubClient() error: unexpected error value %v", err)
			}

		})
	}
}
