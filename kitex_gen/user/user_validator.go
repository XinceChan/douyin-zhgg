// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *User) IsValid() error {
	return nil
}
func (p *BaseResp) IsValid() error {
	return nil
}
func (p *DouyinUserRegisterRequest) IsValid() error {
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *DouyinUserRegisterResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *DouyinUserLoginRequest) IsValid() error {
	if len(p.Username) > int(32) {
		return fmt.Errorf("field Username max_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) > int(32) {
		return fmt.Errorf("field Password max_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *DouyinUserLoginResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *DouyinUserRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *DouyinUserResponse) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
